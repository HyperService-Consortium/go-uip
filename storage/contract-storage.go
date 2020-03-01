package storage



type VM struct {
	l          LocalStorage
	varStorage MerkMap

	callbacks []Callback
}

type Callback interface {
	AfterDone(v *VM) error
}

func NewVM(l LocalStorage) *VM {
	return &VM{l: l, varStorage: l.ArrangeSlot("")}
}

func (v *VM) ArrangeSlot(newSlot string) MerkMap {
	if len(newSlot) == 0 {
		return nil
	}
	return v.l.ArrangeSlot(newSlot)
}

func (v *VM) Callback() (err error) {
	for _, c := range v.callbacks {
		if err = c.AfterDone(v); err != nil {
			return
		}
	}
	return nil
}

// Bytes
// String
// Bool
// Map

// Uint8
// Uint16
// Uint32
// Uint64
// Uint128
// Uint256

// Int8
// Int16
// Int32
// Int64
// Int128
// Int256
