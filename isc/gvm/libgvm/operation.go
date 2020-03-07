package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/gvm-builtin"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

// Add implements arithmetical addition
func Add(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Add, l, r)
}

// Sub implements arithmetical subtraction
func Sub(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Sub, l, r)
}

// Mul implements arithmetical multiplication
func Mul(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Mul, l, r)
}

// Div implements arithmetical Euclidean division
func Div(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Div, l, r)
}

// Quo implements arithmetical truncated division
func Quo(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Quo, l, r)
}

// Rem implements arithmetical truncated modulus
func Rem(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Rem, l, r)
}

// LE implements arithmetical less equal (<=) operation
func LE(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.LessEqual, l, r)
}

// LT implements arithmetical less than (<) operation
func LT(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.Less, l, r)
}

// GE implements arithmetical greater equal (>=) operation
func GE(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.GreaterEqual, l, r)
}

// GT implements arithmetical greater than (>) operation
func GT(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.Greater, l, r)
}

// EQ implements arithmetical equal (==) operation
func EQ(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	if l.GetGVMType() != r.GetGVMType() {
		l, r, err = TryAlignArithmetic(l, r)
		if err != nil {
			return nil, err
		}
	}
	return Bool(gvm_builtin.Equal(ExtendType(l), ExtendType(r))), nil
}

// NEQ implements arithmetical not equal (!=) operation
func NEQ(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	if l.GetGVMType() != r.GetGVMType() {
		l, r, err = TryAlignArithmetic(l, r)
		if err != nil {
			return nil, err
		}
	}
	return Bool(!gvm_builtin.Equal(ExtendType(l), ExtendType(r))), nil
}

// LAnd implements logical and
func LAnd(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return booleanCal(gvm_builtin.LogicAnd, l, r)
}

// LOr implements logical or
func LOr(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return booleanCal(gvm_builtin.LogicOr, l, r)
}

// LNot implements logical not
func LNot(l abstraction.Ref) (abstraction.Ref, error) {
	x, err := ToBool(l)
	if err != nil {
		return nil, err
	}
	return Bool(!x), nil
}

// arithmeticCal performs arithmetic on two reference
func arithmeticCal(protoCal func(interface{}, interface{}) interface{}, l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {

	if l, r, err = TryAlignArithmetic(l, r); err != nil {
		return nil, err
	}
	return UnsafeShrinkType(protoCal(ExtendType(l), ExtendType(r)), l.GetGVMType())
}

// arithmeticCompare performs arithmetical comparison on two reference
func arithmeticCompare(protoCal func(interface{}, interface{}) bool, l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {

	if l, r, err = TryAlignArithmetic(l, r); err != nil {
		return nil, err
	}
	return Bool(protoCal(ExtendType(l), ExtendType(r))), nil
}

// booleanCal performs boolean operation on two reference
func booleanCal(protoCal func(bool, bool) bool, l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	x, err := ToBool(l)
	if err != nil {
		return nil, err
	}
	y, err := ToBool(r)
	if err != nil {
		return nil, err
	}
	return Bool(protoCal(x, y)), nil
}
