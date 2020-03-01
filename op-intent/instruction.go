package opintent

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type LazyInstruction interface {
	uip.Instruction
	Deserialize() (uip.Instruction, error)
}

type lazyInstructionImpl struct {
	g document.Document
	Type instruction_type.Type
}

func (l lazyInstructionImpl) GetType() instruction_type.Type {
	return l.Type
}

func (l lazyInstructionImpl) Deserialize() (uip.Instruction, error) {
	panic("implement me")
}

func HandleInstruction(b []byte) (LazyInstruction, error) {

	g, err := document.NewGJSONDocument(b)
	if err != nil {
		return nil, err
	}

	t := g.Get("itype")
	if !t.Exists() {
		return nil, errors.New("itype not exists")
	}

	return &lazyInstructionImpl{
		g:g,
		Type: t.Uint(),
	}, nil
}
