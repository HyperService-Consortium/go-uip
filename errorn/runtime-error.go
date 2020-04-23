package errorn

import (
	"github.com/Myriad-Dreamin/catcher"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

const (
	RuntimeUnCalcError = iota
	RuntimeBiCalcError
	RuntimeFunctionNotFound
	RuntimeGVMExecuteError
	RuntimeDeserializeError
	RuntimeTypeAssertionError
	RuntimeEvalError
	RuntimeSaveLocalStateVariableError
	RuntimeNotTranslated
)

func _newRuntimeError(err error, errCode int) RuntimeError {
	return RuntimeError{catcher.WrapN(catcher.BaseSkip+2, errCode, err)}
}

func NewRuntimeBiCalcError(l, r gvm.RefType, signType gvm_type.SignType) RuntimeError {
	return _newRuntimeError(&BiCalcError{L: l, R: r, Sign: signType}, RuntimeBiCalcError)
}

func NewRuntimeUnCalcError(l gvm.RefType, signType gvm_type.SignType) RuntimeError {
	return _newRuntimeError(&UnCalcError{L: l, Sign: signType}, RuntimeUnCalcError)
}

func NewFunctionNotFound(funcName string) RuntimeError {
	return _newRuntimeError(&FunctionNotFound{FuncName: funcName}, RuntimeFunctionNotFound)
}

func NewGVMExecuteError(err error) RuntimeError {
	return _newRuntimeError(err, RuntimeGVMExecuteError)
}

func NewDeserializeError(err error) RuntimeError {
	return _newRuntimeError(err, RuntimeDeserializeError)
}

func NewRuntimeTypeAssertionError(expected gvm.RefType, actual gvm.Ref) RuntimeError {
	return _newRuntimeError(TypeAssertionError{
		Expected: expected, Actual: actual}, RuntimeTypeAssertionError)
}

func NewEvalError(err error) RuntimeError {
	return _newRuntimeError(err, RuntimeEvalError)
}

func NewSaveLocalStateVariableError(err error) RuntimeError {
	return _newRuntimeError(err, RuntimeSaveLocalStateVariableError)
}

func NewRuntimeNotTranslated() RuntimeError {
	return _newRuntimeError(ErrNotTranslated, RuntimeNotTranslated)
}

type RuntimeError struct {
	Err error
}

func (p RuntimeError) Error() string {
	return p.Err.Error()
}
