package internal

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
	"math/big"
)

func EncodeVTok(w io.Writer, v uip.VTok, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, v.GetGVMTok(), err)
	if v.GetGVMTok() == token.Constant {
		serial.Write(w, v.GetGVMType(), err)
	}
	v.Marshal(w, err)
	return
}

func DecodeVTok(r io.Reader, v *uip.VTok, err *error) {
	if *err != nil {
		return
	}
	var t gvm.TokType
	*err = binary.Read(r, binary.BigEndian, &t)
	DecodeVTokWithType(r, v, t, err)
}

func EncodeAccount(w io.Writer, v Account, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, v.GetType(), err)
	v.Marshal(w, err)
}

func DecodeAccount(r io.Reader, v *Account, err *error) {
	if *err != nil {
		return
	}
	var t gvm.TokType
	*err = binary.Read(r, binary.BigEndian, &t)
	DecodeAccountWithType(r, v, t, err)
}

func DecodeAccountWithType(r io.Reader, v *Account, t gvm.TokType, err *error) {
	if *err != nil {
		return
	}
	switch t {
	case token.NamespacedNameAccount:
		*v = new(NamespacedNameAccount)
	case token.NameAccount:
		*v = new(NameAccount)
	case token.NamespacedRawAccount:
		*v = new(NamespacedRawAccount)
	case token.RawAccount:
		*v = new(RawAccount)
	default:
		*err = errors.New("bad account type")
		return
	}

	(*v).Unmarshal(r, v, err)
}

func DecodeVTokWithType(r io.Reader, v *uip.VTok, t gvm.TokType, err *error) {
	if *err != nil {
		return
	}
	switch t {
	case token.Constant:
		var vt gvm.RefType
		serial.Read(r, &vt, err)
		if *err != nil {
			return
		}
		ReadConstant(r, v, vt, err)
		return
	case token.BinaryExpression:
		*v = new(DeterminedBinaryExpression)
	case token.UnaryExpression:
		*v = new(DeterminedUnaryExpression)
	case token.LocalStateVariable:
		*v = new(LocalStateVariable)
	case token.StateVariable:
		*v = new(StateVariable)
	default:
		panic("todo")
	}

	(*v).Unmarshal(r, v, err)
	return
}

func ReadConstant(r io.Reader, v *uip.VTok, t gvm.RefType, err *error) {
	switch t {
	case gvm_type.RefUint8:
		*v = Uint8(0)
	case gvm_type.RefUint16:
		*v = Uint16(0)
	case gvm_type.RefUint32:
		*v = Uint32(0)
	case gvm_type.RefUint64:
		*v = Uint64(0)

	case gvm_type.RefInt8:
		*v = Int8(0)
	case gvm_type.RefInt16:
		*v = Int16(0)
	case gvm_type.RefInt32:
		*v = Int32(0)
	case gvm_type.RefInt64:
		*v = Int64(0)

	case gvm_type.RefUint128:
		*v = (*Uint128)(big.NewInt(0))
	case gvm_type.RefUint256:
		*v = (*Uint256)(big.NewInt(0))
	case gvm_type.RefInt128:
		*v = (*Int128)(big.NewInt(0))
	case gvm_type.RefInt256:
		*v = (*Int256)(big.NewInt(0))

	case gvm_type.RefString:
		*v = String("")
	case gvm_type.RefBytes:
		*v = Bytes(nil)
	case gvm_type.RefBool:
		*v = Bool(false)
	case gvm_type.RefUnknown:
		*v = Undefined
	default:
		panic(fmt.Errorf("unknown reference type: %v", gvm_type.ExplainGVMType(t)))
	}
	(*v).Unmarshal(r, v, err)
}
