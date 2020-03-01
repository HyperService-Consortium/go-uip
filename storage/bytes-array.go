package storage

import (
	"github.com/HyperService-Consortium/go-uip/internal/util"
)

type BytesArray struct {
	name   string
	merk   MerkMap
	length uint64
}

func (v *VM) NewBytesArray(arrName string) *BytesArray {
	barr := &BytesArray{
		name:   arrName,
		merk:   v.ArrangeSlot(arrName),
		length: v.GetUint64(arrName),
	}
	v.callbacks = append(v.callbacks, barr)
	return barr
}

func (barr *BytesArray) AfterDone(v *VM) error {
	v.SetUint64(barr.name, barr.length)
	return nil
}

func (barr *BytesArray) Length() uint64 {
	return barr.length
}

func (barr *BytesArray) Set(o uint64, value []byte) (err error) {
	err = barr.merk.Update(util.Uint64ToBytes(o), value)
	if err != nil {
		return err
	}
	if barr.length < o {
		barr.length = o + 1
	}
	return
}

func (barr *BytesArray) Get(o uint64) []byte {
	k, err := barr.merk.Get(util.Uint64ToBytes(o))
	if err != nil {
		panic(wrapGet(err))
	}
	return k
}

func (barr *BytesArray) Delete(o uint64) {
	err := barr.merk.Delete(util.Uint64ToBytes(o))
	if err != nil {
		panic(wrapDelete(err))
	}
}

func (barr *BytesArray) Append(value []byte) {
	err := barr.merk.Update(util.Uint64ToBytes(barr.length), value)
	if err != nil {
		panic(err)
	}
	barr.length++
}
