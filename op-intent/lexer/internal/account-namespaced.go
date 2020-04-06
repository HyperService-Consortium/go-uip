package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type NamespacedAccount struct {
	Name    string                    `json:"name"`
	ChainID uip.ChainIDUnderlyingType `json:"chain_id"`
}

func (n NamespacedAccount) GetName() string {
	return n.Name
}

func (n NamespacedAccount) GetChainID() uip.ChainIDUnderlyingType {
	return n.ChainID
}

func (n NamespacedAccount) GetType() token.Type {
	return token.NamespacedNameAccount
}
