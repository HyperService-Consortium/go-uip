package gvm_builtin

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
)

func ArithmeticalError(l, r interface{}) error {
	return fmt.Errorf("runtime error: cant do arithemetic operation on type %v and %v", reflect.TypeOf(l), reflect.TypeOf(r))
}

func Add(l, r interface{}) interface{} {
	switch l.(type) {
	case uint64:
		return l.(uint64) + r.(uint64)
	case int64:
		return l.(int64) + r.(int64)
	case uint:
		return l.(uint) + r.(uint)
	case int:
		return l.(int) + r.(int)
	case *big.Int:
		return big.NewInt(0).Add(l.(*big.Int), r.(*big.Int))
	}
	panic(ArithmeticalError(l, r))
}

func Sub(l, r interface{}) interface{} {
	switch l.(type) {
	case uint64:
		return l.(uint64) - r.(uint64)
	case int64:
		return l.(int64) - r.(int64)
	case uint:
		return l.(uint) - r.(uint)
	case int:
		return l.(int) - r.(int)
	case *big.Int:
		return big.NewInt(0).Sub(l.(*big.Int), r.(*big.Int))
	}
	panic(ArithmeticalError(l, r))
}

func Mul(l, r interface{}) interface{} {
	switch l.(type) {
	case uint64:
		return l.(uint64) + r.(uint64)
	case int64:
		return l.(int64) + r.(int64)
	case uint:
		return l.(uint) + r.(uint)
	case int:
		return l.(int) + r.(int)
	case *big.Int:
		return big.NewInt(0).Add(l.(*big.Int), r.(*big.Int))
	}
	panic(ArithmeticalError(l, r))
}

// todo make Euclidean division (unlike Go)
func Div(l, r interface{}) interface{} {
	switch l.(type) {
	case uint64:
		return l.(uint64) / r.(uint64)
	case int64:
		return l.(int64) / r.(int64)
	case uint:
		return l.(uint) / r.(uint)
	case int:
		return l.(int) / r.(int)
	case *big.Int:
		return big.NewInt(0).Div(l.(*big.Int), r.(*big.Int))
	}
	panic(ArithmeticalError(l, r))
}

func Quo(l, r interface{}) interface{} {
	switch l.(type) {
	case uint64:
		return l.(uint64) / r.(uint64)
	case int64:
		return l.(int64) / r.(int64)
	case uint:
		return l.(uint) / r.(uint)
	case int:
		return l.(int) / r.(int)
	case *big.Int:
		return big.NewInt(0).Quo(l.(*big.Int), r.(*big.Int))
	}
	panic(ArithmeticalError(l, r))
}

func Rem(l, r interface{}) interface{} {
	switch l.(type) {
	case uint64:
		return l.(uint64) % r.(uint64)
	case int64:
		return l.(int64) % r.(int64)
	case uint:
		return l.(uint) % r.(uint)
	case int:
		return l.(int) % r.(int)
	case *big.Int:
		return big.NewInt(0).Rem(l.(*big.Int), r.(*big.Int))
	}
	panic(ArithmeticalError(l, r))
}

func LessEqual(l, r interface{}) bool {
	switch l.(type) {
	case uint64:
		return l.(uint64) <= r.(uint64)
	case int64:
		return l.(int64) <= r.(int64)
	case uint:
		return l.(uint) <= r.(uint)
	case int:
		return l.(int) <= r.(int)
	case *big.Int:
		return l.(*big.Int).Cmp(r.(*big.Int)) <= 0
	}
	panic(ArithmeticalError(l, r))
}

func Less(l, r interface{}) bool {
	switch l.(type) {
	case uint64:
		return l.(uint64) < r.(uint64)
	case int64:
		return l.(int64) < r.(int64)
	case uint:
		return l.(uint) < r.(uint)
	case int:
		return l.(int) < r.(int)
	case *big.Int:
		return l.(*big.Int).Cmp(r.(*big.Int)) < 0
	}
	panic(ArithmeticalError(l, r))
}

func Greater(l, r interface{}) bool {
	switch l.(type) {
	case uint64:
		return l.(uint64) > r.(uint64)
	case int64:
		return l.(int64) > r.(int64)
	case uint:
		return l.(uint) > r.(uint)
	case int:
		return l.(int) > r.(int)
	case *big.Int:
		return l.(*big.Int).Cmp(r.(*big.Int)) > 0
	}
	panic(ArithmeticalError(l, r))
}

func GreaterEqual(l, r interface{}) bool {
	switch l.(type) {
	case uint64:
		return l.(uint64) >= r.(uint64)
	case int64:
		return l.(int64) >= r.(int64)
	case uint:
		return l.(uint) >= r.(uint)
	case int:
		return l.(int) >= r.(int)
	case *big.Int:
		return l.(*big.Int).Cmp(r.(*big.Int)) >= 0
	}
	panic(ArithmeticalError(l, r))
}

func Equal(l, r interface{}) bool {
	switch l.(type) {
	case bool:
		return l.(bool) == r.(bool)
	case string:
		return l.(string) == r.(string)
	case []byte:
		return bytes.Equal(l.([]byte), r.([]byte))
	case uint64:
		return l.(uint64) == r.(uint64)
	case int64:
		return l.(int64) == r.(int64)
	case uint:
		return l.(uint) == r.(uint)
	case int:
		return l.(int) == r.(int)
	case *big.Int:
		return l.(*big.Int).Cmp(r.(*big.Int)) == 0
	}
	panic(ArithmeticalError(l, r))
}

func LogicAnd(a, b bool) bool { return a && b }
func LogicOr(a, b bool) bool  { return a || b }
func LogicNot(a bool) bool    { return !a }
