package errorn

import (
	"errors"
	"fmt"
)

var (
	InvalidOpType                 = errors.New("there is at least an unexpected op_type in OpIntents")
	UnknownUnitType               = errors.New("unknown unit type")
	UnknownDependencyType         = errors.New("unknown dependency type")
	UnknownParam = errors.New("unknown param type")
	UnknownAccount = errors.New("unknown account format")
	ErrNoAccountBaseProvided      = errors.New("uip.AccountBase must be provided")
	ErrNoBlockChainGetterProvided = errors.New("uip.BlockChainGetter must be provided")
	ErrNotEnoughParamInformation  = errors.New("not enough param information")
	ErrTypeError = errors.New("unexpected field type")
)


type ValueTypeNotFound struct {
	ValueType string
}

func (e ValueTypeNotFound) Error() string {
	return fmt.Sprintf("value type %v not found", e.ValueType)
}

type FieldNotFound struct {
	Field string
}

func (e FieldNotFound) Error() string {
	return fmt.Sprintf("field %v not found", e.Field)
}

type OpNameNotFound struct {
	OpName string
}

func (e OpNameNotFound) Error() string {
	return fmt.Sprintf("op name %v not found", e.OpName)
}
