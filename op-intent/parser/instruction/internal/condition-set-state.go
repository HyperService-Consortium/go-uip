package internal

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

type ConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          string                `json:"target"`
	RightExpression uip.VTok              `json:"expression"`
	Condition       uip.VTok              `json:"condition"`
}

func (inst ConditionSetState) GetType() instruction_type.Type {
	return instruction_type.ConditionSetState
}

func (inst ConditionSetState) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, inst.Type, err)
	serial.Write(w, inst.Target, err)
	lexer.EncodeVTok(w, inst.RightExpression, err)
	lexer.EncodeVTok(w, inst.Condition, err)
}

func (inst ConditionSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &inst.Type, err)
	serial.Read(r, &inst.Target, err)
	lexer.DecodeVTok(r, &inst.RightExpression, err)
	lexer.DecodeVTok(r, &inst.Condition, err)
	*i = inst
}

func (inst ConditionSetState) Exec(g *gvm.ExecCtx) error {
	v, err := inst.Condition.Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != gvm_type.RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}

	if v.Unwrap().(bool) {
		k, err := inst.RightExpression.Eval(g)
		if err != nil {
			return err
		}
		err = g.Save(inst.Target, k)
		if err != nil {
			return err
		}
	}

	g.PC++
	return nil
}

func NewConditionSetState(t value_type.Type, target string, rhs, cond uip.VTok) *ConditionSetState {
	return &ConditionSetState{
		Type:            t,
		Target:          target,
		RightExpression: rhs,
		Condition:       cond,
	}
}
