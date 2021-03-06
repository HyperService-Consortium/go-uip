package parser

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
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

func (ier *Parser) fillIndex(intents TxIntentsImpl) (TxIntentsImpl, error) {
	var nameMap = make(map[string]int)
	for i := range intents {
		if j, ok := nameMap[intents[i].GetName()]; ok {
			return nil, fmt.Errorf("name conflict: %v and %v with name %v", i, j, intents[i].GetName())
		}
		nameMap[intents[i].GetName()] = i
	}
	for i := range intents {
		inst := intents[i].GetInstruction()
		switch inst.GetType() {
		case instruction_type.RawConditionGoto:
			ri := inst.(*instruction.RawConditionGoto)
			cond, err := ri.Condition.Determine(ier)
			if err != nil {
				return nil, err
			}
			index := nameMap[ri.IndexName] + ri.Offset
			if index < 0 {
				return nil, errorn.NewInvalidFieldError(errorn.ErrNegativeIndex)
			}
			intents[i].SetInstruction(instruction.NewConditionGoto(uint64(index), cond))
		case instruction_type.RawGoto:
			ri := inst.(*instruction.RawGoto)
			index := nameMap[ri.IndexName] + ri.Offset
			if index < 0 {
				return nil, errorn.NewInvalidFieldError(errorn.ErrNegativeIndex)
			}
			intents[i].SetInstruction(instruction.NewGoto(uint64(index)))
		case instruction_type.RawConditionSetState:
			//ri := inst.(*RawGoto)
		case instruction_type.RawSetState:
			ri := inst.(*instruction.RawSetState)
			rhs, err := ri.RightExpression.Determine(ier)
			if err != nil {
				return nil, err
			}
			intents[i].SetInstruction(instruction.NewSetState(ri.Target.Type, string(ri.Target.Field), rhs))
		}
	}
	return intents, nil
}

type LexerResult struct {
	RootIntents     *lexer.RootIntents
	RawDependencies *lexer.RawDependencies
	ContractMapping lexer_types.AccountMap
	AccountMapping  lexer_types.AccountMap
}
