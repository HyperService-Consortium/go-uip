package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Account = token.Token

type Param interface {
	token.Param
	Determine(f InstantiateAccountF) (Param, error)
}

type InstantiateAccountF = func(a Account) (uip.Account, error)
