package gvm

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/builtin"
)

func Eval(p VTok) (Ref, error) {
	return EvalG(instance, p)
}

func EvalG(g GVM, p VTok) (Ref, error) {

	if i, err := g.GetTypePrototype(p.GetGVMTok()); err != nil {
		return nil, err
	} else {
		return i.Eval(g, p)
	}
}

var TObjs = [...]TypePrototype{
	TokConstant:           (*_Constant)(nil),
	TokBinaryExpression:   (*_BinaryExpression)(nil),
	TokUnaryExpression:    (*_UnaryExpression)(nil),
	TokLocalStateVariable: (*_LocalStateVariable)(nil),
	TokStateVariable:      (*_NotImplT)(nil),
}

type _Constant struct{}

func (*_Constant) Eval(g GVM, p VTok) (Ref, error) {
	return p.(Ref), nil
}

type _LocalStateVariable struct {
}

func (l *_LocalStateVariable) Eval(g GVM, p VTok) (Ref, error) {
	v := p.(LocalStateVariable)
	return g.Load(v.GetFieldGVM(), v.GetGVMType())
}

type _BinaryExpression struct{}

func (*_BinaryExpression) Eval(g GVM, p VTok) (Ref, error) {
	v := p.(BinaryExpression)
	l, err := EvalG(g, v.GetLeftTok())
	if err != nil {
		return nil, err
	}
	r, err := EvalG(g, v.GetRightTok())
	if err != nil {
		return nil, err
	}
	switch v.GetSign() {
	case sign_type.EQ:
		return eq(l, r)
	case sign_type.LE:
		return le(l, r)
	case sign_type.LT:
		return lt(l, r)
	case sign_type.GE:
		return ge(l, r)
	case sign_type.GT:
		return gt(l, r)
	case sign_type.LAnd:
		return land(l, r)
	case sign_type.LOr:
		return lor(l, r)
	case sign_type.ADD:
		return add(l, r)
	case sign_type.SUB:
		return sub(l, r)
	case sign_type.MUL:
		return mul(l, r)
	case sign_type.QUO:
		return quo(l, r)
	case sign_type.REM:
		return rem(l, r)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", v.GetSign())
	}
}

type _UnaryExpression struct{}

func (*_UnaryExpression) Eval(g GVM, p VTok) (Ref, error) {
	v := p.(BinaryExpression)
	l, err := EvalG(g, v.GetLeftTok())
	if err != nil {
		return nil, err
	}
	switch v.GetSign() {
	case sign_type.LNot:
		return lnot(l)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", v.GetSign())
	}
}

type _StateVariable struct{}

//var StateVariableT = (*_StateVariable)(nil)
func (*_StateVariable) Eval(g GVM, p VTok) (Ref, error) {
	return nil, errors.New("todo")
}

type _NotImplT struct{}

func (*_NotImplT) Eval(g GVM, p VTok) (Ref, error) {
	return nil, fmt.Errorf("not implemented type to eval: %v", p.GetGVMTok())
}


func arithmeticCal(protoCal func(interface{}, interface{}) interface{}, l Ref, r Ref) (b Ref, err error) {
	if l, r, err = convArithmetic(l, r); err != nil {
		return nil, err
	}
	return ref{t: l.GetGVMType(), v: unsafeShrinkType(protoCal(
		extendType(l), extendType(r)), l.GetGVMType())}, nil
}

func arithmeticCompare(protoCal func(interface{}, interface{}) bool, l Ref, r Ref) (b Ref, err error) {

	if l, r, err = convArithmetic(l, r); err != nil {
		return nil, err
	}
	return ref{t: RefBool, v: protoCal(extendType(l), extendType(r))}, nil
}

func booleanCal(protoCal func(bool, bool) bool, l Ref, r Ref) (b Ref, err error) {
	x, err := toBool(l)
	if err != nil {
		return nil, err
	}
	y, err := toBool(r)
	if err != nil {
		return nil, err
	}
	return ref{t: RefBool, v: protoCal(x, y)}, nil
}

func add(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCal(builtin.Add, l, r)
}

func sub(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCal(builtin.Sub, l, r)
}

func mul(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCal(builtin.Mul, l, r)
}

func div(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCal(builtin.Div, l, r)
}

func quo(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCal(builtin.Quo, l, r)
}

func rem(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCal(builtin.Rem, l, r)
}

func le(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCompare(builtin.LessEqual, l, r)
}

func lt(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCompare(builtin.Less, l, r)
}

func ge(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCompare(builtin.GreaterEqual, l, r)
}

func gt(l Ref, r Ref) (b Ref, err error) {
	return arithmeticCompare(builtin.Greater, l, r)
}

func eq(l Ref, r Ref) (b Ref, err error) {
	if l.GetGVMType() != r.GetGVMType() {
		l, r, err = tryConv(l, r)
		if err != nil {
			return nil, err
		}
	}
	return ref{t: RefBool, v: builtin.Equal(extendType(l), extendType(r))}, nil
}

func land(l Ref, r Ref) (b Ref, err error) {
	return booleanCal(builtin.LogicAnd, l, r)
}

func lor(l Ref, r Ref) (b Ref, err error) {
	return booleanCal(builtin.LogicOr, l, r)
}

func lnot(l Ref) (Ref, error) {
	x, err := toBool(l)
	if err != nil {
		return nil, err
	}
	return ref{
		t:  RefBool,
		v: x,
	}, nil
}

