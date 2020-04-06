package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"io"
)

type RawAccount struct {
	Address []byte `json:"address"`
}

func (n RawAccount) Marshal(w io.Writer, err *error) {
	serial.Write(w, n.Address, err)
}

func (n RawAccount) Unmarshal(r io.Reader, i *Account, err *error) {
	serial.Read(r, &n.Address, err)
	*i = n
}

func (n RawAccount) GetAddress() []byte {
	return n.Address
}

func (n RawAccount) GetType() token.Type {
	return token.RawAccount
}
