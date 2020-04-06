package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Account = token.Token

type Param interface {
	GetParamType() value_type.Type
	Determine(f InstantiateAccountF) (token.Param, error)
}

type InstantiateAccountF = func(a Account) (uip.Account, error)
