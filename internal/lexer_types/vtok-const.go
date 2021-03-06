package lexer_types

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
	"math/big"
)

type Uint128 big.Int

func (v *Uint128) GetGVMType() gvm.RefType { return gvm.RefType(value_type.Uint128) }
func (v *Uint128) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Uint128) Encode() ([]byte, error) { return serial.EncodeBigInt((*big.Int)(v)) }
func (v *Uint128) Decode(b []byte) (gvm.Ref, error) {
	return (*Uint128)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Uint128) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v *Uint128) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v *Uint128) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }

type Uint256 big.Int

func (v *Uint256) GetGVMType() gvm.RefType { return gvm.RefType(value_type.Uint256) }
func (v *Uint256) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Uint256) Encode() ([]byte, error) { return serial.EncodeBigInt((*big.Int)(v)) }
func (v *Uint256) Decode(b []byte) (gvm.Ref, error) {
	return (*Uint256)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Uint256) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v *Uint256) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v *Uint256) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }

type Int128 big.Int

func (v *Int128) GetGVMType() gvm.RefType { return gvm.RefType(value_type.Int128) }
func (v *Int128) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Int128) Encode() ([]byte, error) { return serial.EncodeBigInt((*big.Int)(v)) }
func (v *Int128) Decode(b []byte) (gvm.Ref, error) {
	return (*Int128)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Int128) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v *Int128) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v *Int128) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }

type Int256 big.Int

func (v *Int256) GetGVMType() gvm.RefType { return gvm.RefType(value_type.Int256) }
func (v *Int256) Unwrap() interface{}     { return (*big.Int)(v) }
func (v *Int256) Encode() ([]byte, error) { return serial.EncodeBigInt((*big.Int)(v)) }
func (v *Int256) Decode(b []byte) (gvm.Ref, error) {
	return (*Int256)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Int256) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v *Int256) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v *Int256) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }

func (v *Int256) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.ReadBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Int256) Marshal(w io.Writer, err *error) { serial.WriteBigInt(w, (*big.Int)(v), err) }
func (v *Int128) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.ReadBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Int128) Marshal(w io.Writer, err *error) { serial.WriteBigInt(w, (*big.Int)(v), err) }
func (v *Uint128) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.ReadBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Uint128) Marshal(w io.Writer, err *error) { serial.WriteBigInt(w, (*big.Int)(v), err) }
func (v *Uint256) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.ReadBigInt(r, (*big.Int)(v), err)
	*pv = v
}
func (v *Uint256) Marshal(w io.Writer, err *error) { serial.WriteBigInt(w, (*big.Int)(v), err) }

type Uint8 uint8

func (v Uint8) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Uint8) }
func (v Uint8) Unwrap() interface{}                                     { return uint8(v) }
func (v Uint8) Encode() ([]byte, error)                                 { return []byte{uint8(v)}, nil }
func (v Uint8) Decode(b []byte) (gvm.Ref, error)                        { return Uint8(b[0]), nil }
func (v Uint8) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Uint8) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Uint8) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Uint8) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Uint8) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Uint16 uint16

func (v Uint16) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Uint16) }
func (v Uint16) Unwrap() interface{}                                     { return uint16(v) }
func (v Uint16) Encode() ([]byte, error)                                 { return util.Uint16ToBytes(uint16(v)), nil }
func (v Uint16) Decode(b []byte) (gvm.Ref, error)                        { return Uint16(util.BytesToUint16(b)), nil }
func (v Uint16) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Uint16) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Uint16) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Uint16) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Uint16) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Uint32 uint32

func (v Uint32) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Uint32) }
func (v Uint32) Unwrap() interface{}                                     { return uint32(v) }
func (v Uint32) Encode() ([]byte, error)                                 { return util.Uint32ToBytes(uint32(v)), nil }
func (v Uint32) Decode(b []byte) (gvm.Ref, error)                        { return Uint32(util.BytesToUint32(b)), nil }
func (v Uint32) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Uint32) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Uint32) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Uint32) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Uint32) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Uint64 uint64

func (v Uint64) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Uint64) }
func (v Uint64) Unwrap() interface{}                                     { return uint64(v) }
func (v Uint64) Encode() ([]byte, error)                                 { return util.Uint64ToBytes(uint64(v)), nil }
func (v Uint64) Decode(b []byte) (gvm.Ref, error)                        { return Uint64(util.BytesToUint64(b)), nil }
func (v Uint64) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Uint64) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Uint64) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Uint64) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Uint64) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Int8 int8

func (v Int8) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Int8) }
func (v Int8) Unwrap() interface{}                                     { return int8(v) }
func (v Int8) Encode() ([]byte, error)                                 { return []byte{uint8(v)}, nil }
func (v Int8) Decode(b []byte) (gvm.Ref, error)                        { return Int8(b[0]), nil }
func (v Int8) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Int8) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Int8) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Int8) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Int8) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Int16 int16

func (v Int16) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Int16) }
func (v Int16) Unwrap() interface{}                                     { return int16(v) }
func (v Int16) Encode() ([]byte, error)                                 { return util.Int16ToBytes(int16(v)), nil }
func (v Int16) Decode(b []byte) (gvm.Ref, error)                        { return Int16(util.BytesToInt16(b)), nil }
func (v Int16) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Int16) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Int16) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Int16) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Int16) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Int32 int32

func (v Int32) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Int32) }
func (v Int32) Unwrap() interface{}                                     { return int32(v) }
func (v Int32) Encode() ([]byte, error)                                 { return util.Int32ToBytes(int32(v)), nil }
func (v Int32) Decode(b []byte) (gvm.Ref, error)                        { return Int32(util.BytesToInt32(b)), nil }
func (v Int32) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Int32) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Int32) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Int32) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Int32) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Int64 int64

func (v Int64) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Int64) }
func (v Int64) Unwrap() interface{}                                     { return int64(v) }
func (v Int64) Encode() ([]byte, error)                                 { return util.Int64ToBytes(int64(v)), nil }
func (v Int64) Decode(b []byte) (gvm.Ref, error)                        { return Int64(util.BytesToInt64(b)), nil }
func (v Int64) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Int64) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Int64) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Int64) Marshal(w io.Writer, err *error)                         { serial.Write(w, v, err) }
func (v Int64) Unmarshal(r io.Reader, pv *uip.VTok, err *error)         { serial.Read(r, &v, err); *pv = v }

type Bytes []byte

func (v Bytes) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Bytes) }
func (v Bytes) Unwrap() interface{}                                     { return []byte(v) }
func (v Bytes) Encode() ([]byte, error)                                 { return v, nil }
func (v Bytes) Decode(b []byte) (gvm.Ref, error)                        { return Bytes(b), nil }
func (v Bytes) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Bytes) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Bytes) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v Bytes) Marshal(w io.Writer, err *error)                         { serial.Write(w, []byte(v), err) }
func (v Bytes) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.Read(r, (*[]byte)(&v), err)
	*pv = v
}

type String string

func (v String) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.String) }
func (v String) Unwrap() interface{}                                     { return string(v) }
func (v String) Encode() ([]byte, error)                                 { return []byte(v), nil }
func (v String) Decode(b []byte) (gvm.Ref, error)                        { return String(b), nil }
func (v String) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v String) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v String) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }
func (v String) Marshal(w io.Writer, err *error)                         { serial.Write(w, string(v), err) }
func (v String) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.Read(r, (*string)(&v), err)
	*pv = v
}

type Bool bool

func (v Bool) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Bool) }
func (v Bool) Unwrap() interface{}                                     { return bool(v) }
func (v Bool) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v Bool) Determine(_ DetermineContext) (token_types.Param, error) { return v, nil }
func (v Bool) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return v, nil }

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
func (v Bool) Marshal(w io.Writer, err *error) { serial.Write(w, bool(v), err) }
func (v Bool) Unmarshal(r io.Reader, pv *uip.VTok, err *error) {
	serial.Read(r, (*bool)(&v), err)
	*pv = v
}

var Undefined *Unknown

type Unknown struct{}

func (u *Unknown) GetGVMType() gvm.RefType                                 { return gvm.RefType(value_type.Unknown) }
func (u *Unknown) Unwrap() interface{}                                     { return nil }
func (u *Unknown) Encode() ([]byte, error)                                 { return nil, nil }
func (u *Unknown) Decode(_ []byte) (gvm.Ref, error)                        { return Undefined, nil }
func (u *Unknown) GetGVMTok() gvm.TokType                                  { return token_type.Constant }
func (v *Unknown) Determine(_ DetermineContext) (token_types.Param, error) { return Undefined, nil }
func (u *Unknown) Eval(_ *gvm.ExecCtx) (gvm.Ref, error)                    { return u, nil }
func (u *Unknown) Marshal(_ io.Writer, _ *error)                           {}
func (u *Unknown) Unmarshal(_ io.Reader, pv *uip.VTok, _ *error)           { *pv = Undefined }
