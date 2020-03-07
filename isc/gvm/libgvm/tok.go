package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

const (
	TokConstant abstraction.TokType = iota
	TokStateVariable
	TokLocalStateVariable
	TokBinaryExpression
	TokUnaryExpression

	TokFuncParam
	TokLocalVariable
)

type stateVariable interface {
	GetFieldGVM() string
}

type LocalStateVariable interface {
	abstraction.VTok
	stateVariable
}

type Constant = abstraction.Ref

type UnaryExpression interface {
	abstraction.VTok
	GetSign() sign_type.Type
	GetLeftTok() abstraction.VTok
}

type BinaryExpression interface {
	UnaryExpression
	GetRightTok() abstraction.VTok
}

type LocalVariable struct {
	Name string
	Type abstraction.RefType
}

func (l LocalVariable) GetGVMTok() abstraction.TokType {
	return TokLocalVariable
}

func (l LocalVariable) GetGVMType() abstraction.RefType {
	return l.Type
}

func (l LocalVariable) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	return g.This[l.Name], nil
}
