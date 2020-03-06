package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"math/big"
)

type BigInt big.Int

func (v *BigInt) Encode() ([]byte, error) {
	return (*big.Int)(v).Bytes(), nil
}

func (v *BigInt) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

type Uint128 big.Int

func (v *Uint128) GetGVMType() abstraction.RefType {
	return RefUint128
}

func (v *Uint128) Unwrap() interface{} {
	return (*big.Int)(v)
}

func (v *Uint128) Encode() ([]byte, error) {
	return (*big.Int)(v).Bytes(), nil
}

func (v *Uint128) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

type Uint256 big.Int

func (v *Uint256) GetGVMType() abstraction.RefType {
	return RefUint256
}

func (v *Uint256) Unwrap() interface{} {
	return (*big.Int)(v)
}

func (v *Uint256) Encode() ([]byte, error) {
	return (*big.Int)(v).Bytes(), nil
}

func (v *Uint256) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

type Int128 big.Int

func (v *Int128) GetGVMType() abstraction.RefType {
	return RefInt128
}

func (v *Int128) Unwrap() interface{} {
	return (*big.Int)(v)
}

func (v *Int128) Encode() ([]byte, error) {
	return (*big.Int)(v).Bytes(), nil
}

func (v *Int128) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

type Int256 big.Int

func (v *Int256) GetGVMType() abstraction.RefType {
	return RefInt256
}

func (v *Int256) Unwrap() interface{} {
	return (*big.Int)(v)
}

func (v *Int256) Encode() ([]byte, error) {
	return (*big.Int)(v).Bytes(), nil
}

func (v *Int256) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

type Uint8 uint8

func (v Uint8) GetGVMType() abstraction.RefType {
	return RefUint8
}

func (v Uint8) Unwrap() interface{} {
	return uint8(v)
}

func (v Uint8) Encode() ([]byte, error) {
	return []byte{uint8(v)}, nil
}

func (v *Uint8) Decode(b []byte) (interface{}, error) {
	return b[0], nil
}

type Uint16 uint16

func (v Uint16) GetGVMType() abstraction.RefType {
	return RefUint16
}

func (v Uint16) Unwrap() interface{} {
	return uint16(v)
}

func (v Uint16) Encode() ([]byte, error) {
	return util.Uint16ToBytes(uint16(v)), nil
}

func (v *Uint16) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint16(b), nil
}

type Uint32 uint32

func (v Uint32) GetGVMType() abstraction.RefType {
	return RefUint32
}

func (v Uint32) Unwrap() interface{} {
	return uint32(v)
}

func (v Uint32) Encode() ([]byte, error) {
	return util.Uint32ToBytes(uint32(v)), nil
}

func (v *Uint32) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint32(b), nil
}

type Uint64 uint64

func (v Uint64) GetGVMType() abstraction.RefType {
	return RefUint64
}

func (v Uint64) Unwrap() interface{} {
	return uint64(v)
}

func (v Uint64) Encode() ([]byte, error) {
	return util.Uint64ToBytes(uint64(v)), nil
}

func (v *Uint64) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint64(b), nil
}

type Int8 int8

func (v Int8) GetGVMType() abstraction.RefType {
	return RefInt8
}

func (v Int8) Unwrap() interface{} {
	return int8(v)
}

func (v Int8) Encode() ([]byte, error) {
	return []byte{uint8(int8(v))}, nil
}

func (v *Int8) Decode(b []byte) (interface{}, error) {
	return int8(b[0]), nil
}

func (v *Int8) Wrap(b interface{}) int8 {
	return b.(int8)
}

type Int16 int16

func (v Int16) GetGVMType() abstraction.RefType {
	return RefInt16
}

func (v Int16) Unwrap() interface{} {
	return int16(v)
}

func (v Int16) Encode() ([]byte, error) {
	return util.Int16ToBytes(int16(v)), nil
}

func (v *Int16) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt16(b), nil
}

func (v *Int16) Wrap(b interface{}) int16 {
	return b.(int16)
}

type Int32 int32

func (v Int32) GetGVMType() abstraction.RefType {
	return RefInt32
}

func (v Int32) Unwrap() interface{} {
	return int32(v)
}

func (v Int32) Encode() ([]byte, error) {
	return util.Int32ToBytes(int32(v)), nil
}

func (v *Int32) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt32(b), nil
}

func (v *Int32) Wrap(b interface{}) int32 {
	return b.(int32)
}

type Int64 int64

func (v Int64) GetGVMType() abstraction.RefType {
	return RefInt64
}

func (v Int64) Unwrap() interface{} {
	return int64(v)
}

func (v Int64) Encode() ([]byte, error) {
	return util.Int64ToBytes(int64(v)), nil
}

func (v *Int64) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt64(b), nil
}

func (v *Int64) Wrap(b interface{}) int64 {
	return b.(int64)
}

type Bytes []byte

func (v Bytes) GetGVMType() abstraction.RefType {
	return RefBytes
}

func (v Bytes) Unwrap() interface{} {
	return []byte(v)
}

func (v Bytes) Encode() ([]byte, error) {
	return v, nil
}

func (v *Bytes) Decode(b []byte) (interface{}, error) {
	return b, nil
}

type String string

func (v String) GetGVMType() abstraction.RefType {
	return RefString
}

func (v String) Unwrap() interface{} {
	return string(v)
}

func (v String) Encode() ([]byte, error) {
	return []byte(v), nil
}

func (v *String) Decode(b []byte) (interface{}, error) {
	return string(b), nil
}

type Bool bool

func (v Bool) GetGVMType() abstraction.RefType {
	return RefBool
}

func (v Bool) Unwrap() interface{} {
	return bool(v)
}

func (v Bool) Encode() ([]byte, error) {
	if v {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

func (v *Bool) Decode(b []byte) (interface{}, error) {
	if len(b) <= 0 {
		return false, nil
	}
	if b[0] != 0 {
		return true, nil
	} else {
		return false, nil
	}
}
