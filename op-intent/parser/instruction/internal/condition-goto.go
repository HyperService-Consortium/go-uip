package internal

import (
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

type ConditionGoto struct {
	Index     uint64          `json:"goto"`
	Condition json.RawMessage `json:"condition"`
}

func (g ConditionGoto) GetType() instruction_type.Type {
	return instruction_type.ConditionGoto
}

func (g ConditionGoto) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, g.Index, err)
	serial.Write(w, []byte(g.Condition), err)
}

func (g ConditionGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &g.Index, err)
	var b []byte
	serial.Read(r, &b, err)
	g.Condition = b
	*i = g
}

func (g ConditionGoto) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func NewConditionGoto(index uint64, condition json.RawMessage) *ConditionGoto {
	return &ConditionGoto{
		Index:     index,
		Condition: condition,
	}
}

func (g *ConditionGoto) Convert() (gg *GVMConditionGoto, err error) {
	gg = &GVMConditionGoto{
		Index: g.Index,
	}
	panic("todo")
	//gg.Condition, err = lexer.ParamUnmarshalJSON(g.Condition)
	return
}

type GVMConditionGoto struct {
	Type      instruction_type.Type `json:"itype"`
	Index     uint64                `json:"goto"`
	Condition uip.VTok              `json:"condition"`
}

func (inst GVMConditionGoto) GetType() instruction_type.Type {
	return instruction_type.ConditionGoto
}

func (inst GVMConditionGoto) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, inst.Index, err)
	lexer.EncodeVTok(w, inst.Condition, err)
}

func (inst GVMConditionGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &inst.Index, err)
	lexer.DecodeVTok(r, &inst.Condition, err)
	*i = inst
}

func (inst GVMConditionGoto) Exec(g *gvm.ExecCtx) error {
	v, err := inst.GetGotoConditionGVMI().Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != gvm_type.RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}
	if v.Unwrap().(bool) {
		g.PC = uint64(inst.GetGotoIndexGVMI())
		return nil
	}
	g.PC++
	return nil
}

func (inst *GVMConditionGoto) GetGotoIndexGVMI() uint64 {
	return inst.Index
}

func (inst *GVMConditionGoto) GetGotoConditionGVMI() gvm.VTok {
	return inst.Condition
}
