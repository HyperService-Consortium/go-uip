package internal

import "github.com/HyperService-Consortium/go-uip/op-intent/token"

type NameAccount struct {
	Name string `json:"name"`
}

func (n NameAccount) GetName() string {
	return n.Name
}

func (n NameAccount) GetType() token.Type {
	return token.NameAccount
}
