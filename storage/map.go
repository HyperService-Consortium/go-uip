package storage

import (
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"math/big"
)

type _BigInt struct{}

func (_ _BigInt) Encode(v interface{}) ([]byte, error) {
	return v.(*big.Int).Bytes(), nil
}

func (_ _BigInt) Decode(b []byte) (interface{}, error) {
	return big.NewInt(0).SetBytes(b), nil
}

func (_ _BigInt) Wrap(b interface{}) *big.Int {
	return b.(*big.Int)
}

type _Uint8 struct{}

func (_ _Uint8) Encode(v interface{}) ([]byte, error) {
	return []byte{v.(uint8)}, nil
}

func (_ _Uint8) Decode(b []byte) (interface{}, error) {
	return b[0], nil
}

func (_ _Uint8) Wrap(b interface{}) uint8 {
	return b.(uint8)
}

type _Uint16 struct{}

func (_ _Uint16) Encode(v interface{}) ([]byte, error) {
	return util.Uint16ToBytes(v.(uint16)), nil
}

func (_ _Uint16) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint16(b), nil
}

func (_ _Uint16) Wrap(b interface{}) uint16 {
	return b.(uint16)
}

type _Uint32 struct{}

func (_ _Uint32) Encode(v interface{}) ([]byte, error) {
	return util.Uint32ToBytes(v.(uint32)), nil
}

func (_ _Uint32) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint32(b), nil
}

func (_ _Uint32) Wrap(b interface{}) uint32 {
	return b.(uint32)
}

type _Uint64 struct{}

func (_ _Uint64) Encode(v interface{}) ([]byte, error) {
	return util.Uint64ToBytes(v.(uint64)), nil
}

func (_ _Uint64) Decode(b []byte) (interface{}, error) {
	return util.BytesToUint64(b), nil
}

func (_ _Uint64) Wrap(b interface{}) uint64 {
	return b.(uint64)
}

type _Uint struct{}

func (_ _Uint) Encode(v interface{}) ([]byte, error) {
	return util.Uint64ToBytes(uint64(v.(uint))), nil
}

func (_ _Uint) Decode(b []byte) (interface{}, error) {
	return uint(util.BytesToUint64(b)), nil
}

func (_ _Uint) Wrap(b interface{}) uint {
	return b.(uint)
}

type _Int8 struct{}

func (_ _Int8) Encode(v interface{}) ([]byte, error) {
	return []byte{uint8(v.(int8))}, nil
}

func (_ _Int8) Decode(b []byte) (interface{}, error) {
	return b[0], nil
}

func (_ _Int8) Wrap(b interface{}) int8 {
	return b.(int8)
}

type _Int16 struct{}

func (_ _Int16) Encode(v interface{}) ([]byte, error) {
	return util.Int16ToBytes(v.(int16)), nil
}

func (_ _Int16) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt16(b), nil
}

func (_ _Int16) Wrap(b interface{}) int16 {
	return b.(int16)
}

type _Int32 struct{}

func (_ _Int32) Encode(v interface{}) ([]byte, error) {
	return util.Int32ToBytes(v.(int32)), nil
}

func (_ _Int32) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt32(b), nil
}

func (_ _Int32) Wrap(b interface{}) int32 {
	return b.(int32)
}

type _Int64 struct{}

func (_ _Int64) Encode(v interface{}) ([]byte, error) {
	return util.Int64ToBytes(v.(int64)), nil
}

func (_ _Int64) Decode(b []byte) (interface{}, error) {
	return util.BytesToInt64(b), nil
}

func (_ _Int64) Wrap(b interface{}) int64 {
	return b.(int64)
}

type _Int struct{}

func (_ _Int) Encode(v interface{}) ([]byte, error) {
	return util.Int64ToBytes(int64(v.(int))), nil
}

func (_ _Int) Decode(b []byte) (interface{}, error) {
	return uint(util.BytesToInt64(b)), nil
}

func (_ _Int) Wrap(b interface{}) int {
	return b.(int)
}

type _Bytes struct{}

func (_ _Bytes) Encode(b interface{}) ([]byte, error) {
	return b.([]byte), nil
}

func (_ _Bytes) Decode(b []byte) (interface{}, error) {
	return b, nil
}

func (_ _Bytes) Wrap(b interface{}) []byte {
	return b.([]byte)
}

type _String struct{}

func (_ _String) Encode(b interface{}) ([]byte, error) {
	return []byte(b.(string)), nil
}

func (_ _String) Decode(b []byte) (interface{}, error) {
	return string(b), nil
}

func (_ _String) Wrap(b interface{}) string {
	return b.(string)
}

type _Bool struct{}

func (_ _Bool) Encode(b interface{}) ([]byte, error) {
	if b.(bool) {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

func (_ _Bool) Decode(b []byte) (interface{}, error) {
	if len(b) <= 0 {
		return false, nil
	}
	if b[0] != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (_ _Bool) Wrap(b interface{}) bool {
	return b.(bool)
}

var (
	Uint8  = _Uint8{}
	Uint16 = _Uint16{}
	Uint32 = _Uint32{}
	Uint64 = _Uint64{}


	Int8  = _Int8{}
	Int16 = _Int16{}
	Int32 = _Int32{}
	Int64 = _Int64{}

	Uint128 = _BigInt{}
	Uint256 = _BigInt{}
	Int128 = _BigInt{}
	Int256 = _BigInt{}

	Uint   = _Uint{}
	Bytes  = _Bytes{}
	String  = _String{}
	Bool   = _Bool{}
)

type MappingUint8ToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingUint8ToBytes(arrName string) *MappingUint8ToBytes {
	p := &MappingUint8ToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Uint8, v: Bytes,
	}}
	return p
}

func (m *MappingUint8ToBytes) Set(k uint8, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint8ToBytes) Get(k uint8) []byte {
	return Bytes.Wrap(m.g.Get(k))

}

func (m *MappingUint8ToBytes) Delete(k uint8) {
	m.g.Delete(k)
}

type MappingUint16ToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingUint16ToBytes(arrName string) *MappingUint16ToBytes {
	p := &MappingUint16ToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Uint16, v: Bytes,
	}}
	return p
}

func (m *MappingUint16ToBytes) Set(k uint16, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint16ToBytes) Get(k uint16) []byte {
	return Bytes.Wrap(m.g.Get(k))

}

func (m *MappingUint16ToBytes) Delete(k uint16) {
	m.g.Delete(k)
}

type MappingUint32ToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingUint32ToBytes(arrName string) *MappingUint32ToBytes {
	p := &MappingUint32ToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Uint32, v: Bytes,
	}}
	return p
}

func (m *MappingUint32ToBytes) Set(k uint32, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint32ToBytes) Get(k uint32) []byte {
	return Bytes.Wrap(m.g.Get(k))

}

func (m *MappingUint32ToBytes) Delete(k uint32) {
	m.g.Delete(k)
}

type MappingUint64ToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingUint64ToBytes(arrName string) *MappingUint64ToBytes {
	p := &MappingUint64ToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Uint64, v: Bytes,
	}}
	return p
}

func (m *MappingUint64ToBytes) Set(k uint64, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint64ToBytes) Get(k uint64) []byte {
	return Bytes.Wrap(m.g.Get(k))

}

func (m *MappingUint64ToBytes) Delete(k uint64) {
	m.g.Delete(k)
}

type MappingUintToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingUintToBytes(arrName string) *MappingUintToBytes {
	p := &MappingUintToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Uint, v: Bytes,
	}}
	return p
}

func (m *MappingUintToBytes) Set(k uint, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUintToBytes) Get(k uint) []byte {
	return Bytes.Wrap(m.g.Get(k))

}

func (m *MappingUintToBytes) Delete(k uint) {
	m.g.Delete(k)
}

type MappingBytesToUint64 struct {
	g *GeneralMap
}

func (v *VM) NewMappingBytesToUint64(arrName string) *MappingBytesToUint64 {
	p := &MappingBytesToUint64{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Bytes, v: Uint64,
	}}
	return p
}

func (m *MappingBytesToUint64) Set(k []byte, v uint64) {
	m.g.Set(k, v)
}

func (m *MappingBytesToUint64) Get(k []byte) uint64 {
	return Uint64.Wrap(m.g.Get(k))

}

func (m *MappingBytesToUint64) Delete(k []byte) {
	m.g.Delete(k)
}

type MappingBytesToUint32 struct {
	g *GeneralMap
}

func (v *VM) NewMappingBytesToUint32(arrName string) *MappingBytesToUint32 {
	p := &MappingBytesToUint32{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Bytes, v: Uint32,
	}}
	return p
}

func (m *MappingBytesToUint32) Set(k []byte, v uint32) {
	m.g.Set(k, v)
}

func (m *MappingBytesToUint32) Get(k []byte) uint32 {
	return Uint32.Wrap(m.g.Get(k))

}

func (m *MappingBytesToUint32) Delete(k []byte) {
	m.g.Delete(k)
}

type MappingBytesToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingBytesToBytes(arrName string) *MappingBytesToBytes {
	p := &MappingBytesToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Bytes, v: Bytes,
	}}
	return p
}

func (m *MappingBytesToBytes) Set(k []byte, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingBytesToBytes) Get(k []byte) []byte {
	return Bytes.Wrap(m.g.Get(k))

}

func (m *MappingBytesToBytes) Delete(k []byte) {
	m.g.Delete(k)
}

type MappingBytesToBool struct {
	g *GeneralMap
}

func (v *VM) NewMappingBytesToBool(arrName string) *MappingBytesToBool {
	p := &MappingBytesToBool{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Bytes, v: Bool,
	}}
	return p
}

func (m *MappingBytesToBool) Set(k []byte, v bool) {
	m.g.Set(k, v)
}

func (m *MappingBytesToBool) Get(k []byte) bool {
	return Bool.Wrap(m.g.Get(k))

}

func (m *MappingBytesToBool) Delete(k []byte) {
	m.g.Delete(k)
}

type MappingUint64ToUint64 struct {
	g *GeneralMap
}

func (v *VM) NewMappingUint64ToUint64(arrName string) *MappingUint64ToUint64 {
	p := &MappingUint64ToUint64{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    Uint64, v: Uint64,
	}}
	return p
}

func (m *MappingUint64ToUint64) Set(k uint64, v uint64) {
	m.g.Set(k, v)
}

func (m *MappingUint64ToUint64) Get(k uint64) uint64 {
	return Uint64.Wrap(m.g.Get(k))
}

func (m *MappingUint64ToUint64) Delete(k uint64) {
	m.g.Delete(k)
}
