package opintent

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type TransactionIntent = instruction.TransactionIntent
type Goto = instruction.Goto
type RawGoto = instruction.RawGoto
type ConditionGoto = instruction.ConditionGoto
type RawConditionGoto = instruction.RawConditionGoto
type SetState = instruction.SetState
type RawSetState = instruction.RawSetState
type ConditionSetState = instruction.ConditionSetState

func NewGoto(index uint64) *Goto {
	return instruction.NewGoto(index)
}

func NewConditionGoto(index uint64, condition uip.VTok) *ConditionGoto {
	return instruction.NewConditionGoto(index, condition)
}

func NewSetState(t value_type.Type, target string, rhs uip.VTok) *SetState {
	return instruction.NewSetState(t, target, rhs)
}

func NewConditionSetState(t value_type.Type, target string, rhs uip.VTok, condition uip.VTok) *ConditionSetState {
	return instruction.NewConditionSetState(t, target, rhs, condition)
}
