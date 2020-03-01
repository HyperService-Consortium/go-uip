package parser

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

// if cond [if-op]
//    op1
//    op2
//    ...
//    opk
// else
//    opk+1
//    opk+2
//    ...
//    opk+n

// ->

// goto index(op1) if cond
// opk+1
// opk+2
// ...
// opk+n
// goto index(opk)+1 if true
// op1
// op2
// ...
// opk
//


type RawConditionGoto struct {
	IndexName     string
	Condition  lexer.Param
	Offset int
}

func (r RawConditionGoto) GetType() instruction_type.Type {
	return instruction_type.RawConditionGoto
}

type RawGoto struct {
	IndexName     string
	Offset int
}

func (r RawGoto) GetType() instruction_type.Type {
	return instruction_type.RawGoto
}


func (ier * Parser) parseIf(rawIntent *lexer.IfIntent) (intents []uip.TxIntentI, err error) {
	ifName := rawIntent.GetName()

	ifIntents, err := ier.parseIntents(rawIntent.If)
	if err != nil {
		return nil, err
	}

	elseIntents, err := ier.parseIntents(rawIntent.Else)
	if err != nil {
		return nil, err
	}

	// goto index(op1) if cond
	gotoOp1IfCond := &RawConditionGoto{
		IndexName: "",
		Condition: rawIntent.Condition,
		Offset:    0,
	}
	txGotoOp1IfCond := newIntent(gotoOp1IfCond, ifName + ".goto.if")
	// goto index(opk)+1
	gotoOpkP1 := &RawGoto{
		IndexName: "",
		Offset:      0,
	}
	txGotoIndexOpkP1 := newIntent(gotoOpkP1, ifName + ".goto.else")


	if ifIntents.Len() != 0 {
		op1 := ifIntents.GetTxIntents()[0]
		gotoOp1IfCond.IndexName = op1.GetName()
	} else {
		gotoOp1IfCond.IndexName = txGotoIndexOpkP1.Name
		gotoOp1IfCond.Offset = 1
	}

	if ifIntents.Len() != 0 {
		opk := ifIntents.GetTxIntents()[ifIntents.Len() - 1]
		gotoOpkP1.IndexName = opk.GetName()
		gotoOpkP1.Offset = 1
	} else {
		gotoOpkP1.IndexName = txGotoIndexOpkP1.Name
		gotoOpkP1.Offset = 1
	}

	// goto index(op1) if cond
	intents = append(intents, txGotoOp1IfCond)
	// opk+1
	// opk+2
	// ...
	// opk+n
	intents = append(intents, elseIntents...)
	// goto index(opk)+1
	intents = append(intents, txGotoIndexOpkP1)
	// op1
	// op2
	// ...
	// opk
	intents = append(intents, ifIntents...)

	return
}