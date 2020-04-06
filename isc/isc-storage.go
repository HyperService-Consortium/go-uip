package isc

import (
	"errors"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"math/big"
)

func (isc *ISC) Load(field string, t gvm.RefType) (gvm.Ref, error) {
	switch t {
	case gvm_type.RefUint8:
		return isc.CreateRef(t, isc.Storage.storage.GetUint8(field)), nil
	case gvm_type.RefUint16:
		return isc.CreateRef(t, uint16(isc.Storage.storage.GetUint64(field))), nil
	case gvm_type.RefUint32:
		return isc.CreateRef(t, uint32(isc.Storage.storage.GetUint64(field))), nil
	case gvm_type.RefUint64:
		return isc.CreateRef(t, isc.Storage.storage.GetUint64(field)), nil
	case gvm_type.RefInt8:
		return isc.CreateRef(t, isc.Storage.storage.GetInt8(field)), nil
	case gvm_type.RefInt16:
		return isc.CreateRef(t, int16(isc.Storage.storage.GetInt64(field))), nil
	case gvm_type.RefInt32:
		return isc.CreateRef(t, int32(isc.Storage.storage.GetInt64(field))), nil
	case gvm_type.RefInt64:
		return isc.CreateRef(t, isc.Storage.storage.GetInt64(field)), nil
	case gvm_type.RefBool:
		return isc.CreateRef(t, isc.Storage.storage.GetBool(field)), nil
	case gvm_type.RefString:
		return isc.CreateRef(t, isc.Storage.storage.GetString(field)), nil
	case gvm_type.RefBytes, gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
		switch t {
		case gvm_type.RefBytes:
			return isc.CreateRef(t, isc.Storage.storage.GetBytes(field)), nil
		case gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
			return isc.CreateRef(t, big.NewInt(0).SetBytes(isc.Storage.storage.GetBytes(field))), nil
		default:
			panic("runtime case error")
		}
	default:
		return nil, errors.New("type error")
	}
}

func (isc *ISC) Save(field string, k gvm.Ref) error {
	switch k.GetGVMType() {
	case gvm_type.RefUint8:
		isc.Storage.storage.SetUint8(field, k.Unwrap().(uint8))
		return nil
	case gvm_type.RefUint16:
		isc.Storage.storage.SetUint64(field, uint64(k.Unwrap().(uint16)))
		return nil
	case gvm_type.RefUint32:
		isc.Storage.storage.SetUint64(field, uint64(k.Unwrap().(uint32)))
		return nil
	case gvm_type.RefUint64:
		isc.Storage.storage.SetUint64(field, k.Unwrap().(uint64))
		return nil
	case gvm_type.RefInt8:
		isc.Storage.storage.SetInt8(field, k.Unwrap().(int8))
		return nil
	case gvm_type.RefInt16:
		isc.Storage.storage.SetInt64(field, int64(k.Unwrap().(int16)))
		return nil
	case gvm_type.RefInt32:
		isc.Storage.storage.SetInt64(field, int64(k.Unwrap().(int32)))
		return nil
	case gvm_type.RefInt64:
		isc.Storage.storage.SetInt64(field, k.Unwrap().(int64))
		return nil
	case gvm_type.RefBool:
		isc.Storage.storage.SetBool(field, k.Unwrap().(bool))
		return nil
	case gvm_type.RefString:
		isc.Storage.storage.SetString(field, k.Unwrap().(string))
		return nil
	case gvm_type.RefBytes, gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
		var v = k.Unwrap()
		if v == nil && k.GetGVMType() == gvm_type.RefBytes {
			isc.Storage.storage.SetBytes(field, []byte{})
			return nil
		}
		switch k.GetGVMType() {
		case gvm_type.RefBytes:
			isc.Storage.storage.SetBytes(field, k.Unwrap().([]byte))
			return nil
		case gvm_type.RefUint128, gvm_type.RefUint256, gvm_type.RefInt128, gvm_type.RefInt256:
			b, err := k.Encode()
			if err != nil {
				return err
			}
			isc.Storage.storage.SetBytes(field, b)
			return nil
		default:
			panic("runtime case error")
		}
	default:
		return errors.New("type error")
	}
}

func (isc *ISC) Delete(_ string) error {
	return nil
}
