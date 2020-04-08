package lexer_types

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
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

func (n NamespacedNameAccount) GetType() token_type.Type {
	return token_type.NamespacedNameAccount
}
