package parser

import (
	"encoding/hex"
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Parser struct {
	Program      *LexerResult
	ContractBase uip.ContractBase
	AccountBase  uip.AccountBase
	ChainGetter  uip.BlockChainGetter
}


func (ier *Parser) marshal(v interface{}) (b []byte, err error) {
	b, err = json.Marshal(v)
	if err != nil {
		return nil, errorn.NewMarshalError(err)
	}
	return
}
//marshal

func (ier *Parser) decodeHex(src string) (b []byte, err error) {
	b, err = hex.DecodeString(src)
	if err != nil {
		return nil, errorn.NewUnmarshalError(err)
	}
	return
}

type LexerResult struct {
	RootIntents     *lexer.RootIntents
	RawDependencies *lexer.RawDependencies
	ContractMapping lexer.AccountMap
	AccountMapping  lexer.AccountMap
}

