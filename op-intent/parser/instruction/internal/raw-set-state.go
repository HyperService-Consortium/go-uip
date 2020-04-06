package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type RawSetState struct {
	Target          *lexer.LocalStateVariable
	RightExpression lexer.Param
}

func (g RawSetState) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g RawSetState) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g RawSetState) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (r RawSetState) GetType() instruction_type.Type {
	return instruction_type.RawSetState
}
