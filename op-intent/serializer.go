package opintent

import (
	"bytes"
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

func EncodeVTok(w io.Writer, v uip.VTok, err *error) {
	lexer_types.EncodeVTok(w, v, err)
}

//noinspection GoUnusedExportedFunction
func EncodeVTokE(w io.Writer, v uip.VTok) (err error) {
	lexer_types.EncodeVTok(w, v, &err)
	return
}

func DecodeVTok(r io.Reader, v *uip.VTok, err *error) {
	lexer_types.DecodeVTok(r, v, err)
}

//noinspection GoUnusedExportedFunction
func DecodeVTokE(r io.Reader) (v uip.VTok, err error) {
	lexer_types.DecodeVTok(r, &v, &err)
	return
}

//noinspection GoUnusedExportedFunction
func EncodeAccount(w io.Writer, v Account, err *error) {
	lexer_types.EncodeAccount(w, v, err)
}

//noinspection GoUnusedExportedFunction
func EncodeAccountE(w io.Writer, v Account) (err error) {
	lexer_types.EncodeAccount(w, v, &err)
	return
}

//noinspection GoUnusedExportedFunction
func DecodeAccount(r io.Reader, v *Account, err *error) {
	lexer_types.DecodeAccount(r, v, err)
}

//noinspection GoUnusedExportedFunction
func DecodeAccountE(r io.Reader) (v Account, err error) {
	lexer_types.DecodeAccount(r, &v, &err)
	return
}

//noinspection GoUnusedExportedFunction
func EncodeInstruction(w io.Writer, i uip.Instruction) (err error) {
	return instruction.EncodeInstruction(i, w)
}

//noinspection GoUnusedExportedFunction
func DecodeInstruction(r io.Reader) (i uip.Instruction, err error) {
	return instruction.DecodeInstruction(r)
}

//noinspection GoUnusedExportedFunction
func DecodeInstructionWithType(r io.Reader, t instruction_type.Type) (i uip.Instruction, err error) {
	return instruction.DecodeInstructionWithType(r, t)
}

//noinspection GoUnusedExportedFunction
func DecodeVTokWithType(r io.Reader, v *uip.VTok, t gvm.TokType, err *error) {
	lexer_types.DecodeVTokWithType(r, v, t, err)
}

//noinspection GoUnusedExportedFunction
func ReadConstant(r io.Reader, v *uip.VTok, t gvm.RefType, err *error) {
	lexer_types.ReadConstant(r, v, t, err)
}

//noinspection GoUnusedExportedFunction
func EncodeInstructions(is []uip.Instruction) (bs [][]byte, err error) {
	return instruction.EncodeInstructions(is)
}

//noinspection GoUnusedExportedFunction
func DecodeInstructions(bs [][]byte) ([]uip.Instruction, error) {
	return instruction.DecodeInstructions(bs)

}

type mcs struct{}
type ts struct{}
type _serializer struct {
	TransactionIntent ts
	Meta              struct {
		Contract mcs
	}
}

//noinspection GoUnusedGlobalVariable
var Serializer = _serializer{}

func (mcs) Unmarshal(b []byte, meta *lexer.ContractInvokeMeta) (err error) {
	var r = bytes.NewReader(b)
	serial.Read(r, &meta.Code, &err)
	serial.Read(r, &meta.Meta, &err)
	serial.Read(r, &meta.FuncName, &err)
	var paramsLength uint64
	serial.Read(r, &paramsLength, &err)
	if err != nil {
		return
	}
	meta.Params = make([]uip.VTok, paramsLength)
	for i := range meta.Params {
		DecodeVTok(r, &meta.Params[i], &err)
	}
	return
}

func (mcs) Marshal(meta *lexer.ContractInvokeMeta) (_ []byte, err error) {
	var w = bytes.NewBuffer(nil)
	serial.Write(w, meta.Code, &err)
	serial.Write(w, meta.Meta, &err)
	serial.Write(w, meta.FuncName, &err)
	serial.Write(w, uint64(len(meta.Params)), &err)
	for i := range meta.Params {
		EncodeVTok(w, meta.Params[i], &err)
	}
	return w.Bytes(), err
}

func (ts) Unmarshal(b []byte, meta *TransactionIntent) error {
	return json.Unmarshal(b, meta)
}

func (ts) Marshal(meta *TransactionIntent) ([]byte, error) {
	return json.Marshal(meta)
}
