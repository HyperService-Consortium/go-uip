package storage

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm"
)

type MappingUint8ToBytes struct {
	g *GeneralMap
}

func (v *VM) NewMappingUint8ToBytes(arrName string) *MappingUint8ToBytes {
	p := &MappingUint8ToBytes{&GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    gvm.Uint8, v: gvm.Bytes,
	}}
	return p
}

func (m *MappingUint8ToBytes) Set(k uint8, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint8ToBytes) Get(k uint8) []byte {
	return gvm.Bytes.Wrap(m.g.Get(k))

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
		k:    gvm.Uint16, v: gvm.Bytes,
	}}
	return p
}

func (m *MappingUint16ToBytes) Set(k uint16, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint16ToBytes) Get(k uint16) []byte {
	return gvm.Bytes.Wrap(m.g.Get(k))

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
		k:    gvm.Uint32, v: gvm.Bytes,
	}}
	return p
}

func (m *MappingUint32ToBytes) Set(k uint32, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint32ToBytes) Get(k uint32) []byte {
	return gvm.Bytes.Wrap(m.g.Get(k))

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
		k:    gvm.Uint64, v: gvm.Bytes,
	}}
	return p
}

func (m *MappingUint64ToBytes) Set(k uint64, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUint64ToBytes) Get(k uint64) []byte {
	return gvm.Bytes.Wrap(m.g.Get(k))

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
		k:    gvm.Uint, v: gvm.Bytes,
	}}
	return p
}

func (m *MappingUintToBytes) Set(k uint, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingUintToBytes) Get(k uint) []byte {
	return gvm.Bytes.Wrap(m.g.Get(k))

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
		k:    gvm.Bytes, v: gvm.Uint64,
	}}
	return p
}

func (m *MappingBytesToUint64) Set(k []byte, v uint64) {
	m.g.Set(k, v)
}

func (m *MappingBytesToUint64) Get(k []byte) uint64 {
	return gvm.Uint64.Wrap(m.g.Get(k))

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
		k:    gvm.Bytes, v: gvm.Uint32,
	}}
	return p
}

func (m *MappingBytesToUint32) Set(k []byte, v uint32) {
	m.g.Set(k, v)
}

func (m *MappingBytesToUint32) Get(k []byte) uint32 {
	return gvm.Uint32.Wrap(m.g.Get(k))

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
		k:    gvm.Bytes, v: gvm.Bytes,
	}}
	return p
}

func (m *MappingBytesToBytes) Set(k []byte, v []byte) {
	m.g.Set(k, v)
}

func (m *MappingBytesToBytes) Get(k []byte) []byte {
	return gvm.Bytes.Wrap(m.g.Get(k))

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
		k:    gvm.Bytes, v: gvm.Bool,
	}}
	return p
}

func (m *MappingBytesToBool) Set(k []byte, v bool) {
	m.g.Set(k, v)
}

func (m *MappingBytesToBool) Get(k []byte) bool {
	return gvm.Bool.Wrap(m.g.Get(k))

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
		k:    gvm.Uint64, v: gvm.Uint64,
	}}
	return p
}

func (m *MappingUint64ToUint64) Set(k uint64, v uint64) {
	m.g.Set(k, v)
}

func (m *MappingUint64ToUint64) Get(k uint64) uint64 {
	return gvm.Uint64.Wrap(m.g.Get(k))
}

func (m *MappingUint64ToUint64) Delete(k uint64) {
	m.g.Delete(k)
}
