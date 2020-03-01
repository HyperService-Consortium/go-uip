package isc

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"math/big"
)

func add(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _add(l, r)
}

func _add(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) + r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) + r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) + r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) + r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) + r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) + r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) + r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) + r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Add(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _sub(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) - r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) - r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) - r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) - r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) - r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) - r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) - r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) - r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Sub(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _mul(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) * r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) * r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) * r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) * r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) * r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) * r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) * r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) * r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Mul(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

// todo make Euclidean division (unlike Go)
func _div(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) / r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) / r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) / r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) / r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) / r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) / r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) / r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) / r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Div(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _quo(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) / r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) / r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) / r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) / r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) / r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) / r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) / r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) / r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Quo(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _rem(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator add")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint8) % r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint16) % r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint32) % r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(uint64) % r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int8) % r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int16) % r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int32) % r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: l.GetConstant().(int64) % r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  l.GetParamType(),
			Const: big.NewInt(0).Quo(l.GetConstant().(*big.Int), r.GetConstant().(*big.Int)),
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func sub(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _sub(l, r)
}

func mul(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _mul(l, r)
}

func div(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _div(l, r)
}

func quo(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _quo(l, r)
}

func rem(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	li, _, _ := getBasicTraits(l)
	ri, _, _ := getBasicTraits(r)
	if !li || !ri {
		return nil, errors.New("operate on non-integral type")
	}
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _rem(l, r)
}

func le(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _le(l, r)
}

func lt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _lt(l, r)
}

func _lt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	b, err = _le(l, r)
	if err != nil {
		return nil, err
	}
	if b.GetConstant().(bool) {
		l, err = _eq(l, r)
		if err != nil {
			return nil, err
		}
		if l.GetConstant().(bool) {
			return &lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: false,
			}, nil
		} else {
			return b, nil
		}
	}
	return &lexer.ConstantVariable{
		Type:  value_type.Bool,
		Const: false,
	}, nil
}

func ge(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _ge(l, r)
}

func _ge(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	b, err = _lt(l, r)
	if err != nil {
		return nil, err
	}
	if b.GetConstant().(bool) {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: false,
		}, nil
	} else {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: true,
		}, nil
	}
}

func gt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _gt(l, r)
}

func _gt(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	b, err = _le(l, r)
	if err != nil {
		return nil, err
	}
	if b.GetConstant().(bool) {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: false,
		}, nil
	} else {
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: true,
		}, nil
	}
}

func eq(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != r.GetParamType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return _eq(l, r)
}

func land(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != value_type.Bool || r.GetParamType() != value_type.Bool {
		return nil, fmt.Errorf("type error: %v %v", l.GetParamType(), r.GetParamType())
	}
	return &lexer.ConstantVariable{Type:value_type.Bool, Const: l.GetConstant().(bool) && r.GetConstant().(bool)}, nil
}

func lor(l token.ConstantI, r token.ConstantI) (b token.ConstantI, err error) {
	if l.GetParamType() != value_type.Bool || r.GetParamType() != value_type.Bool {
		return nil, fmt.Errorf("type error: %v %v", l.GetParamType(), r.GetParamType())
	}
	return &lexer.ConstantVariable{Type:value_type.Bool, Const: l.GetConstant().(bool) || r.GetConstant().(bool)}, nil
}

func lnot(l token.ConstantI) (token.ConstantI, error) {
	if l.GetParamType() != value_type.Bool {
		return nil, fmt.Errorf("type error: %v", l.GetParamType())
	}
	return &lexer.ConstantVariable{Type:value_type.Bool, Const: !l.GetConstant().(bool)}, nil
}

func _eq(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64,
		value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64, value_type.Bool, value_type.String:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant() == r.GetConstant(),
		}, nil
	case value_type.Bytes:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: bytes.Equal(r.GetConstant().([]byte), r.GetConstant().([]byte)),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(*big.Int).Cmp(r.GetConstant().(*big.Int)) == 0,
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}

func _le(l token.ConstantI, r token.ConstantI) (token.ConstantI, error) {
	switch l.GetParamType() {
	case value_type.String, value_type.Bytes, value_type.Bool:
		return nil, errors.New("string/bytes/bool var has not operator le")
	case value_type.Uint8:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint8) <= r.GetConstant().(uint8),
		}, nil
	case value_type.Uint16:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint16) <= r.GetConstant().(uint16),
		}, nil
	case value_type.Uint32:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint32) <= r.GetConstant().(uint32),
		}, nil
	case value_type.Uint64:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(uint64) <= r.GetConstant().(uint64),
		}, nil
	case value_type.Int8:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int8) <= r.GetConstant().(int8),
		}, nil
	case value_type.Int16:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int16) <= r.GetConstant().(int16),
		}, nil
	case value_type.Int32:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int32) <= r.GetConstant().(int32),
		}, nil
	case value_type.Int64:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(int64) <= r.GetConstant().(int64),
		}, nil
	case value_type.Uint128, value_type.Uint256, value_type.Int256, value_type.Int128:
		return &lexer.ConstantVariable{
			Type:  value_type.Bool,
			Const: l.GetConstant().(*big.Int).Cmp(r.GetConstant().(*big.Int)) <= 0,
		}, nil
	default:
		return nil, errors.New("value type not found")
	}
}