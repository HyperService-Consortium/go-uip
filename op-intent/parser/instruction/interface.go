package instruction

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/instruction/internal"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser/trap"
	"github.com/HyperService-Consortium/go-uip/serial"
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

func NewConditionGoto(index uint64, condition uip.VTok) *ConditionGoto {
	return internal.NewConditionGoto(index, condition)
}

func NewSetState(t value_type.Type, target string, rhs uip.VTok) *SetState {
	return internal.NewSetState(t, target, rhs)
}

func NewConditionSetState(t value_type.Type, target string, rhs uip.VTok, condition uip.VTok) *ConditionSetState {
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
	if *err != nil {
		return
	}
	serial.Write(w, tx.TransType, err)
	serial.Write(w, tx.Src, err)
	serial.Write(w, tx.Dst, err)
	serial.Write(w, tx.Amt, err)
	serial.Write(w, tx.ChainID, err)
	serial.Write(w, []byte(tx.Meta), err)
}

func (tx *TransactionIntent) Exec(c *gvm.ExecCtx) error {
	isc := c.Machine.(uip.ISCMachine)
	if isc.GetMuPC() != TxState.Closed {
		return trap.ClaimRequest
	}

	isc.SetMuPC(TxState.Unknown)
	c.PC++
	return nil
}

func (tx *TransactionIntent) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &tx.TransType, err)
	serial.Read(r, &tx.Src, err)
	serial.Read(r, &tx.Dst, err)
	serial.Read(r, &tx.Amt, err)
	serial.Read(r, &tx.ChainID, err)
	var b []byte
	serial.Read(r, &b, err)
	tx.Meta = b
	*i = tx
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
