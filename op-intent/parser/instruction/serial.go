package instruction

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"io"
)

func EncodeInstructions(is []uip.Instruction) (bs [][]byte, err error) {
	bs = make([][]byte, len(is))
	var b = bytes.NewBuffer(nil)
	for i := range is {
		b.Reset()
		err = EncodeInstruction(is[i], b)
		if err != nil {
			return nil, err
		}
		bs[i] = make([]byte, b.Len())
		copy(bs[i], b.Bytes())
	}
	return bs, nil
}

func EncodeInstruction(i uip.Instruction, w io.Writer) (err error) {
	serial.Write(w, i.GetType(), &err)
	if err != nil {
		return err
	}
	i.Marshal(w, &err)
	return err
}

func DecodeInstruction(r io.Reader) (i uip.Instruction, err error) {
	var t instruction_type.Type
	err = binary.Read(r, binary.BigEndian, &t)
	if err != nil {
		return nil, err
	}
	return DecodeInstructionWithType(r, t)
}

func DecodeInstructionWithType(r io.Reader, t instruction_type.Type) (i uip.Instruction, err error) {
	switch t {
	case instruction_type.SetState:
		i = new(SetState)
	case instruction_type.ConditionSetState:
		i = new(ConditionSetState)
	case instruction_type.Goto:
		i = new(Goto)
	case instruction_type.ConditionGoto:
		i = new(ConditionGoto)
	case instruction_type.ContractInvoke, instruction_type.Payment:
		i = new(TransactionIntent)
	//case instruction_type.RawSetState:
	//	fallthrough
	//case instruction_type.RawGoto:
	//	fallthrough
	default:
		return nil, errors.New("invalid type")
	}
	i.Unmarshal(r, &i, &err)
	return
}
