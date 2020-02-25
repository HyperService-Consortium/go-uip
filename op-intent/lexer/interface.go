package lexer

import "github.com/HyperService-Consortium/go-uip/op-intent/token"

type Token interface {
	GetType() token.Type
}

type Intent interface {
	Token
	GetName() string
}

