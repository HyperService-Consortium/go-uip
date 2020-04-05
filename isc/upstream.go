package isc

import (
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
	"math/big"
)

func execSetState(g *gvm.ExecCtx, target string, rhs gvm.VTok) error {
	k, err := rhs.Eval(g)
	if err != nil {
		return err
	}
	err = g.Save(target, k)
	if err != nil {
		return err
	}

	g.PC++
	return nil
}

func encodeBigInt(b *big.Int) ([]byte, error) {
	if b == nil {
		return nil, nil
	}
	return b.Bytes(), nil
}

func writeBigInt(w io.Writer, i *big.Int, err *error) {
	write(w, i.Bytes(), err)
}

func readBigInt(r io.Reader, i *big.Int, err *error) {
	var b []byte
	read(r, &b, err)
	i.SetBytes(b)
}

type Uint128 big.Int

func (v *Uint128) GetGVMType() gvm.RefType { return gvm_type.RefUint128 }
func (v *Uint128) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Uint128) Encode() ([]byte, error) { return encodeBigInt((*big.Int)(v)) }
func (v *Uint128) Decode(b []byte) (gvm.Ref, error) {
	return (*Uint128)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Uint128) GetGVMTok() gvm.TokType               { return gvm_type.TokConstant }
func (v *Uint128) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) { return v, nil }

type Uint256 big.Int

func (v *Uint256) GetGVMType() gvm.RefType { return gvm_type.RefUint256 }
func (v *Uint256) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Uint256) Encode() ([]byte, error) { return encodeBigInt((*big.Int)(v)) }
func (v *Uint256) Decode(b []byte) (gvm.Ref, error) {
	return (*Uint256)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Uint256) GetGVMTok() gvm.TokType               { return gvm_type.TokConstant }
func (v *Uint256) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) { return v, nil }

type Int128 big.Int

func (v *Int128) GetGVMType() gvm.RefType { return gvm_type.RefInt128 }
func (v *Int128) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Int128) Encode() ([]byte, error) { return encodeBigInt((*big.Int)(v)) }
func (v *Int128) Decode(b []byte) (gvm.Ref, error) {
	return (*Int128)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Int128) GetGVMTok() gvm.TokType               { return gvm_type.TokConstant }
func (v *Int128) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) { return v, nil }

type Int256 big.Int

func (v *Int256) GetGVMType() gvm.RefType { return gvm_type.RefInt256 }
func (v *Int256) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Int256) Encode() ([]byte, error) { return encodeBigInt((*big.Int)(v)) }
func (v *Int256) Decode(b []byte) (gvm.Ref, error) {
	return (*Int256)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Int256) GetGVMTok() gvm.TokType               { return gvm_type.TokConstant }
func (v *Int256) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) { return v, nil }

func (v *Int256) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	readBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Int256) Marshal(w io.Writer, err *error) { writeBigInt(w, (*big.Int)(v), err) }
func (v *Int128) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	readBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Int128) Marshal(w io.Writer, err *error) { writeBigInt(w, (*big.Int)(v), err) }
func (v *Uint128) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	readBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Uint128) Marshal(w io.Writer, err *error) { writeBigInt(w, (*big.Int)(v), err) }
func (v *Uint256) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	readBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Uint256) Marshal(w io.Writer, err *error) { writeBigInt(w, (*big.Int)(v), err) }

type Uint8 uint8

func (v Uint8) GetGVMType() gvm.RefType                         { return gvm_type.RefUint8 }
func (v Uint8) Unwrap() interface{}                             { return uint8(v) }
func (v Uint8) Encode() ([]byte, error)                         { return []byte{uint8(v)}, nil }
func (v Uint8) Decode(b []byte) (gvm.Ref, error)                { return Uint8(b[0]), nil }
func (v Uint8) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Uint8) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Uint8) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Uint8) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Uint16 uint16

func (v Uint16) GetGVMType() gvm.RefType                         { return gvm_type.RefUint16 }
func (v Uint16) Unwrap() interface{}                             { return uint16(v) }
func (v Uint16) Encode() ([]byte, error)                         { return util.Uint16ToBytes(uint16(v)), nil }
func (v Uint16) Decode(b []byte) (gvm.Ref, error)                { return Uint16(util.BytesToUint16(b)), nil }
func (v Uint16) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Uint16) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Uint16) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Uint16) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Uint32 uint32

func (v Uint32) GetGVMType() gvm.RefType                         { return gvm_type.RefUint32 }
func (v Uint32) Unwrap() interface{}                             { return uint32(v) }
func (v Uint32) Encode() ([]byte, error)                         { return util.Uint32ToBytes(uint32(v)), nil }
func (v Uint32) Decode(b []byte) (gvm.Ref, error)                { return Uint32(util.BytesToUint32(b)), nil }
func (v Uint32) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Uint32) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Uint32) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Uint32) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Uint64 uint64

func (v Uint64) GetGVMType() gvm.RefType                         { return gvm_type.RefUint64 }
func (v Uint64) Unwrap() interface{}                             { return uint64(v) }
func (v Uint64) Encode() ([]byte, error)                         { return util.Uint64ToBytes(uint64(v)), nil }
func (v Uint64) Decode(b []byte) (gvm.Ref, error)                { return Uint64(util.BytesToUint64(b)), nil }
func (v Uint64) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Uint64) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Uint64) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Uint64) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Int8 int8

func (v Int8) GetGVMType() gvm.RefType                         { return gvm_type.RefInt8 }
func (v Int8) Unwrap() interface{}                             { return int8(v) }
func (v Int8) Encode() ([]byte, error)                         { return []byte{uint8(v)}, nil }
func (v Int8) Decode(b []byte) (gvm.Ref, error)                { return Int8(b[0]), nil }
func (v Int8) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Int8) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Int8) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Int8) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Int16 int16

func (v Int16) GetGVMType() gvm.RefType                         { return gvm_type.RefInt16 }
func (v Int16) Unwrap() interface{}                             { return int16(v) }
func (v Int16) Encode() ([]byte, error)                         { return util.Int16ToBytes(int16(v)), nil }
func (v Int16) Decode(b []byte) (gvm.Ref, error)                { return Int16(util.BytesToInt16(b)), nil }
func (v Int16) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Int16) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Int16) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Int16) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Int32 int32

func (v Int32) GetGVMType() gvm.RefType                         { return gvm_type.RefInt32 }
func (v Int32) Unwrap() interface{}                             { return int32(v) }
func (v Int32) Encode() ([]byte, error)                         { return util.Int32ToBytes(int32(v)), nil }
func (v Int32) Decode(b []byte) (gvm.Ref, error)                { return Int32(util.BytesToInt32(b)), nil }
func (v Int32) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Int32) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Int32) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Int32) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Int64 int64

func (v Int64) GetGVMType() gvm.RefType                         { return gvm_type.RefInt64 }
func (v Int64) Unwrap() interface{}                             { return int64(v) }
func (v Int64) Encode() ([]byte, error)                         { return util.Int64ToBytes(int64(v)), nil }
func (v Int64) Decode(b []byte) (gvm.Ref, error)                { return Int64(util.BytesToInt64(b)), nil }
func (v Int64) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Int64) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Int64) Marshal(w io.Writer, err *error)                 { write(w, v, err) }
func (v Int64) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, &v, err); *pv = v }

type Bytes []byte

func (v Bytes) GetGVMType() gvm.RefType                         { return gvm_type.RefBytes }
func (v Bytes) Unwrap() interface{}                             { return []byte(v) }
func (v Bytes) Encode() ([]byte, error)                         { return v, nil }
func (v Bytes) Decode(b []byte) (gvm.Ref, error)                { return Bytes(b), nil }
func (v Bytes) GetGVMTok() gvm.TokType                          { return gvm_type.TokConstant }
func (v Bytes) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)            { return v, nil }
func (v Bytes) Marshal(w io.Writer, err *error)                 { write(w, []byte(v), err) }
func (v Bytes) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, (*[]byte)(&v), err); *pv = v }

type String string

func (v String) GetGVMType() gvm.RefType              { return gvm_type.RefString }
func (v String) Unwrap() interface{}                  { return string(v) }
func (v String) Encode() ([]byte, error)              { return []byte(v), nil }
func (v String) Decode(b []byte) (gvm.Ref, error)     { return String(b), nil }
func (v String) GetGVMTok() gvm.TokType               { return gvm_type.TokConstant }
func (v String) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) { return v, nil }
func (v String) Marshal(w io.Writer, err *error)      { write(w, string(v), err) }
func (v String) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	read(r, (*string)(&v), err)
	*pv = v
}

type Bool bool

func (v Bool) GetGVMType() gvm.RefType              { return gvm_type.RefBool }
func (v Bool) Unwrap() interface{}                  { return bool(v) }
func (v Bool) GetGVMTok() gvm.TokType               { return gvm_type.TokConstant }
func (v Bool) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) { return v, nil }

func (v Bool) Encode() ([]byte, error) {
	if v {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}
func (v Bool) Decode(b []byte) (gvm.Ref, error) {
	if len(b) <= 0 {
		return Bool(false), nil
	}
	if b[0] != 0 {
		return Bool(true), nil
	} else {
		return Bool(false), nil
	}
}
func (v Bool) Marshal(w io.Writer, err *error)                 { write(w, bool(v), err) }
func (v Bool) Unmarshal(r io.Reader, pv *uip.VTok, err *error) { read(r, (*bool)(&v), err); *pv = v }

var Undefined *Unknown

type Unknown struct{}

func (u *Unknown) GetGVMType() gvm.RefType                       { return gvm_type.RefUnknown }
func (u *Unknown) Unwrap() interface{}                           { return nil }
func (u *Unknown) Encode() ([]byte, error)                       { return nil, nil }
func (u *Unknown) Decode(_ []byte) (gvm.Ref, error)              { return Undefined, nil }
func (u *Unknown) GetGVMTok() gvm.TokType                        { return gvm_type.TokConstant }
func (u *Unknown) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)          { return u, nil }
func (u *Unknown) Marshal(_ io.Writer, _ *error)                 {}
func (u *Unknown) Unmarshal(_ io.Reader, pv *uip.VTok, _ *error) { *pv = Undefined }
