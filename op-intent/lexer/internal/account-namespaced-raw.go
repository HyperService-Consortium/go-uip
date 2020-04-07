package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"io"
)

type NamespacedRawAccount struct {
	Address []byte                    `json:"address"`
	ChainID uip.ChainIDUnderlyingType `json:"chain_id"`
}

func (n NamespacedRawAccount) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, n.Address, err)
	serial.Write(w, n.ChainID, err)
}

func (n NamespacedRawAccount) Unmarshal(r io.Reader, i *Account, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &n.Address, err)
	serial.Read(r, &n.ChainID, err)
	*i = n
}

func (n NamespacedRawAccount) GetChainId() uip.ChainID {
	return n.ChainID
}

func (n NamespacedRawAccount) GetAddress() uip.Address {
	return n.Address
}

func NewNamespacedRawAccount(a uip.Account) *NamespacedRawAccount {
	return &NamespacedRawAccount{
		Address: a.GetAddress(),
		ChainID: a.GetChainId(),
	}
}

func (n NamespacedRawAccount) GetType() token.Type {
	return token.NamespacedRawAccount
}
