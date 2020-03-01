package storage

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"math/big"
)

func Encode(variable token.ConstantI) ([]byte, error) {
	switch variable.GetParamType() {
	case value_type.Bool:
		return Bool.Encode(variable.GetConstant())
	case value_type.Bytes:
		return Bytes.Encode(variable.GetConstant())
	case value_type.String:
		return String.Encode(variable.GetConstant())
	case value_type.Int8:
		return Int8.Encode(variable.GetConstant())
	case value_type.Int16:
		return Int16.Encode(variable.GetConstant())
	case value_type.Int32:
		return Int32.Encode(variable.GetConstant())
	case value_type.Int64:
		return Int64.Encode(variable.GetConstant())
	case value_type.Int128:
		return Int128.Encode(variable.GetConstant())
	case value_type.Int256:
		return Int256.Encode(variable.GetConstant())
	case value_type.Uint8:
		return Uint8.Encode(variable.GetConstant())
	case value_type.Uint16:
		return Uint16.Encode(variable.GetConstant())
	case value_type.Uint32:
		return Uint32.Encode(variable.GetConstant())
	case value_type.Uint64:
		return Uint64.Encode(variable.GetConstant())
	case value_type.Uint128:
		return Uint128.Encode(variable.GetConstant())
	case value_type.Uint256:
		return Uint256.Encode(variable.GetConstant())
		//	SliceUint8
		//	SliceUint16
		//	SliceUint32
		//	SliceUint64
		//	SliceUint128 // 19
		//	SliceUint256
		//	SliceInt8
		//	SliceInt16
		//	SliceInt32
		//	SliceInt64 // 24
		//	SliceInt128
		//	SliceInt256
	default:
		return nil, errors.New("bad value_type")
	}
}

func Decode(x []byte, t value_type.Type) (token.ConstantI, error) {
	switch t {
	case value_type.Bytes:
		b, err := Bytes.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.([]byte),
		}, nil
	case value_type.String:
		b, err := String.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(string),
		}, nil
	case value_type.Uint8:
		b, err := Uint8.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(uint8),
		}, nil
	case value_type.Uint16:
		b, err := Uint16.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(uint16),
		}, nil
	case value_type.Uint32:
		b, err := Uint32.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(uint32),
		}, nil
	case value_type.Uint64:
		b, err := Uint64.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(uint64),
		}, nil
	case value_type.Uint128:
		b, err := Uint128.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(*big.Int),
		}, nil
	case value_type.Uint256:
		b, err := Uint256.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(*big.Int),
		}, nil
	case value_type.Int8:
		b, err := Int8.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(int8),
		}, nil
	case value_type.Int16:
		b, err := Int16.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(int16),
		}, nil
	case value_type.Int32:
		b, err := Int32.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(int32),
		}, nil
	case value_type.Int64:
		b, err := Int64.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(int64),
		}, nil
	case value_type.Int128:
		b, err := Int128.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(*big.Int),
		}, nil
	case value_type.Int256:
		b, err := Int256.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(*big.Int),
		}, nil
	case value_type.Bool:
		b, err := Bool.Decode(x)
		if err != nil {
			return nil, err
		}
		return lexer.ConstantVariable{
			Type:  t,
			Const: b.(bool),
		}, nil
	//	SliceUint8
	//	SliceUint16
	//	SliceUint32
	//	SliceUint64
	//	SliceUint128 // 19
	//	SliceUint256
	//	SliceInt8
	//	SliceInt16
	//	SliceInt32
	//	SliceInt64 // 24
	//	SliceInt128
	//	SliceInt256
	default:
		return nil, errors.New("bad value_type")
	}
}
