package lexer

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"strings"
)

type Param interface {
	Token
	GetParamType() value_type.Type
}

type ParamImpl struct {
	Type string `json:"Type"`
}

func ParamUnmarshalJSON(b []byte) (r Param, err error) {
	c, err := document.NewGJSONDocument(b)
	if err != nil {
		return
	}
	return ParamUnmarshalResult(c)
}

func DecodeAddress(src string) ([]byte, error) {
	if strings.HasPrefix(src, "0x") {
		return hex.DecodeString(src[2:])
	}
	return hex.DecodeString(src)
}

func DecodeContractPos(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

type ConstantVariable struct {
	Type  value_type.Type `json:"type"`
	Const interface{}
}

func (p ConstantVariable) GetType() token.Type {
	return token.Constant
}

func (p ConstantVariable) GetParamType() value_type.Type {
	return p.Type
}

type StateVariable struct {
	Type     value_type.Type `json:"type"`
	Contract Account
	Pos      []byte `json:"pos"`
	Field    []byte `json:"field"`
}

func (e StateVariable) GetType() token.Type {
	return token.StateVariable
}

func (e StateVariable) GetParamType() value_type.Type {
	return e.Type
}

func ParamUnmarshalResult(content document.Document) (p Param, err error) {
	v := content.Get(FieldOpIntentsType)
	if !v.Exists() {
		return nil, errorn.NewFieldNotFound(FieldOpIntentsType)
	}
	var intDesc value_type.Type
	if intDesc = value_type.FromString(v.String()); intDesc == value_type.Unknown {
		return nil, errorn.NewValueTypeNotFound(v.String())
	}
	content = content.Get(FieldOpIntentsValue)
	if !content.Exists() {
		return nil, errorn.NewFieldNotFound(FieldOpIntentsValue)
	}

	v = content.Get(FieldValueConstant)
	if v.Exists() {
		//todo v.Value() -> intDesc
		return &ConstantVariable{Type: intDesc, Const: v.Value()}, nil
	} else if content.Get(FieldContractPos).Exists() &&
		content.Get(FieldContractField).Exists() {
		//if err != nil {
		//	return nil, errorn.NewDecodeAddressError(err)
		//}
		pos, err := DecodeContractPos(content.Get(FieldContractPos).String())
		if err != nil {
			return nil, errorn.NewDecodeContractPosError(err)
		}
		acc, err := AccountUnmarshalResult(FieldContractAccount, content)
		if err != nil {
			return nil, err
		}
		return &StateVariable{Type: intDesc, Contract: acc,
			Pos: pos, Field: []byte(content.Get(FieldContractField).String())}, nil

	}
	//else {
	//	return nil, errorn.NewNotEnoughParamInformation()
	//}

	return nil, errorn.NewInvalidFieldError(errorn.UnknownParam)
}

func initParamsR(i document.Document) (params []Param, err error) {
	if i.Exists() && !i.IsArray() {
		return nil, errorn.NewInvalidFieldError(errorn.ErrTypeError).Desc(errorn.AtOpIntentField{Field: FieldOpIntentsParameters})
	}
	rawParams := i.Array()
	params = make([]Param, rawParams.Len())
	for i := 0; i < rawParams.Len(); i++ {
		params[i], err = ParamUnmarshalResult(rawParams.Index(i))
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentParameterPos{Pos: i})
		}
	}
	return
}
