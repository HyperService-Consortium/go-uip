package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
)

type Account = token.Token

type Param interface {
	GetGVMType() gvm.RefType
	Determine(f InstantiateAccountF) (token.Param, error)
}

type DeterminedParam interface {
	Determine(f InstantiateAccountF) (token.Param, error)
	token.Param
}

type InstantiateAccountF = func(a Account) (uip.Account, error)
