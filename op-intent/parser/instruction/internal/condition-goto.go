package internal

import (
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

type ConditionGoto struct {
	Type      instruction_type.Type `json:"itype"`
	Index     uint64                `json:"goto"`
	Condition json.RawMessage       `json:"condition"`
}

func (g *ConditionGoto) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g *ConditionGoto) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g *ConditionGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func NewConditionGoto(index uint64, condition json.RawMessage) *ConditionGoto {
	return &ConditionGoto{
		Type:      instruction_type.ConditionGoto,
		Index:     index,
		Condition: condition,
	}
}

func (g *ConditionGoto) Convert() (gg *GVMConditionGoto, err error) {
	gg = &GVMConditionGoto{
		Type:  g.Type,
		Index: g.Index,
	}
	gg.Condition, err = lexer.ParamUnmarshalJSON(g.Condition)
	return
}

type GVMConditionGoto struct {
	Type      instruction_type.Type `json:"itype"`
	Index     uint64                `json:"goto"`
	Condition gvm.VTok              `json:"condition"`
}

func (g GVMConditionGoto) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g GVMConditionGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (inst *GVMConditionGoto) Exec(g *gvm.ExecCtx) error {
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

func (g *ConditionGoto) GetType() instruction_type.Type {
	return instruction_type.ConditionGoto
}
