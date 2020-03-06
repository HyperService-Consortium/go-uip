package libgvm

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"math/big"
	"reflect"
)

// GetArithmeticalCharacteristic gets the arithmetical characteristic from a runtime value
func GetArithmeticalCharacteristic(r abstraction.Ref) (integral bool, unsigned bool, bitCount int) {
	return GetArithmeticalCharacteristicT(r.GetGVMType())
}

// GetArithmeticalCharacteristic gets the arithmetical characteristic from a type index
// if t is not a arithmetical type, then integral == false, and the value of unsigned and bitCount are uncertain.
func GetArithmeticalCharacteristicT(t abstraction.RefType) (integral bool, unsigned bool, bitCount int) {
	switch t {
	case RefUint8, RefUint16, RefUint32, RefUint64, RefUint128, RefUint256:
		unsigned, integral = true, true
		switch t {
		case RefUint8:
			bitCount = 8
		case RefUint16:
			bitCount = 16
		case RefUint32:
			bitCount = 32
		case RefUint64:
			bitCount = 64
		case RefUint128:
			bitCount = 128
		case RefUint256:
			bitCount = 256
		}
	case RefInt8, RefInt16, RefInt32, RefInt64, RefInt128, RefInt256:
		unsigned, integral = false, true
		switch t {
		case RefInt8:
			bitCount = 8
		case RefInt16:
			bitCount = 16
		case RefInt32:
			bitCount = 32
		case RefInt64:
			bitCount = 64
		case RefInt128:
			bitCount = 128
		case RefInt256:
			bitCount = 256
		}
	default:
		integral = false
	}
	return
}

// TryAlignArithmetic aligns l and r to the minimal type that can contain both of them.
func TryAlignArithmetic(l abstraction.Ref, r abstraction.Ref) (_ abstraction.Ref, _ abstraction.Ref, err error) {
	if l.GetGVMType() == r.GetGVMType() {
		return l, r, nil
	}
	li, lu, lk := GetArithmeticalCharacteristic(l)
	ri, ru, rk := GetArithmeticalCharacteristic(r)
	if li && ri {
		if lk < rk {
			l, err = ExtendBit(l, rk)
			if lu != ru && err == nil {
				if ru {
					l, err = TryConvertUnsigned(l)
				} else {
					l, err = TryConvertSigned(l)
				}
			}
			return l, r, err
		} else if lk > rk {
			r, err = ExtendBit(r, lk)
			if lu != ru && err == nil {
				if lu {
					r, err = TryConvertUnsigned(r)
				} else {
					r, err = TryConvertSigned(r)
				}
			}
			return l, r, err
		}
		if lu && !ru {
			if r, err = TryConvertUnsigned(r); err != nil {
				return nil, nil, err
			}
		} else if ru && !lu {
			if l, err = TryConvertUnsigned(l); err != nil {
				return nil, nil, err
			}
		}

	} else {
		return nil, nil, ConvertError(l, r)
	}

	return l, r, nil
}

// ExtendType converts r.value to uint64/int64/uint256/int256.
func ExtendType(l abstraction.Ref) interface{} {
	switch l.GetGVMType() {
	case RefUint8:
		return uint64(l.Unwrap().(uint8))
	case RefUint16:
		return uint64(l.Unwrap().(uint16))
	case RefUint32:
		return uint64(l.Unwrap().(uint32))
	case RefUint64:
		return l.Unwrap().(uint64)
	case RefInt8:
		return int64(l.Unwrap().(int8))
	case RefInt16:
		return int64(l.Unwrap().(int16))
	case RefInt32:
		return int64(l.Unwrap().(int32))
	case RefInt64:
		return l.Unwrap().(int64)
	case RefUint128, RefUint256, RefInt128, RefInt256:
		return l.Unwrap().(*big.Int)
	default:
		return l
	}
}

// ToBool converts r.value to boolean type. if r.value is not boolean or arithmetical, then it returns an error
func ToBool(r abstraction.Ref) (bool, error) {
	switch r.GetGVMType() {
	case RefBool:
		return r.Unwrap().(bool), nil
	case RefUint8:
		return (r.Unwrap().(uint8)) != 0, nil
	case RefUint16:
		return (r.Unwrap().(uint16)) != 0, nil
	case RefUint32:
		return (r.Unwrap().(uint32)) != 0, nil
	case RefUint64:
		return r.Unwrap().(uint64) != 0, nil
	case RefInt8:
		return (r.Unwrap().(int8)) != 0, nil
	case RefInt16:
		return (r.Unwrap().(int16)) != 0, nil
	case RefInt32:
		return (r.Unwrap().(int32)) != 0, nil
	case RefInt64:
		return r.Unwrap().(int64) != 0, nil
	case RefUint128, RefUint256, RefInt128, RefInt256:
		return r.Unwrap().(*big.Int).Sign() != 0, nil
	}
	return false, fmt.Errorf("error in toBool: type %v, true-type %v", ExplainGVMType(r.GetGVMType()), reflect.TypeOf(r.Unwrap()))
}

// UnsafeShrinkType wraps value to the reference with ref-type paramType
func UnsafeShrinkType(value interface{}, paramType abstraction.RefType) (abstraction.Ref, error) {
	switch paramType {
	case RefUint8:
		return Uint8(value.(uint64)), nil
	case RefUint16:
		return Uint16(value.(uint64)), nil
	case RefUint32:
		return Uint32(value.(uint64)), nil
	case RefUint64:
		return Uint64(value.(uint64)), nil
	case RefInt8:
		return Int8(value.(int64)), nil
	case RefInt16:
		return Int16(value.(int64)), nil
	case RefInt32:
		return Int32(value.(int64)), nil
	case RefInt64:
		return Int64(value.(int64)), nil
	case RefUint128:
		return (*Uint128)(value.(*big.Int)), nil
	case RefInt128:
		return (*Int128)(value.(*big.Int)), nil
	case RefUint256:
		return (*Uint256)(value.(*big.Int)), nil
	case RefInt256:
		return (*Int256)(value.(*big.Int)), nil
	}

	return nil, fmt.Errorf("error in unsafeShrinkType: type %v, true-type %v", ExplainGVMType(paramType), reflect.TypeOf(value))
}

// ExtendBit extends a number by applying sign extension
// if the ref value is not arithmetical, it returns error
// we represent the type of ref by a-x. if a-x or a-bitCount is not a valid type, it returns an error.
// it does not check whether x > bitCount or not.
func ExtendBit(r abstraction.Ref, bitCount int) (abstraction.Ref, error) {
	switch r.GetGVMType() {
	case RefUint8, RefUint16, RefUint32, RefUint64:
		k := reflect.ValueOf(r.Unwrap()).Uint()
		switch bitCount {
		case 128:
			return (*Uint128)(big.NewInt(0).SetUint64(k)), nil
		case 256:
			return (*Uint256)(big.NewInt(0).SetUint64(k)), nil
		default:
			switch bitCount {
			case 8:
				return Uint8(k), nil
			case 16:
				return Uint16(k), nil
			case 32:
				return Uint32(k), nil
			case 64:
				return Uint64(k), nil
			default:
				return nil, fmt.Errorf("type Uint%v not found", bitCount)
			}
		}
	case RefInt8, RefInt16, RefInt32, RefInt64:
		k := reflect.ValueOf(r.Unwrap()).Int()
		switch bitCount {
		case 128:
			return (*Int128)(big.NewInt(k)), nil
		case 256:
			return (*Int256)(big.NewInt(k)), nil
		default:
			switch bitCount {
			case 8:
				return Int8(k), nil
			case 16:
				return Int16(k), nil
			case 32:
				return Int32(k), nil
			case 64:
				return Int64(k), nil
			default:
				return nil, fmt.Errorf("type Int%v not found", bitCount)
			}
		}
	case RefUint128, RefUint256:
		return (*Uint256)(r.Unwrap().(*big.Int)), nil
	case RefInt256, RefInt128:
		return (*Int256)(r.Unwrap().(*big.Int)), nil
	}
	return nil, InvalidTypeError(r)
}

// TryConvertUnsigned converts a number type to an unsigned number type.
// if the ref value is not arithmetical or negative, it returns an error
func TryConvertUnsigned(k abstraction.Ref) (abstraction.Ref, error) {
	switch k.GetGVMType() {
	case RefUint8, RefUint16, RefUint32, RefUint64, RefUint128, RefUint256:
		return k, nil
	case RefInt8:
		if k.Unwrap().(int8) < 0 {
			return nil, ConvertUnsignedError(k)
		}
		return Uint8(k.Unwrap().(int8)), nil
	case RefInt16:
		if k.Unwrap().(int16) < 0 {
			return nil, ConvertUnsignedError(k)
		}
		return Uint16(k.Unwrap().(int16)), nil
	case RefInt32:
		if k.Unwrap().(int32) < 0 {
			return nil, ConvertUnsignedError(k)
		}
		return Uint32(k.Unwrap().(int32)), nil
	case RefInt64:
		if k.Unwrap().(int64) < 0 {
			return nil, ConvertUnsignedError(k)
		}
		return Uint64(k.Unwrap().(int64)), nil
	case RefInt128, RefInt256:
		if k.Unwrap().(*big.Int).Sign() < 0 {
			return nil, ConvertUnsignedError(k)
		}
		if k.GetGVMType() == RefInt128 {
			return (*Uint128)(k.Unwrap().(*big.Int)), nil
		} else {
			return (*Uint256)(k.Unwrap().(*big.Int)), nil
		}
	}
	return nil, InvalidTypeError(k)
}

// TryConvertSigned converts a number type to an signed number type.
// if the ref value is not arithmetical, it returns an error
// we represent the type of ref by a-x. it maybe lose precision when value of ref v >= 2^{x/2}, but will always be right
// if it helps convert uint-x1 to int-x2 where x1 < x2
func TryConvertSigned(k abstraction.Ref) (abstraction.Ref, error) {
	switch k.GetGVMType() {
	case RefInt8, RefInt16, RefInt32, RefInt64, RefInt128, RefInt256:
		return k, nil
	case RefUint8:
		return Int8(k.Unwrap().(uint8)), nil
	case RefUint16:
		return Int16(k.Unwrap().(uint16)), nil
	case RefUint32:
		return Int32(k.Unwrap().(uint32)), nil
	case RefUint64:
		return Int64(k.Unwrap().(uint64)), nil
	case RefUint128, RefUint256:
		if k.GetGVMType() == RefUint128 {
			return (*Int128)(k.Unwrap().(*big.Int)), nil
		} else {
			return (*Int256)(k.Unwrap().(*big.Int)), nil
		}
	}
	return nil, InvalidTypeError(k)
}
