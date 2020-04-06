package lexer

import "github.com/HyperService-Consortium/go-uip/op-intent/lexer/internal"

type ContractInvokeMeta struct {
	Code     []byte           `json:"contract_code"` // key
	FuncName string           `json:"func"`
	Params   []internal.Param `json:"parameters"`
	Meta     []byte           `json:"meta"`
}
