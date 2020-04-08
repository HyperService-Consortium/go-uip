package lexer

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"math/big"
)

//noinspection GoUnusedExportedFunction
func ParamUnmarshalJSON(b []byte) (r lexer_types.Param, err error) {
	c, err := document.NewGJSONDocument(b)
	if err != nil {
		return
	}
	return ParamUnmarshalResult(c)
}

func CreateConstantFromJSON(t value_type.Type, v document.Document) lexer_types.DeterminedParam {
	switch t {
	case value_type.Uint256:
		return (*lexer_types.Uint256)(big.NewInt(v.Int()))
	}
	panic("implement me")
}

//Greater

var SignTable = map[string]sign_type.Type{
	"Greater": sign_type.GT,
}

func ParamUnmarshalResult(content document.Document) (p lexer_types.Param, err error) {

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
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: FieldKeyLeft})
		}

		v = content.Get(FieldKeyRight)
		if !v.Exists() {
			return lexer_types.UnaryExpression{Type: value_type.Type(left.GetGVMType()), Sign: sign, Left: left}, nil
		}

		right, err := ParamUnmarshalResult(v)
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: FieldKeyRight})
		}

		//if left.GetParamType() != right.GetParamType() {
		//	return nil,
		//}
		t := value_type.Type(left.GetGVMType())
		if sign_type.IsLogic(sign) {
			t = value_type.Bool
		}
		// todo: determine param type of non-boolean expression
		return lexer_types.BinaryExpression{Type: t, Sign: sign, Left: left, Right: right}, nil
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
		return CreateConstantFromJSON(intDesc, v), nil
	} else if content.Get(FieldContractPos).Exists() &&
		content.Get(FieldContractField).Exists() {
		//if err != nil {
		//	return nil, errorn.NewDecodeAddressError(err)
		//}
		pos, err := lexer_types.DecodeContractPos(content.Get(FieldContractPos).String())
		if err != nil {
			return nil, errorn.NewDecodeContractPosError(err)
		}
		acc, err := AccountUnmarshalResult(FieldContractAccount, content)
		if err != nil {
			return nil, err
		}
		return &lexer_types.StateVariable{Type: intDesc, Contract: acc,
			Pos: pos, Field: []byte(content.Get(FieldContractField).String())}, nil

	}
	//else {
	//	return nil, errorn.NewNotEnoughParamInformation()
	//}

	return nil, errorn.NewInvalidFieldError(errorn.UnknownParam)
}

func initParamsR(i document.Document) (params []lexer_types.Param, err error) {
	if i.Exists() && !i.IsArray() {
		return nil, errorn.NewInvalidFieldError(errorn.ErrTypeError).Desc(errorn.AtOpIntentField{Field: FieldOpIntentsParameters})
	}
	rawParams := i.Array()
	params = make([]lexer_types.Param, rawParams.Len())
	for i := 0; i < rawParams.Len(); i++ {
		params[i], err = ParamUnmarshalResult(rawParams.Index(i))
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentParameterPos{Pos: i})
		}
	}
	return
}
