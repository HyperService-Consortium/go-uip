package internal

import (
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

func (g RawConditionGoto) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (g RawConditionGoto) Exec(c *gvm.ExecCtx) error {
	panic("implement me")
}

func (g RawConditionGoto) Unmarshal(r io.Reader, i *uip.Instruction, err *error) {
	panic("implement me")
}

func (g RawConditionGoto) GetType() instruction_type.Type {
	return instruction_type.RawConditionGoto
}
