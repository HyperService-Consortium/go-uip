package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type NamespacedRawAccount struct {
	Address []byte                    `json:"address"`
	ChainID uip.ChainIDUnderlyingType `json:"chain_id"`
}

func (n NamespacedRawAccount) GetChainId() uip.ChainID {
	return n.ChainID
}

func (n NamespacedRawAccount) GetAddress() uip.Address {
	return n.Address
}

func NewNamespacedRawAccount(a uip.Account) Account {
	return &NamespacedRawAccount{
		Address: a.GetAddress(),
		ChainID: a.GetChainId(),
	}
}

func (n NamespacedRawAccount) GetType() token.Type {
	return token.NamespacedRawAccount
}
