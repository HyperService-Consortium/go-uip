package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

func NewConditionGoto(index uint64, condition uip.VTok) *ConditionGoto {
	return &ConditionGoto{
		Index:     index,
		Condition: condition,
	}
}

type ConditionGoto struct {
	Index     uint64   `json:"goto"`
	Condition uip.VTok `json:"condition"`
}

func (inst ConditionGoto) GetType() instruction_type.Type {
	return instruction_type.ConditionGoto
}

func (inst ConditionGoto) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, inst.Index, err)
	lexer_types.EncodeVTok(w, inst.Condition, err)
}

func (inst ConditionGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &inst.Index, err)
	lexer_types.DecodeVTok(r, &inst.Condition, err)
	*i = inst
}

func (inst ConditionGoto) Exec(g *gvm.ExecCtx) error {
	v, err := inst.GetGotoConditionGVMI().Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != gvm_type.RefBool {
		return errorn.NewRuntimeTypeAssertionError(gvm_type.RefBool, v)
	}
	if v.Unwrap().(bool) {
		g.PC = inst.GetGotoIndexGVMI()
		return nil
	}
	g.PC++
	return nil
}

func (inst *ConditionGoto) GetGotoIndexGVMI() uint64 {
	return inst.Index
}

func (inst *ConditionGoto) GetGotoConditionGVMI() gvm.VTok {
	return inst.Condition
}
