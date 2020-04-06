package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"io"
)

type NamespacedNameAccount struct {
	Name    string                    `json:"name"`
	ChainID uip.ChainIDUnderlyingType `json:"chain_id"`
}

func (n NamespacedNameAccount) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, n.Name, err)
	serial.Write(w, n.ChainID, err)
}

func (n NamespacedNameAccount) Unmarshal(r io.Reader, i *Account, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &n.Name, err)
	serial.Read(r, &n.ChainID, err)
	*i = n
}

func (n NamespacedNameAccount) GetName() string {
	return n.Name
}

func (n NamespacedNameAccount) GetChainID() uip.ChainIDUnderlyingType {
	return n.ChainID
}

func (n NamespacedNameAccount) GetType() token.Type {
	return token.NamespacedNameAccount
}
