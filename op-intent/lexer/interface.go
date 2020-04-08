package lexer

import "github.com/HyperService-Consortium/go-uip/internal/token_types"

type Intent interface {
	token_types.Token
	GetName() string
}
