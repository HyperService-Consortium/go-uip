package gvm

import (
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"math/big"
)

var (
	Uint8  = (*_Uint8)(nil)
	Uint16 = (*_Uint16)(nil)
	Uint32 = (*_Uint32)(nil)
	Uint64 = (*_Uint64)(nil)

	Int8  = (*_Int8)(nil)
	Int16 = (*_Int16)(nil)
	Int32 = (*_Int32)(nil)
	Int64 = (*_Int64)(nil)

	XUint = Uint64
	XInt  = Int64

	Uint128 = (*_BigInt)(nil)
	Uint256 = (*_BigInt)(nil)
	Int128  = (*_BigInt)(nil)
	Int256  = (*_BigInt)(nil)

	Bytes  = (*_Bytes)(nil)
	String = (*_String)(nil)
	Bool   = (*_Bool)(nil)

	PObjs = [...]PackPrototype{
		RefBool:    Bool,
		RefBytes:   Bytes,
		RefString:  String,
		RefUint8:   Uint8,
		RefUint16:  Uint16,
		RefUint32:  Uint32,
		RefUint64:  Uint64,
		RefUint128: Uint128,
		RefUint256: Uint256,
		RefInt8:    Int8,
		RefInt16:   Int16,
		RefInt32:   Int32,
		RefInt64:   Int64,
		RefInt128:  Int128,
		RefInt256:  Int256,
	}
)

func (_ *_BigInt) Encode(v interface{}) ([]byte, error) {
	return v.(*big.Int).Bytes(), nil
}

type _BigInt struct{}

func (_ *_BigInt) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

func (_ *_BigInt) Wrap(b interface{}) *big.Int {
	return b.(*big.Int)
}

type _Uint8 struct{}

func (_ *_Uint8) Encode(v interface{}) ([]byte, error) {
	return []byte{v.(uint8)}, nil
}

func (_ *_Uint8) Decode(b []byte) (interface{}, error) {
	return b[0], nil
}

func (_ *_Uint8) Wrap(b interface{}) uint8 {
	return b.(uint8)
}

type _Uint16 struct{}

func (_ *_Uint16) Encode(v interface{}) ([]byte, error) {
	return util.Uint16ToBytes(v.(uint16)), nil
}

func (_ *_Uint16) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint16(b), nil
}

func (_ *_Uint16) Wrap(b interface{}) uint16 {
	return b.(uint16)
}

type _Uint32 struct{}

func (_ *_Uint32) Encode(v interface{}) ([]byte, error) {
	return util.Uint32ToBytes(v.(uint32)), nil
}

func (_ *_Uint32) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint32(b), nil
}

func (_ *_Uint32) Wrap(b interface{}) uint32 {
	return b.(uint32)
}

type _Uint64 struct{}

func (_ *_Uint64) Encode(v interface{}) ([]byte, error) {
	return util.Uint64ToBytes(v.(uint64)), nil
}

func (_ *_Uint64) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint64(b), nil
}

func (_ *_Uint64) Wrap(b interface{}) uint64 {
	return b.(uint64)
}

type _Int8 struct{}

func (_ *_Int8) Encode(v interface{}) ([]byte, error) {
	return []byte{uint8(v.(int8))}, nil
}

func (_ *_Int8) Decode(b []byte) (interface{}, error) {
	return int8(b[0]), nil
}

func (_ *_Int8) Wrap(b interface{}) int8 {
	return b.(int8)
}

type _Int16 struct{}

func (_ *_Int16) Encode(v interface{}) ([]byte, error) {
	return util.Int16ToBytes(v.(int16)), nil
}

func (_ *_Int16) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt16(b), nil
}

func (_ *_Int16) Wrap(b interface{}) int16 {
	return b.(int16)
}

type _Int32 struct{}

func (_ *_Int32) Encode(v interface{}) ([]byte, error) {
	return util.Int32ToBytes(v.(int32)), nil
}

func (_ *_Int32) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt32(b), nil
}

func (_ *_Int32) Wrap(b interface{}) int32 {
	return b.(int32)
}

type _Int64 struct{}

func (_ *_Int64) Encode(v interface{}) ([]byte, error) {
	return util.Int64ToBytes(v.(int64)), nil
}

func (_ *_Int64) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt64(b), nil
}

func (_ *_Int64) Wrap(b interface{}) int64 {
	return b.(int64)
}

type _Bytes struct{}

func (_ *_Bytes) Encode(b interface{}) ([]byte, error) {
	return b.([]byte), nil
}

func (_ *_Bytes) Decode(b []byte) (interface{}, error) {
	return b, nil
}

func (_ *_Bytes) Wrap(b interface{}) []byte {
	return b.([]byte)
}

type _String struct{}

func (_ *_String) Encode(b interface{}) ([]byte, error) {
	return []byte(b.(string)), nil
}

func (_ *_String) Decode(b []byte) (interface{}, error) {
	return string(b), nil
}

func (_ *_String) Wrap(b interface{}) string {
	return b.(string)
}

type _Bool struct{}

func (_ *_Bool) Encode(b interface{}) ([]byte, error) {
	if b.(bool) {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

func (_ *_Bool) Decode(b []byte) (interface{}, error) {
	if len(b) <= 0 {
		return false, nil
	}
	if b[0] != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (_ *_Bool) Wrap(b interface{}) bool {
	return b.(bool)
}

