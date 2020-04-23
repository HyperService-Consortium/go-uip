package errorn

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

var (
	InvalidOpType                 = errors.New("there is at least an unexpected op_type in OpIntents")
	UnknownUnitType               = errors.New("unknown unit type")
	UnknownDependencyType         = errors.New("unknown dependency type")
	UnknownParam                  = errors.New("unknown param type")
	UnknownAccount                = errors.New("unknown account format")
	ErrNoAccountBaseProvided      = errors.New("uip.AccountBase must be provided")
	ErrNoBlockChainGetterProvided = errors.New("uip.BlockChainGetter must be provided")
	ErrNotEnoughParamInformation  = errors.New("not enough param information")
	ErrTypeError                  = errors.New("unexpected field type")
	ErrNoDeterminedChainID        = errors.New("no determined chain id")
	ErrNotTranslated              = errors.New("not translated")
	ErrNegativeIndex              = errors.New("index negative")
)

type ValueTypeNotFound struct {
	ValueType string
}

func (e ValueTypeNotFound) Error() string {
	return fmt.Sprintf("value type %v not found", e.ValueType)
}

type TokenTypeNotFound struct {
	TokenType int
}

func (e TokenTypeNotFound) Error() string {
	return fmt.Sprintf("token type %v not found", e.TokenType)
}

type GVMTypeNotFound struct {
	GVMType int
}

func (e GVMTypeNotFound) Error() string {
	return fmt.Sprintf("gvm type %v not found", e.GVMType)
}

type AccountTypeNotFound struct {
	AccountType int
}

func (e AccountTypeNotFound) Error() string {
	return fmt.Sprintf("account type %v not found", e.AccountType)
}

type AccountNotFound struct {
	Name    string
	ChainID uip.ChainIDUnderlyingType
}

func (e AccountNotFound) Error() string {
	return fmt.Sprintf(`account <"%v",%v> not found`, e.Name, e.ChainID)
}

type FieldNotFound struct {
	Field string
}

func (e FieldNotFound) Error() string {
	return fmt.Sprintf("field %v not found", e.Field)
}

type SignNotFound struct {
	Sign string
}

func (e SignNotFound) Error() string {
	return fmt.Sprintf("sign %v not found", e.Sign)
}

type OpNameNotFound struct {
	OpName string
}

func (e OpNameNotFound) Error() string {
	return fmt.Sprintf("op name %v not found", e.OpName)
}

type BiCalcError struct {
	L    gvm.RefType
	R    gvm.RefType
	Sign gvm_type.SignType
}

func (e BiCalcError) Error() string {
	return fmt.Sprintf("<L:%v,R:%v,Sign:%v>", e.L, e.R, e.Sign)
}

type UnCalcError struct {
	L    gvm.RefType
	Sign gvm_type.SignType
}

func (e UnCalcError) Error() string {
	return fmt.Sprintf("<L:%v,Sign:%v>", e.L, e.Sign)
}

type FunctionNotFound struct {
	FuncName string
}

func (e FunctionNotFound) Error() string {
	return fmt.Sprintf("function not found: %v", e.FuncName)
}

type TypeAssertionError struct {
	Expected gvm.RefType
	Actual   gvm.Ref
}

func (e TypeAssertionError) Error() string {
	return fmt.Sprintf("runtime type assertion error: expected %v, actual %v", e.Expected, e.Actual.GetGVMType())
}

type ChainIDNotEqual struct {
	U, V uip.ChainIDUnderlyingType
}

func (c ChainIDNotEqual) Error() string {
	return fmt.Sprintf("chain id not equal: %v and %v", c.U, c.V)
}
