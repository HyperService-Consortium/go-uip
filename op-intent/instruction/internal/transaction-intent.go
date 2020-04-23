package internal

import (
	"bytes"
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	TxState "github.com/HyperService-Consortium/go-uip/const/transaction_state_type"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/op-intent/trap"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

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

func (tx *TransactionIntent) GetTxType() trans_type.Type {
	return tx.TransType
}

func (tx *TransactionIntent) GetChainID() uip.ChainIDUnderlyingType {
	return tx.ChainID
}

func (tx *TransactionIntent) GetSrc() []byte {
	return tx.Src
}

func (tx *TransactionIntent) GetDst() []byte {
	return tx.Dst
}

func (tx *TransactionIntent) GetMeta() []byte {
	return tx.Meta
}

func (tx *TransactionIntent) GetAmt() string {
	return tx.Amt
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

func (tx *TransactionIntent) Serialize() (_ []byte, err error) {
	var b = bytes.NewBuffer(nil)
	tx.Marshal(b, &err)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
