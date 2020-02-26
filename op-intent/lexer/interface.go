package lexer

import "github.com/HyperService-Consortium/go-uip/op-intent/token"

type Intent interface {
	token.Token
	GetName() string
}

