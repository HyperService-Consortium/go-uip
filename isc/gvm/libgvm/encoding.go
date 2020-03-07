package libgvm

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"math/big"
)

func CreateRef(t abstraction.RefType, v interface{}) abstraction.Ref {
	switch t {
	case RefUint8:
		return Uint8(v.(uint8))
	case RefUint16:
		return Uint16(v.(uint16))
	case RefUint32:
		return Uint32(v.(uint32))
	case RefUint64:
		return Uint64(v.(uint64))
	case RefInt8:
		return Int8(v.(int8))
	case RefInt16:
		return Int16(v.(int16))
	case RefInt32:
		return Int32(v.(int32))
	case RefInt64:
		return Int64(v.(int64))
	case RefUint128:
		return (*Uint128)(v.(*big.Int))
	case RefUint256:
		return (*Uint256)(v.(*big.Int))
	case RefInt128:
		return (*Int128)(v.(*big.Int))
	case RefInt256:
		return (*Int256)(v.(*big.Int))
	case RefBool:
		return Bool(v.(bool))
	case RefString:
		return String(v.(string))
	case RefBytes:
		return Bytes(v.([]byte))
	case RefUnknown:
		return Undefined
	}
	panic(fmt.Errorf("unknown reference type: %v", ExplainGVMType(t)))
}

func DecodeRef(t abstraction.RefType, r []byte) (abstraction.Ref, error) {
	switch t {
	case RefUint8:
		return Uint8(0).Decode(r)
	case RefUint16:
		return Uint16(0).Decode(r)
	case RefUint32:
		return Uint32(0).Decode(r)
	case RefUint64:
		return Uint64(0).Decode(r)

	case RefInt8:
		return Int8(0).Decode(r)
	case RefInt16:
		return Int16(0).Decode(r)
	case RefInt32:
		return Int32(0).Decode(r)
	case RefInt64:
		return Int64(0).Decode(r)

	case RefUint128:
		return (*Uint128)(nil).Decode(r)
	case RefUint256:
		return (*Uint256)(nil).Decode(r)
	case RefInt128:
		return (*Int128)(nil).Decode(r)
	case RefInt256:
		return (*Int256)(nil).Decode(r)

	case RefString:
		return String("").Decode(r)
	case RefBytes:
		return Bytes(nil).Decode(r)
	case RefBool:
		return Bool(false).Decode(r)
	case RefUnknown:
		return Undefined, nil
	}
	panic(fmt.Errorf("unknown reference type: %v", ExplainGVMType(t)))
}

var Undefined *Unknown

type Unknown struct{}

func (u *Unknown) GetGVMType() abstraction.RefType                      { return RefUnknown }
func (u *Unknown) Unwrap() interface{}                                  { return nil }
func (u *Unknown) Encode() ([]byte, error)                              { return nil, nil }
func (u *Unknown) Decode(b []byte) (abstraction.Ref, error)             { return Undefined, nil }
func (u *Unknown) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (u *Unknown) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return u, nil }

type Uint128 big.Int

func (v *Uint128) GetGVMType() abstraction.RefType { return RefUint128 }
func (v *Uint128) Unwrap() interface{}             { return (*big.Int)(v) }
func (v *Uint128) Encode() ([]byte, error)         { return (*big.Int)(v).Bytes(), nil }
func (v *Uint128) Decode(b []byte) (abstraction.Ref, error) {
	return (*Uint128)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Uint128) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v *Uint128) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Uint256 big.Int

func (v *Uint256) GetGVMType() abstraction.RefType { return RefUint256 }
func (v *Uint256) Unwrap() interface{}             { return (*big.Int)(v) }
func (v *Uint256) Encode() ([]byte, error)         { return (*big.Int)(v).Bytes(), nil }
func (v *Uint256) Decode(b []byte) (abstraction.Ref, error) {
	return (*Uint256)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Uint256) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v *Uint256) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Int128 big.Int

func (v *Int128) GetGVMType() abstraction.RefType { return RefInt128 }
func (v *Int128) Unwrap() interface{}             { return (*big.Int)(v) }
func (v *Int128) Encode() ([]byte, error)         { return (*big.Int)(v).Bytes(), nil }
func (v *Int128) Decode(b []byte) (abstraction.Ref, error) {
	return (*Int128)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Int128) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v *Int128) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Int256 big.Int

func (v *Int256) GetGVMType() abstraction.RefType { return RefInt256 }
func (v *Int256) Unwrap() interface{}             { return (*big.Int)(v) }
func (v *Int256) Encode() ([]byte, error)         { return (*big.Int)(v).Bytes(), nil }
func (v *Int256) Decode(b []byte) (abstraction.Ref, error) {
	return (*Int256)(big.NewInt(0).SetBytes(b)), nil
}
func (v *Int256) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v *Int256) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Uint8 uint8

func (v Uint8) GetGVMType() abstraction.RefType                      { return RefUint8 }
func (v Uint8) Unwrap() interface{}                                  { return uint8(v) }
func (v Uint8) Encode() ([]byte, error)                              { return []byte{uint8(v)}, nil }
func (v Uint8) Decode(b []byte) (abstraction.Ref, error)             { return Uint8(b[0]), nil }
func (v Uint8) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Uint8) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Uint16 uint16

func (v Uint16) GetGVMType() abstraction.RefType                      { return RefUint16 }
func (v Uint16) Unwrap() interface{}                                  { return uint16(v) }
func (v Uint16) Encode() ([]byte, error)                              { return util.Uint16ToBytes(uint16(v)), nil }
func (v Uint16) Decode(b []byte) (abstraction.Ref, error)             { return Uint16(util.BytesToUint16(b)), nil }
func (v Uint16) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Uint16) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Uint32 uint32

func (v Uint32) GetGVMType() abstraction.RefType                      { return RefUint32 }
func (v Uint32) Unwrap() interface{}                                  { return uint32(v) }
func (v Uint32) Encode() ([]byte, error)                              { return util.Uint32ToBytes(uint32(v)), nil }
func (v Uint32) Decode(b []byte) (abstraction.Ref, error)             { return Uint32(util.BytesToUint32(b)), nil }
func (v Uint32) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Uint32) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Uint64 uint64

func (v Uint64) GetGVMType() abstraction.RefType                      { return RefUint64 }
func (v Uint64) Unwrap() interface{}                                  { return uint64(v) }
func (v Uint64) Encode() ([]byte, error)                              { return util.Uint64ToBytes(uint64(v)), nil }
func (v Uint64) Decode(b []byte) (abstraction.Ref, error)             { return Uint64(util.BytesToUint64(b)), nil }
func (v Uint64) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Uint64) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Int8 int8

func (v Int8) GetGVMType() abstraction.RefType                      { return RefInt8 }
func (v Int8) Unwrap() interface{}                                  { return int8(v) }
func (v Int8) Encode() ([]byte, error)                              { return []byte{uint8(int8(v))}, nil }
func (v Int8) Decode(b []byte) (abstraction.Ref, error)             { return Int8(b[0]), nil }
func (v Int8) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Int8) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

func (v *Int8) Wrap(b interface{}) int8 { return b.(int8) }

type Int16 int16

func (v Int16) GetGVMType() abstraction.RefType                      { return RefInt16 }
func (v Int16) Unwrap() interface{}                                  { return int16(v) }
func (v Int16) Encode() ([]byte, error)                              { return util.Int16ToBytes(int16(v)), nil }
func (v Int16) Decode(b []byte) (abstraction.Ref, error)             { return Int16(util.BytesToInt16(b)), nil }
func (v Int16) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Int16) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

func (v *Int16) Wrap(b interface{}) int16 { return b.(int16) }

type Int32 int32

func (v Int32) GetGVMType() abstraction.RefType                      { return RefInt32 }
func (v Int32) Unwrap() interface{}                                  { return int32(v) }
func (v Int32) Encode() ([]byte, error)                              { return util.Int32ToBytes(int32(v)), nil }
func (v Int32) Decode(b []byte) (abstraction.Ref, error)             { return Int32(util.BytesToInt32(b)), nil }
func (v Int32) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Int32) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

func (v *Int32) Wrap(b interface{}) int32 { return b.(int32) }

type Int64 int64

func (v Int64) GetGVMType() abstraction.RefType                      { return RefInt64 }
func (v Int64) Unwrap() interface{}                                  { return int64(v) }
func (v Int64) Encode() ([]byte, error)                              { return util.Int64ToBytes(int64(v)), nil }
func (v Int64) Decode(b []byte) (abstraction.Ref, error)             { return Int64(util.BytesToInt64(b)), nil }
func (v Int64) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Int64) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

func (v *Int64) Wrap(b interface{}) int64 { return b.(int64) }

type Bytes []byte

func (v Bytes) GetGVMType() abstraction.RefType                      { return RefBytes }
func (v Bytes) Unwrap() interface{}                                  { return []byte(v) }
func (v Bytes) Encode() ([]byte, error)                              { return v, nil }
func (v Bytes) Decode(b []byte) (abstraction.Ref, error)             { return Bytes(b), nil }
func (v Bytes) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Bytes) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type String string

func (v String) GetGVMType() abstraction.RefType                      { return RefString }
func (v String) Unwrap() interface{}                                  { return string(v) }
func (v String) Encode() ([]byte, error)                              { return []byte(v), nil }
func (v String) Decode(b []byte) (abstraction.Ref, error)             { return String(b), nil }
func (v String) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v String) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

type Bool bool

func (v Bool) GetGVMType() abstraction.RefType                      { return RefBool }
func (v Bool) Unwrap() interface{}                                  { return bool(v) }
func (v Bool) GetGVMTok() abstraction.TokType                       { return TokConstant }
func (v Bool) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) { return v, nil }

func (v Bool) Encode() ([]byte, error) {
	if v {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}
func (v Bool) Decode(b []byte) (abstraction.Ref, error) {
	if len(b) <= 0 {
		return Bool(false), nil
	}
	if b[0] != 0 {
		return Bool(true), nil
	} else {
		return Bool(false), nil
	}
}
