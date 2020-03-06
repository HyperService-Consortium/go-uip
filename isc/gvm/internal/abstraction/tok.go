package abstraction

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
)

type TokType = uint64

const (
	TokConstant TokType = iota
	TokStateVariable
	TokLocalStateVariable
	TokBinaryExpression
	TokUnaryExpression
)

type Tok interface {
	GetGVMTok() TokType
}

type stateVariable interface {
	GetFieldGVM() string
}

type LocalStateVariable interface {
	VTok
	stateVariable
}

type Constant = Ref

type UnaryExpression interface {
	VTok
	GetSign() sign_type.Type
	GetLeftTok() VTok
}

type BinaryExpression interface {
	UnaryExpression
	GetRightTok() VTok
}

//
