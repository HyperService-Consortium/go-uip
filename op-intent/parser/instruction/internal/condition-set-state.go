package internal

import (
	"encoding/json"
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
	Type            value_type.Type `json:"value_type"`
	Target          []byte          `json:"target"`
	RightExpression json.RawMessage `json:"expression"`
	Condition       json.RawMessage `json:"condition"`
}

func (g ConditionSetState) GetType() instruction_type.Type {
	return instruction_type.ConditionSetState
}

func (g ConditionSetState) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, g.Type, err)
	serial.Write(w, g.Target, err)
	serial.Write(w, []byte(g.RightExpression), err)
	serial.Write(w, []byte(g.Condition), err)
}

func (g ConditionSetState) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g ConditionSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &g.Type, err)
	serial.Read(r, &g.Target, err)
	var b []byte
	serial.Read(r, &b, err)
	g.RightExpression = b
	serial.Read(r, &b, err)
	g.Condition = b
	*i = g
}

type GVMConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          string                `json:"target"`
	RightExpression uip.VTok              `json:"expression"`
	Condition       uip.VTok              `json:"condition"`
}

func (inst GVMConditionSetState) GetType() instruction_type.Type {
	return instruction_type.ConditionSetState
}

func (inst GVMConditionSetState) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, inst.Type, err)
	serial.Write(w, inst.Target, err)
	lexer.EncodeVTok(w, inst.RightExpression, err)
	lexer.EncodeVTok(w, inst.Condition, err)
}

func (inst GVMConditionSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &inst.Type, err)
	serial.Read(r, &inst.Target, err)
	lexer.DecodeVTok(r, &inst.RightExpression, err)
	lexer.DecodeVTok(r, &inst.Condition, err)
	*i = inst
}

func (g *ConditionSetState) Convert() (gg *GVMConditionSetState, err error) {
	gg = &GVMConditionSetState{
		Type:   g.Type,
		Target: string(g.Target),
	}
	panic("todo")
	//g.RightExpression, err = lexer.ParamUnmarshalJSON(tx.RightExpression)
	//if err != nil {
	//	return nil, err
	//}
	//g.Condition, err = lexer.ParamUnmarshalJSON(tx.Condition)
	return
}

func (inst GVMConditionSetState) Exec(g *gvm.ExecCtx) error {
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

func (inst GVMConditionSetState) GetGotoConditionGVMI() gvm.VTok {
	return inst.Condition
}

func (inst GVMConditionSetState) GetRefNameGVMI() string {
	return inst.Target
}

func (inst GVMConditionSetState) GetRightHandStatementGVMI() gvm.VTok {
	return inst.RightExpression
}

func NewConditionSetState(t value_type.Type, target []byte, rhs, cond json.RawMessage) *ConditionSetState {
	return &ConditionSetState{
		Type:            t,
		Target:          target,
		RightExpression: rhs,
		Condition:       cond,
	}
}
