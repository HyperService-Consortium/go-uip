package opintent

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

func EncodeVTok(w io.Writer, v uip.VTok, err *error) {
	lexer_types.EncodeVTok(w, v, err)
}

func DecodeVTok(r io.Reader, v *uip.VTok, err *error) {
	lexer_types.DecodeVTok(r, v, err)
}

func EncodeAccount(w io.Writer, v Account, err *error) {
	lexer_types.EncodeAccount(w, v, err)
}

func DecodeAccount(r io.Reader, v *Account, err *error) {
	lexer_types.DecodeAccount(r, v, err)
}

func DecodeVTokWithType(r io.Reader, v *uip.VTok, t gvm.TokType, err *error) {
	lexer_types.DecodeVTokWithType(r, v, t, err)
}

func ReadConstant(r io.Reader, v *uip.VTok, t gvm.RefType, err *error) {
	lexer_types.ReadConstant(r, v, t, err)
}

func EncodeInstructions(is []uip.Instruction) (bs [][]byte, err error) {
	return instruction.EncodeInstructions(is)
}

func DecodeInstructions(bs [][]byte) ([]uip.Instruction, error) {
	return instruction.DecodeInstructions(bs)

}

func EncodeInstruction(w io.Writer, i uip.Instruction) (err error) {
	return instruction.EncodeInstruction(i, w)
}

func DecodeInstruction(r io.Reader) (i uip.Instruction, err error) {
	return instruction.DecodeInstruction(r)
}

func DecodeInstructionWithType(r io.Reader, t instruction_type.Type) (i uip.Instruction, err error) {
	return instruction.DecodeInstructionWithType(r, t)
}

type mcs struct{}
type ts struct{}
type _serializer struct {
	TransactionIntent ts
	Meta              struct {
		Contract mcs
	}
}

var Serializer = _serializer{}

func (mcs) Unmarshal(b []byte, meta *lexer.ContractInvokeMeta) error {
	return json.Unmarshal(b, meta)
}

func (mcs) Marshal(meta *lexer.ContractInvokeMeta) ([]byte, error) {
	return json.Marshal(meta)
}

func (ts) Unmarshal(b []byte, meta *TransactionIntent) error {
	return json.Unmarshal(b, meta)
}

func (ts) Marshal(meta *TransactionIntent) ([]byte, error) {
	return json.Marshal(meta)
}
