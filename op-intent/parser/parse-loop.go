package parser

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)
// loop times
//    op1
//    op2
//    ...
//    opk

// ->

// loopBegin: goto reset if loopVar >= times
// op1
// op2
// ...
// opk
// loopVar := loopVar + 1
// loopEnd: goto loopBegin if true
// reset: loopVar := 0


type RawSetState struct {
	Target *lexer.LocalStateVariable
	RightExpression lexer.Param
}

func (r RawSetState) GetType() instruction_type.Type {
	return instruction_type.RawSetState
}

// WARNING: will lead to a wrong result if using function

const loopVarType = value_type.Int64
func (ier *Parser) parseLoop(intent *lexer.LoopIntent) (intents []uip.TxIntentI, _ error) {
	loopIntents, err := ier.parseIntents(intent.Loop)
	if err != nil {
		return nil, err
	}

	loopVar := &lexer.LocalStateVariable{
		Type:  loopVarType,
		Pos:   nil,
		Field: []byte(intent.GetName() + ".loopVar"),
	}

	addOpLoopVar := newIntent(&RawSetState{
		Target:          loopVar,
		RightExpression: &lexer.BinaryExpression{
			Type: loopVarType,
			Sign: sign_type.ADD,
			Left: loopVar,
			Right: &lexer.ConstantVariable{
				Type:  loopVarType,
				Const: 1,
			},
		},
	}, intent.GetName() + ".addLoopVar")

	resetLoopVar := newIntent(&RawSetState{
		Target:          loopVar,
		RightExpression: &lexer.ConstantVariable{
			Type:  loopVarType,
			Const: 0,
		},
	}, intent.GetName() + ".resetLoopVar")

	loopBegin := newIntent(&RawConditionGoto{
		IndexName: resetLoopVar.GetName(),
		Condition: &lexer.BinaryExpression{
			Type:  value_type.Bool,
			Sign:  sign_type.GE,
			Left:  loopVar,
			// todo: convert Times to loopVarType
			Right: &lexer.ConstantVariable{
				Type:  loopVarType,
				Const: int64(intent.Times),
			},
		},
		Offset:    0,
	}, intent.GetName() + ".loopBegin")

	loopEnd := newIntent(&RawGoto{
		IndexName: loopBegin.GetName(),
		Offset:    0,
	}, intent.GetName() + ".loopEnd")


	// loopBegin: goto reset if loopVar < times
	intents = append(intents, loopBegin)
	// op1
	// op2
	// ...
	// opk
	intents = append(intents, loopIntents...)
	// loopVar := loopVar + 1
	intents = append(intents, addOpLoopVar)
	// loopEnd: goto loopBegin if true
	intents = append(intents, loopEnd)
	// reset: loopVar := 0
	intents = append(intents, resetLoopVar)
	return
}
