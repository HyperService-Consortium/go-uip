package internal

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type SetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          []byte                `json:"target"`
	RightExpression json.RawMessage       `json:"expression"`
}

func (g SetState) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g SetState) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g SetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func NewSetState(t value_type.Type, target []byte, rhs json.RawMessage) *SetState {
	return &SetState{
		IType:           instruction_type.SetState,
		Type:            t,
		Target:          target,
		RightExpression: rhs,
	}
}

func (G *GVMSetState) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, G.Target, err)
	lexer.EncodeVTok(w, G.RightExpression, err)
}

func (G *GVMSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &G.Target, err)
	lexer.DecodeVTok(r, &G.RightExpression, err)
	*i = G
}

type GVMSetState struct {
	Type            value_type.Type `json:"value_type"`
	Target          string          `json:"target"`
	RightExpression uip.VTok        `json:"expression"`
}

func (G *GVMSetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}

func (G GVMSetState) Exec(g *gvm.ExecCtx) error {
	return execSetState(g, G.Target, G.RightExpression)
}

func execSetState(g *gvm.ExecCtx, target string, rhs gvm.VTok) error {
	k, err := rhs.Eval(g)
	if err != nil {
		return err
	}
	err = g.Save(target, k)
	if err != nil {
		return err
	}

	g.PC++
	return nil
}

func (G GVMSetState) GetRefNameGVMI() string {
	return G.Target
}

func (G GVMSetState) GetRightHandStatementGVMI() gvm.VTok {
	return G.RightExpression
}

func (tx *SetState) Convert() (g *GVMSetState, err error) {
	g = &GVMSetState{
		Type:   tx.Type,
		Target: string(tx.Target),
	}
	panic("todo")
	//g.RightExpression, err = lexer.ParamUnmarshalJSON(tx.RightExpression)
	return
}

func (tx *SetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}
