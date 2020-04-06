package instruction

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/instruction/internal"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type Goto = internal.Goto
type RawGoto = internal.RawGoto
type ConditionGoto = internal.ConditionGoto
type RawConditionGoto = internal.RawConditionGoto

type SetState = internal.SetState
type RawSetState = internal.RawSetState
type ConditionSetState = internal.ConditionSetState

func NewGoto(index uint64) *Goto {
	return internal.NewGoto(index)
}

func NewConditionGoto(index uint64, condition json.RawMessage) *ConditionGoto {
	return internal.NewConditionGoto(index, condition)
}

func NewSetState(t value_type.Type, target []byte, rhs json.RawMessage) *SetState {
	return internal.NewSetState(t, target, rhs)
}

func NewConditionSetState(t value_type.Type, target []byte, rhs json.RawMessage, condition json.RawMessage) *ConditionSetState {
	return internal.NewConditionSetState(t, target, rhs, condition)
}

type hexstring = string

type NamedIntent struct {
	Name string `json:"-"`
}

func (n NamedIntent) GetName() string {
	return n.Name
}

type TransactionIntent struct {
	TransType trans_type.Type `json:"itype"`
	Src       []byte          `json:"src"`
	Dst       []byte          `json:"dst"`
	Meta      json.RawMessage `json:"meta"`
	Amt       hexstring       `json:"amt"`
	ChainID   uint64          `json:"chain_id"`
}

func (tx *TransactionIntent) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (tx *TransactionIntent) Exec(g *gvm.ExecCtx) error {
	panic("implement me")
}

func (tx *TransactionIntent) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (tx *TransactionIntent) GetType() instruction_type.Type {
	return tx.TransType
}

func (tx *TransactionIntent) Bytes() []byte {
	b, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return b
}
