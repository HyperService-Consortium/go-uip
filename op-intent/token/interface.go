package token

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type NamespacedRawAccountI = uip.Account
type RawAccountI interface {
	GetAddress() []byte
}

type NamespacedNameAccountI interface {
	GetName() string
	GetChainID() uip.ChainIDUnderlyingType
}

type NameAccountI interface {
	GetName() string
}

type Token interface {
	GetType() Type
}

type Param interface {
	Token
	GetParamType() value_type.Type
}


type stateVariableI interface {
	GetPos() []byte
	GetField() []byte
}

type StateVariableI interface {
	GetContract() Token
	LocalStateVariableI
}

type LocalStateVariableI interface {
	Param
	stateVariableI
}

type ConstantI interface {
	Param
	GetConstant() interface{}
}

type UnaryExpressionI interface {
	Param
	GetSign() sign_type.Type
	GetLeft() Param
}

type BinaryExpressionI interface {
	UnaryExpressionI
	GetRight() Param
}

//type StateVariable struct {
//	Type     value_type.Type `json:"type"`
//	Contract Account `json:"contract"`
//	Pos      []byte `json:"pos"`
//	Field    []byte `json:"field"`
//}