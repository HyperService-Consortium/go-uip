package isc

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"math/big"
	"reflect"
)


func convErr(l, r token.ConstantI) (err error) {
	return fmt.Errorf("cant conv param type %v and %v", l.GetParamType(), r.GetParamType())
}

func getBasicTraits(l token.ConstantI) (integral bool, unsigned bool, k int) {
	integral, unsigned, k = true, false, 0
	switch l.GetParamType() {
	case value_type.Bool, value_type.String, value_type.Bytes:
		integral = false
	case value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64, value_type.Uint128, value_type.Uint256:
		unsigned = false
		switch l.GetParamType() {
		case value_type.Uint8:
			k = 8
		case value_type.Uint16:
			k = 16
		case value_type.Uint32:
			k = 32
		case value_type.Uint64:
			k = 64
		case value_type.Uint128:
			k = 128
		case value_type.Uint256:
			k = 256
		}
	case value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64, value_type.Int128, value_type.Int256:
		unsigned = true
		switch l.GetParamType() {
		case value_type.Int8:
			k = 8
		case value_type.Int16:
			k = 16
		case value_type.Int32:
			k = 32
		case value_type.Int64:
			k = 64
		case value_type.Int128:
			k = 128
		case value_type.Int256:
			k = 256
		}
	default:
		panic(errors.New("value type not found"))
	}
	return
}

func tryConv(l token.ConstantI, r token.ConstantI) (_ token.ConstantI, _ token.ConstantI, err error) {
	if l.GetParamType() == r.GetParamType() {
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

func extentBit(l token.ConstantI, rk int) token.ConstantI {
	switch l.GetParamType() {
	case value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64:
		k := reflect.ValueOf(l.GetConstant()).Uint()
		switch rk {
		case 128:
			return &lexer.ConstantVariable{
				Type:  value_type.Uint128,
				Const: big.NewInt(0).SetUint64(k),
			}
		case 256:
			return &lexer.ConstantVariable{
				Type:  value_type.Uint256,
				Const: big.NewInt(0).SetUint64(k),
			}
		default:
			switch rk {
			case 8:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint8,
					Const: uint8(k),
				}
			case 16:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint16,
					Const: uint16(k),
				}
			case 32:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint32,
					Const: uint32(k),
				}
			case 64:
				return &lexer.ConstantVariable{
					Type:  value_type.Uint64,
					Const: uint64(k),
				}
			default:
				panic(fmt.Errorf("bit not found: %v", rk))
			}
		}
	case value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64:
		k := reflect.ValueOf(l.GetConstant()).Int()
		switch rk {
		case 128:
			return &lexer.ConstantVariable{
				Type:  value_type.Int128,
				Const: big.NewInt(k),
			}
		case 256:
			return &lexer.ConstantVariable{
				Type:  value_type.Int256,
				Const: big.NewInt(k),
			}
		default:
			switch rk {
			case 8:
				return &lexer.ConstantVariable{
					Type:  value_type.Int8,
					Const: int8(k),
				}
			case 16:
				return &lexer.ConstantVariable{
					Type:  value_type.Int16,
					Const: int16(k),
				}
			case 32:
				return &lexer.ConstantVariable{
					Type:  value_type.Int32,
					Const: int32(k),
				}
			case 64:
				return &lexer.ConstantVariable{
					Type:  value_type.Int64,
					Const: int64(k),
				}
			default:
				panic(fmt.Errorf("bit not found: %v", rk))
			}
		}
	case value_type.Bytes, value_type.Bool, value_type.String:
		panic(fmt.Errorf("wrong extendBit on non integral type"))
	case value_type.Uint128, value_type.Uint256:
		return &lexer.ConstantVariable{
			Type:  value_type.Uint256,
			Const: l.GetConstant(),
		}
	case value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  value_type.Int256,
			Const: l.GetConstant(),
		}
	default:
		panic(errors.New("value type not found"))
	}
}

func tryConvUnsigned(k token.ConstantI) (token.ConstantI, error) {
	switch k.GetParamType() {
	case value_type.Int8:
		if k.GetConstant().(int8) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint8,
			Const:uint8( k.GetConstant().(int8)),
		}, nil
	case value_type.Int16:
		if k.GetConstant().(int16) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint16,
			Const: uint16(k.GetConstant().(int16)),
		}, nil
	case value_type.Int32:
		if k.GetConstant().(int32) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint32,
			Const: uint32(k.GetConstant().(int32)),
		}, nil
	case value_type.Int64:
		if k.GetConstant().(int64) < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
		return &lexer.ConstantVariable{
			Type:  value_type.Uint64,
			Const: uint64(k.GetConstant().(int64)),
		}, nil
	case value_type.Int128, value_type.Int256:
		if k.GetConstant().(*big.Int).Sign() < 0 {
			return nil, errors.New("cant conv negative integer to unsigned")
		}
	}
	return nil, errors.New("value type invalid")
}
