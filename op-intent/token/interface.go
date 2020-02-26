package token

import (
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

type StateVariableI interface {
	GetParamType() value_type.Type
	GetContract() Token
	GetPos() []byte
	GetField() []byte
}

type ConstantI interface {
	GetParamType() value_type.Type
	GetConstant() interface{}
}
//type StateVariable struct {
//	Type     value_type.Type `json:"type"`
//	Contract Account `json:"contract"`
//	Pos      []byte `json:"pos"`
//	Field    []byte `json:"field"`
//}