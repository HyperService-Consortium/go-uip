package lexer

import (
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
)

type ContractInvokeMeta struct {
	Code     []byte              `json:"contract_code"` // key
	FuncName string              `json:"func"`
	Params   []token_types.Param `json:"parameters"`
	Meta     []byte              `json:"meta"`
}
