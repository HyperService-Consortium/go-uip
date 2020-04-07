package internal

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type RawConditionGoto struct {
	IndexName string
	Condition lexer.Param
	Offset    int
}

var ErrNotTranslated = errors.New("not translated")

func (g RawConditionGoto) Marshal(_ io.Writer, err *error) {
	if *err != nil {
		return
	}
	*err = ErrNotTranslated
}

func (g RawConditionGoto) Exec(_ *gvm.ExecCtx) error {
	return ErrNotTranslated
}

func (g RawConditionGoto) Unmarshal(_ io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	*err = ErrNotTranslated
	*i = g
}

func (g RawConditionGoto) GetType() instruction_type.Type {
	return instruction_type.RawConditionGoto
}
