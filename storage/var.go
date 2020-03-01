package storage

import (
	"github.com/HyperService-Consortium/go-uip/internal/util"
	"github.com/HyperService-Consortium/go-uip/math"
)

func (v *VM) SetBytes(variName string, value []byte) {
	err := v.varStorage.Update([]byte(variName), value)
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetString(variName string, value string) {
	err := v.varStorage.Update([]byte(variName), []byte(value))
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetUint64(variName string, value uint64) {
	err := v.varStorage.Update([]byte(variName), util.Uint64ToBytes(value))
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetUint256(variName string, value *math.Uint256) {
	err := v.varStorage.Update([]byte(variName), value.Bytes())
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetInt64(variName string, value int64) {
	err := v.varStorage.Update([]byte(variName), util.Int64ToBytes(value))
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetInt8(variName string, value int8) {
	err := v.varStorage.Update([]byte(variName), []byte{uint8(value)})
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetUint8(variName string, value uint8) {
	err := v.varStorage.Update([]byte(variName), []byte{value})
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) SetBool(variName string, value bool) {
	var err error
	if value {
		err = v.varStorage.Update([]byte(variName), []byte{1})
	} else {
		err = v.varStorage.Update([]byte(variName), []byte{0})
	}
	if err != nil {
		panic(err)
	}
	return
}

type Bytable interface {
	Bytes() []byte
}

func (v *VM) SetAny(variName string, value Bytable) {
	err := v.varStorage.Update([]byte(variName), value.Bytes())
	if err != nil {
		panic(err)
	}
	return
}

func (v *VM) GetBytes(variName string) []byte {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	return bt
}

func (v *VM) GetString(variName string) string {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	return string(bt)
}

func (v *VM) GetUint64(variName string) uint64 {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 8 {
		panic("Decode Error: the length of getting value is more than 8")
	}
	return util.BytesToUint64(append(make([]byte, 8 - len(bt), 8), bt...))
}

func (v *VM) GetUint256(variName string) *math.Uint256 {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 32 {
		panic("Decode Error: the length of getting value is more than 32")
	}
	return math.NewUint256FromBytes(append(make([]byte, 32 - len(bt), 32), bt...))
}

func (v *VM) GetInt64(variName string) int64 {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 8 {
		panic("Decode Error: the length of getting value is more than 8")
	}
	return util.BytesToInt64(append(make([]byte, 8 - len(bt), 8), bt...))
}

func (v *VM) GetInt8(variName string) int8 {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 1 {
		panic("Decode Error: the length of getting value is more than 1")
	}
	if bt == nil || len(bt) == 0 {
		return 0
	} else {
		return int8(bt[0])
	}
}

func (v *VM) GetUint8(variName string) uint8 {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) > 1 {
		panic("Decode Error: the length of getting value is more than 1")
	}
	if bt == nil || len(bt) == 0 {
		return 0
	} else {
		return uint8(bt[0])
	}
}

func (v *VM) GetBool(variName string) bool {
	bt, err := v.varStorage.Get([]byte(variName))
	if err != nil {
		panic(err)
	}
	if len(bt) != 1 {
		panic("Decode Error: the length of getting value is not 1")
	}
	return bt[0] != 0
}