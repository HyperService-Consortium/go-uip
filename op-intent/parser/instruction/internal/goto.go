package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type Goto struct {
	Type  instruction_type.Type `json:"itype"`
	Index uint64                `json:"goto"`
}

func (g Goto) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g Goto) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g Goto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (tx *Goto) GetGotoIndexGVMI() uint64 {
	return tx.Index
}

func NewGoto(index uint64) *Goto {
	return &Goto{
		Type:  instruction_type.Goto,
		Index: index,
	}
}

func (tx *Goto) GetType() instruction_type.Type {
	return instruction_type.Goto
}
