package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type RawConditionGoto struct {
	IndexName string
	Condition lexer_types.Param
	Offset    int
}

func (g RawConditionGoto) Marshal(_ io.Writer, err *error) {
	if *err != nil {
		return
	}
	*err = errorn.NewRuntimeNotTranslated()
}

func (g RawConditionGoto) Exec(_ *gvm.ExecCtx) error {
	return errorn.NewRuntimeNotTranslated()
}

func (g RawConditionGoto) Unmarshal(_ io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	*err = errorn.NewRuntimeNotTranslated()
	*i = g
}

func (g RawConditionGoto) GetType() instruction_type.Type {
	return instruction_type.RawConditionGoto
}
