package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
)

type ContractInvokeMeta struct {
	Code     []byte        `json:"contract_code"` // key
	FuncName string        `json:"func"`
	Params   []token.Param `json:"parameters"`
	Meta     []byte        `json:"meta"`
}
