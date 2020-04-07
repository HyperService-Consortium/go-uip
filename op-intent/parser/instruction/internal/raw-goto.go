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

func (g RawGoto) Marshal(_ io.Writer, err *error) {
	if *err != nil {
		return
	}
	*err = ErrNotTranslated
}

func (g RawGoto) Exec(_ *gvm.ExecCtx) error {
	return ErrNotTranslated
}

func (g RawGoto) Unmarshal(_ io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	*err = ErrNotTranslated
	*i = g
}

func (r RawGoto) GetType() instruction_type.Type {
	return instruction_type.RawGoto
}
