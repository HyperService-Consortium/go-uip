package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

func NewSetState(t value_type.Type, target string, rhs uip.VTok) *SetState {
	return &SetState{
		Type:            t,
		Target:          target,
		RightExpression: rhs,
	}
}

func (G SetState) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, G.Type, err)
	serial.Write(w, G.Target, err)
	lexer_types.EncodeVTok(w, G.RightExpression, err)
}

func (G SetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &G.Type, err)
	serial.Read(r, &G.Target, err)
	lexer_types.DecodeVTok(r, &G.RightExpression, err)
	*i = G
}

type SetState struct {
	Type            value_type.Type `json:"value_type"`
	Target          string          `json:"target"`
	RightExpression uip.VTok        `json:"expression"`
}

func (G SetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}

func (G SetState) Exec(g *gvm.ExecCtx) error {
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

func (G SetState) GetRefNameGVMI() string {
	return G.Target
}

func (G SetState) GetRightHandStatementGVMI() gvm.VTok {
	return G.RightExpression
}
