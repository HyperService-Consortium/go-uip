package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type RawSetState struct {
	Target          *lexer_types.LocalStateVariable
	RightExpression lexer_types.Param
}

func (g RawSetState) Marshal(_ io.Writer, err *error) {
	if *err != nil {
		return
	}
	*err = ErrNotTranslated
}

func (g RawSetState) Exec(_ *gvm.ExecCtx) error {
	return ErrNotTranslated
}

func (g RawSetState) Unmarshal(_ io.Reader, i *uip.Instruction, err *error) {
	if *err != nil {
		return
	}
	*err = ErrNotTranslated
	*i = g
}

func (r RawSetState) GetType() instruction_type.Type {
	return instruction_type.RawSetState
}
