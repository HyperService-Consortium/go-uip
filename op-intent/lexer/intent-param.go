package lexer

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
	"strings"
)

type InstantiateAccountF = func (a Account) (uip.Account, error)
type Param interface {
	token.Token
	GetParamType() value_type.Type
	Determine(f InstantiateAccountF) (Param, error)
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

func decodeHex(src string) (b []byte, err error) {
	b, err = hex.DecodeString(src)
	if err != nil {
		return nil, errorn.NewDecodeHexError(err)
	}
	return b, nil
}

func DecodeAddress(src string) ([]byte, error) {
	if strings.HasPrefix(src, "0x") {
		return decodeHex(src[2:])
	}
	return decodeHex(src)
}

func DecodeContractPos(src string) ([]byte, error) {
	if strings.HasPrefix(src, "0x") {
		return decodeHex(src[2:])
	}
	return decodeHex(src)
}

type ConstantVariable struct {
	Type  value_type.Type `json:"type"`
	Const interface{} `json:"constant"`
}

func (p ConstantVariable) GetConstant() interface{} {

	return p.Const
}

func (p ConstantVariable) Determine(_ InstantiateAccountF) (Param, error) {
	return p, nil
}

func (p ConstantVariable) GetType() token.Type {
	return token.Constant
}

func (p ConstantVariable) GetParamType() value_type.Type {
	return p.Type
}

type LocalStateVariable struct {
	Type     value_type.Type `json:"type"`
	Pos      []byte `json:"pos"`
	Field    []byte `json:"field"`
}

func (l LocalStateVariable) GetType() token.Type {
	return token.LocalStateVariable
}

func (l LocalStateVariable) GetParamType() value_type.Type {
	return l.Type
}

func (e LocalStateVariable) GetPos() []byte {
	return e.Pos
}

func (e LocalStateVariable) GetField() []byte {
	return e.Field
}

func (l LocalStateVariable) Determine(f InstantiateAccountF) (Param, error) {
	return l, nil
}

type StateVariable struct {
	Type     value_type.Type `json:"type"`
	Contract Account `json:"contract"`
	Pos      []byte `json:"pos"`
	Field    []byte `json:"field"`
}

func (e StateVariable) GetType() token.Type {
	return token.StateVariable
}

func (e StateVariable) GetParamType() value_type.Type {
	return e.Type
}

func (e StateVariable) GetContract() token.Token {
	return e.Contract
}

func (e StateVariable) GetPos() []byte {
	return e.Pos
}

func (e StateVariable) GetField() []byte {
	return e.Field
}

func (e StateVariable) Determine(f InstantiateAccountF) (Param, error) {
	a, err := f(e.Contract)
	if err!= nil {
		return nil, err
	}
	e.Contract = NewNamespacedRawAccount(a)
	return e, nil
}

type BinaryExpression struct {
	Type  value_type.Type `json:"type"`
	Sign sign_type.Type `json:"sign"`
	Left Param `json:"left"`
	Right Param `json:"right"`
}
func (b BinaryExpression) GetType() token.Type {
	return token.BinaryExpression
}

func (b BinaryExpression) GetSign() sign_type.Type {
	return b.Sign
}

func (b BinaryExpression) GetLeft() token.Param {
	return b.Left
}

func (b BinaryExpression) GetRight() token.Param {
	return b.Right
}

func (b BinaryExpression) GetParamType() value_type.Type {
	return b.Type
}

func (b BinaryExpression) Determine(f InstantiateAccountF) (_ Param, err error) {
	b.Left, err = b.Left.Determine(f)
	if err != nil {
		return nil, err
	}

	b.Right, err = b.Right.Determine(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Param.type == value_type.Bool
type UnaryExpression struct {
	Type  value_type.Type `json:"type"`
	Sign sign_type.Type `json:"sign"`
	Left Param `json:"left"`
}

func (u UnaryExpression) GetType() token.Type {
	return token.UnaryExpression
}

func (u UnaryExpression) GetSign() sign_type.Type {
	return u.Sign
}

func (u UnaryExpression) GetLeft() token.Param {
	return u.Left
}

func (u UnaryExpression) GetParamType() value_type.Type {
	return u.Type
}

func (u UnaryExpression) Determine(f InstantiateAccountF) (_ Param, err error) {
	u.Left, err = u.Left.Determine(f)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Greater

var SignTable = map[string]sign_type.Type{
	"Greater": sign_type.GT,
}



func ParamUnmarshalResult(content document.Document) (p Param, err error) {

	v := content.Get(FieldOpIntentsSign)
	if v.Exists() {
		sign, ok := SignTable[v.String()]
		if !ok {
			return nil, errorn.NewSignTypeNotFound(v.String())
		}

		v = content.Get(FieldKeyLeft)
		if !v.Exists() {
			return nil, errorn.NewFieldNotFound(FieldKeyLeft)
		}

		left, err := ParamUnmarshalResult(content.Get(FieldKeyLeft))
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field:FieldKeyLeft})
		}

		v = content.Get(FieldKeyRight)
		if !v.Exists() {
			return UnaryExpression{Type:left.GetParamType(), Sign: sign, Left:left}, nil
		}

		right, err := ParamUnmarshalResult(v)
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field:FieldKeyRight})
		}

		//if left.GetParamType() != right.GetParamType() {
		//	return nil,
		//}
		t := left.GetParamType()
		if sign_type.IsLogic(sign) {
			t = value_type.Bool
		}
		// todo: determine param type of non-boolean expression
		return BinaryExpression{Type:t, Sign: sign, Left:left, Right:right}, nil
	}

	v = content.Get(FieldKeyType)
	if !v.Exists() {
		return nil, errorn.NewFieldNotFound(FieldKeyType)
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
