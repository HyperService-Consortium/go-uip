package gvm

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
)

func convErr(l, r Ref) (err error) {
	return fmt.Errorf("cant conv param type %v and %v", l.GetGVMType(), r.GetGVMType())
}

func convArithmetic(l Ref, r Ref) (_ Ref, _ Ref, err error) {

	li, _, _ := getBasicTraits(l)

	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, nil, fmt.Errorf("operate on non-integral type: %v and %v", l.GetGVMType(), r.GetGVMType())
	}
	if l.GetGVMType() != r.GetGVMType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, nil, err
		}
	}
	return l, r, nil
}

func getBasicTraits(l Ref) (integral bool, unsigned bool, k int) {
	integral, unsigned, k = true, false, 0
	switch l.GetGVMType() {
	case RefBool, RefString, RefBytes:
		integral = false
	case RefUint8, RefUint16, RefUint32, RefUint64, RefUint128, RefUint256:
		unsigned = false
		switch l.GetGVMType() {
		case RefUint8:
			k = 8
		case RefUint16:
			k = 16
		case RefUint32:
			k = 32
		case RefUint64:
			k = 64
		case RefUint128:
			k = 128
		case RefUint256:
			k = 256
		}
	case RefInt8, RefInt16, RefInt32, RefInt64, RefInt128, RefInt256:
		unsigned = true
		switch l.GetGVMType() {
		case RefInt8:
			k = 8
		case RefInt16:
			k = 16
		case RefInt32:
			k = 32
		case RefInt64:
			k = 64
		case RefInt128:
			k = 128
		case RefInt256:
			k = 256
		}
	default:
		panic(errors.New("value type not found"))
	}
	return
}

func extendType(l Ref) interface{} {
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
	case RefBool, RefString, RefBytes:
		return l
	}

	panic(fmt.Errorf("error in extendType: type %v, true-type %v", l.GetGVMType(), reflect.TypeOf(l.Unwrap())))
}

func toBool(l Ref) (bool, error) {
	switch l.GetGVMType() {
	case RefBool:
		return l.Unwrap().(bool), nil
	case RefBytes:
		return len(l.Unwrap().([]byte)) != 0, nil
	case RefString:
		return len(l.Unwrap().(string)) != 0, nil
	case RefUint8:
		return (l.Unwrap().(uint8)) != 0, nil
	case RefUint16:
		return (l.Unwrap().(uint16)) != 0, nil
	case RefUint32:
		return (l.Unwrap().(uint32)) != 0, nil
	case RefUint64:
		return l.Unwrap().(uint64) != 0, nil
	case RefInt8:
		return (l.Unwrap().(int8)) != 0, nil
	case RefInt16:
		return (l.Unwrap().(int16)) != 0, nil
	case RefInt32:
		return (l.Unwrap().(int32)) != 0, nil
	case RefInt64:
		return l.Unwrap().(int64) != 0, nil
	case RefUint128, RefUint256, RefInt128, RefInt256:
		return l.Unwrap().(*big.Int).Sign() == 0, nil
	}
	return false, fmt.Errorf("error in toBool: type %v, true-type %v", l.GetGVMType(), reflect.TypeOf(l.Unwrap()))
}

func unsafeShrinkType(value interface{}, paramType RefType) interface{} {
	switch paramType {
	case RefUint8:
		return uint8(value.(uint64))
	case RefUint16:
		return uint16(value.(uint64))
	case RefUint32:
		return uint32(value.(uint64))
	case RefUint64:
		return value.(uint64)
	case RefInt8:
		return int8(value.(int64))
	case RefInt16:
		return int16(value.(int64))
	case RefInt32:
		return int32(value.(int64))
	case RefInt64:
		return value.(int64)
	case RefUint128, RefUint256, RefInt128, RefInt256:
		return value.(*big.Int)
	}

	panic(fmt.Errorf("error in unsafeShrinkType: type %v, true-type %v", paramType, reflect.TypeOf(value)))
}

func tryConv(l Ref, r Ref) (_ Ref, _ Ref, err error) {
	if l.GetGVMType() == r.GetGVMType() {
		return l, r, nil
	}
	li, lu, lk := getBasicTraits(l)
	ri, ru, rk := getBasicTraits(r)
	if li && ri {
		if lu && !ru {
			if r, err = tryConvUnsigned(r); err != nil {
				return nil, nil, err
			}
			ru = true
		}
		if ru && !lu {
			if l, err = tryConvUnsigned(l); err != nil {
				return nil, nil, err
			}
			lu = true
		}
		if lk != rk {
			if lk < rk {
				return extentBit(l, rk), r, nil
			} else {
				return l, extentBit(r, lk), nil
			}
		}

	} else {
		return nil, nil, convErr(l, r)
	}

	return l, r, nil
}

func extentBit(l Ref, rk int) Ref {
	switch l.GetGVMType() {
	case RefUint8, RefUint16, RefUint32, RefUint64:
		k := reflect.ValueOf(l.Unwrap()).Uint()
		switch rk {
		case 128:
			return ref{
				t:  RefUint128,
				v: big.NewInt(0).SetUint64(k),
			}
		case 256:
			return ref{
				t:  RefUint256,
				v: big.NewInt(0).SetUint64(k),
			}
		default:
			switch rk {
			case 8:
				return ref{
					t:  RefUint8,
					v: uint8(k),
				}
			case 16:
				return ref{
					t:  RefUint16,
					v: uint16(k),
				}
			case 32:
				return ref{
					t:  RefUint32,
					v: uint32(k),
				}
			case 64:
				return ref{
					t:  RefUint64,
					v: uint64(k),
				}
			default:
				panic(fmt.Errorf("bit not found: %v", rk))
			}
		}
	case RefInt8, RefInt16, RefInt32, RefInt64:
		k := reflect.ValueOf(l.Unwrap()).Int()
		switch rk {
		case 128:
			return ref{
				t:  RefInt128,
				v: big.NewInt(k),
			}
		case 256:
			return ref{
				t:  RefInt256,
				v: big.NewInt(k),
			}
		default:
			switch rk {
			case 8:
				return ref{
					t:  RefInt8,
					v: int8(k),
				}
			case 16:
				return ref{
					t:  RefInt16,
					v: int16(k),
				}
			case 32:
				return ref{
					t:  RefInt32,
					v: int32(k),
				}
			case 64:
				return ref{
					t:  RefInt64,
					v: int64(k),
				}
			default:
				panic(fmt.Errorf("bit not found: %v", rk))
			}
		}
	case RefBytes, RefBool, RefString:
		panic(fmt.Errorf("wrong extendBit on non integral type"))
	case RefUint128, RefUint256:
		return ref{
			t:  RefUint256,
			v: l.Unwrap(),
		}
	case RefInt256, RefInt128:
		return ref{
			t:  RefInt256,
			v: l.Unwrap(),
		}
	default:
		panic(errors.New("value type not found"))
	}
}

func tryConvUnsigned(k Ref) (Ref, error) {
	switch k.GetGVMType() {
	case RefInt8:
		if k.Unwrap().(int8) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return ref{
			t:  RefUint8,
			v: uint8(k.Unwrap().(int8)),
		}, nil
	case RefInt16:
		if k.Unwrap().(int16) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return ref{
			t:  RefUint16,
			v: uint16(k.Unwrap().(int16)),
		}, nil
	case RefInt32:
		if k.Unwrap().(int32) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return ref{
			t:  RefUint32,
			v: uint32(k.Unwrap().(int32)),
		}, nil
	case RefInt64:
		if k.Unwrap().(int64) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return ref{
			t: RefUint64,
			v: uint64(k.Unwrap().(int64)),
		}, nil
	case RefInt128, RefInt256:
		if k.Unwrap().(*big.Int).Sign() < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}

	}
	return nil, errors.New("value type invalid")
}
