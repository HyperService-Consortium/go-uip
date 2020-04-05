package storage

import (
	"github.com/pkg/errors"
)

type LFunc func(interface{}) ([]byte, error)
type RFunc func([]byte) (interface{}, error)

type Object interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

type GeneralMap struct {
	name string
	merk MerkMap
	k, v Object
}

func (v *VM) NewGeneralMap(arrName string, key, value Object) *GeneralMap {
	barr := &GeneralMap{
		name: arrName,
		merk: v.ArrangeSlot(arrName),
		k:    key, v: value,
	}
	return barr
}

func wrapErr(err error, msg string) error {
	return errors.WithStack(errors.Wrap(err, msg))
}

func wrapEncode(err error) error {
	return wrapErr(err, "encode error, report this to github repo")
}

func wrapDecode(err error) error {
	return wrapErr(err, "decode error, report this to github repo")
}

func wrapGet(err error) error {
	return wrapErr(err, "get error, report this to github repo")
}

func wrapDelete(err error) error {
	return errors.WithStack(errors.Wrap(err, "delete error, report this to github repo"))
}

func wrapUpdate(err error) error {
	return wrapErr(err, "update error, report this to github repo")
}

func (udMap *GeneralMap) Set(o interface{}, value interface{}) {
	if b, err := udMap.k.Encode(o); err != nil {
		panic(wrapEncode(err))
	} else if c, err := udMap.v.Encode(value); err != nil {
		panic(wrapDecode(err))
	} else if err = udMap.merk.Update(b, c); err != nil {
		panic(wrapUpdate(err))
	}
}

func (udMap *GeneralMap) Get(o interface{}) interface{} {
	if b, err := udMap.k.Encode(o); err != nil {
		panic(wrapEncode(err))
	} else if b, err = udMap.merk.Get(b); err != nil {
		panic(wrapGet(err))
	} else {
		i, err := udMap.v.Decode(b)
		if err != nil {
			panic(wrapDecode(err))
		}
		return i
	}
}

func (udMap *GeneralMap) Delete(o interface{}) {
	if b, err := udMap.k.Encode(o); err != nil {
		panic(wrapEncode(err))
	} else if err = udMap.merk.Delete(b); err != nil {
		panic(wrapDelete(err))
	}
}
