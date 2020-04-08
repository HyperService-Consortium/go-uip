package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type Goto struct {
	Index uint64 `json:"goto"`
}

func (tx Goto) GetType() instruction_type.Type {
	return instruction_type.Goto
}

func (g Goto) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, g.Index, err)
}

func (g Goto) Exec(c *gvm.ExecCtx) error {
	c.PC = g.Index
	return nil
}

func (g Goto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &g.Index, err)
	*i = g
}

func (tx *Goto) GetGotoIndexGVMI() uint64 {
	return tx.Index
}

func NewGoto(index uint64) *Goto {
	return &Goto{
		Index: index,
	}
}
