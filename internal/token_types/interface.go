package token_types

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type NamespacedRawAccountI = uip.Account
type RawAccountI interface {
	GetAddress() []byte
}

type FullAccountI interface {
	uip.Account
	GetName() string
}

type NamespacedNameAccountI interface {
	GetName() string
	GetChainID() uip.ChainIDUnderlyingType
}

type NameAccountI interface {
	GetName() string
}

type Token interface {
	GetType() token_type.Type
}

type Param = uip.VTok

type stateVariableI interface {
	GetPos() []byte
	GetField() []byte
}

type LocalStateVariableI interface {
	Param
	stateVariableI
}

type StateVariableI interface {
	GetContract() Token
	LocalStateVariableI
}

type ConstantI interface {
	Param
	Unwrap() interface{}
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
