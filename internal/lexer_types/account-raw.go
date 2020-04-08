package lexer_types

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
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

func (n RawAccount) GetType() token_type.Type {
	return token_type.RawAccount
}
