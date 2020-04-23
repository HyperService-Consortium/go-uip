package instruction

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction/internal"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Goto = internal.Goto
type RawGoto = internal.RawGoto
type ConditionGoto = internal.ConditionGoto
type RawConditionGoto = internal.RawConditionGoto

type SetState = internal.SetState
type RawSetState = internal.RawSetState
type ConditionSetState = internal.ConditionSetState
type TransactionIntent = internal.TransactionIntent
type NamedIntent = internal.NamedIntent

func NewGoto(index uint64) *Goto {
	return internal.NewGoto(index)
}

func NewConditionGoto(index uint64, condition uip.VTok) *ConditionGoto {
	return internal.NewConditionGoto(index, condition)
}

func NewSetState(t value_type.Type, target string, rhs uip.VTok) *SetState {
	return internal.NewSetState(t, target, rhs)
}

func NewConditionSetState(t value_type.Type, target string, rhs uip.VTok, condition uip.VTok) *ConditionSetState {
	return internal.NewConditionSetState(t, target, rhs, condition)
}
