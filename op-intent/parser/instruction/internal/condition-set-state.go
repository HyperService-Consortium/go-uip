package internal

import (
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

type ConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          []byte                `json:"target"`
	RightExpression json.RawMessage       `json:"expression"`
	Condition       json.RawMessage       `json:"condition"`
}

func (g ConditionSetState) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g ConditionSetState) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g ConditionSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

type GVMConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          string                `json:"target"`
	RightExpression gvm.VTok              `json:"expression"`
	Condition       gvm.VTok              `json:"condition"`
}

func (g GVMConditionSetState) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g GVMConditionSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (tx *ConditionSetState) Convert() (g *GVMConditionSetState, err error) {
	g = &GVMConditionSetState{
		IType:  tx.IType,
		Type:   tx.Type,
		Target: string(tx.Target),
	}
	g.RightExpression, err = lexer.ParamUnmarshalJSON(tx.RightExpression)
	if err != nil {
		return nil, err
	}
	g.Condition, err = lexer.ParamUnmarshalJSON(tx.Condition)
	return
}

func (inst *GVMConditionSetState) Exec(g *gvm.ExecCtx) error {
	v, err := inst.GetGotoConditionGVMI().Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != gvm_type.RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}

	if v.Unwrap().(bool) {
		k, err := inst.GetRightHandStatementGVMI().Eval(g)
		if err != nil {
			return err
		}
		err = g.Save(inst.GetRefNameGVMI(), k)
		if err != nil {
			return err
		}
	}

	g.PC++
	return nil
}

func (G GVMConditionSetState) GetGotoConditionGVMI() gvm.VTok {
	return G.Condition
}

func (G GVMConditionSetState) GetRefNameGVMI() string {
	return G.Target
}

func (G GVMConditionSetState) GetRightHandStatementGVMI() gvm.VTok {
	return G.RightExpression
}

func NewConditionSetState(t value_type.Type, target []byte, rhs, cond json.RawMessage) *ConditionSetState {
	return &ConditionSetState{
		IType:           instruction_type.ConditionSetState,
		Type:            t,
		Target:          target,
		RightExpression: rhs,
		Condition:       cond,
	}
}

func (tx *ConditionSetState) GetType() instruction_type.Type {
	return instruction_type.ConditionSetState
}
