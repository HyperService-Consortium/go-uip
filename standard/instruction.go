package standard

import (
	"encoding/json"
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type _jsonX struct{}
func (*_jsonX) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
func (*_jsonX) Unmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}
func (*_jsonX) NewDocument(b []byte) (document.Document, error) {
	return document.NewGJSONDocument(b)
}

var jsonX *_jsonX
func IsValidInstructionImpl(i uip.Instruction) error {
	return IsValidInstructionImpl_(i, jsonX)
}

func IsValidInstructionType(t instruction_type.Type) error {
	if instruction_type.IsStandard(t) {
		return nil
	}
	return errors.New("not valid instruction type")
}

func IsValidInstructionImpl_(i uip.Instruction, marshaller interface{
	Marshal(v interface{}) ([]byte, error)
	NewDocument(b []byte) (document.Document, error)
}) error {
	x, err := marshaller.Marshal(i)
	if err != nil {
		return err
	}

	g, err := marshaller.NewDocument(x)
	if err != nil {
		return err
	}

	t := g.Get("itype")
	if !t.Exists() {
		return errors.New("type not found")
	}

	return nil
}
