package internal

import "github.com/HyperService-Consortium/go-uip/op-intent/token"

type RawAccount struct {
	Address []byte `json:"address"`
}

func (n RawAccount) GetAddress() []byte {
	return n.Address
}

func (n RawAccount) GetType() token.Type {
	return token.RawAccount
}
