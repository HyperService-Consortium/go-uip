package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type RawGoto struct {
	IndexName string
	Offset    int
}

func (g RawGoto) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g RawGoto) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g RawGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (r RawGoto) GetType() instruction_type.Type {
	return instruction_type.RawGoto
}
