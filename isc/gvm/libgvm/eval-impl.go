package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/gvm-builtin"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

func Add(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Add, l, r)
}

func Sub(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Sub, l, r)
}

func Mul(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Mul, l, r)
}

func div(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Div, l, r)
}

func Quo(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Quo, l, r)
}

func Rem(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCal(gvm_builtin.Rem, l, r)
}

func LE(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.LessEqual, l, r)
}

func LT(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.Less, l, r)
}

func GE(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.GreaterEqual, l, r)
}

func GT(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return arithmeticCompare(gvm_builtin.Greater, l, r)
}

func EQ(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	if l.GetGVMType() != r.GetGVMType() {
		l, r, err = TryAlignArithmetic(l, r)
		if err != nil {
			return nil, err
		}
	}
	return Bool(gvm_builtin.Equal(ExtendType(l), ExtendType(r))), nil
}

func LAnd(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return booleanCal(gvm_builtin.LogicAnd, l, r)
}

func LOr(l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	return booleanCal(gvm_builtin.LogicOr, l, r)
}

func LNot(l abstraction.Ref) (abstraction.Ref, error) {
	x, err := ToBool(l)
	if err != nil {
		return nil, err
	}
	return Bool(x), nil
}

func arithmeticCal(protoCal func(interface{}, interface{}) interface{}, l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {
	if l, r, err = TryAlignArithmetic(l, r); err != nil {
		return nil, err
	}
	return UnsafeShrinkType(protoCal(ExtendType(l), ExtendType(r)), l.GetGVMType())
}

func arithmeticCompare(protoCal func(interface{}, interface{}) bool, l abstraction.Ref, r abstraction.Ref) (b abstraction.Ref, err error) {

	if l, r, err = TryAlignArithmetic(l, r); err != nil {
		return nil, err
	}
	return Bool(protoCal(ExtendType(l), ExtendType(r))), nil
}

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
