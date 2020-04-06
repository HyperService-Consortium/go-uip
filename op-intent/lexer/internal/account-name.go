package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"io"
)

type NameAccount struct {
	Name string `json:"name"`
}

func (n NameAccount) Marshal(w io.Writer, err *error) {
	serial.Write(w, n.Name, err)
}

func (n NameAccount) Unmarshal(r io.Reader, i *Account, err *error) {
	serial.Read(r, &n.Name, err)
	*i = n
}

func (n NameAccount) GetName() string {
	return n.Name
}

func (n NameAccount) GetType() token.Type {
	return token.NameAccount
}
