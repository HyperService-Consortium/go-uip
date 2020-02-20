package opintent

import (
	"errors"
	"fmt"
)

var (
	invalidOpType                 = errors.New("there is at least an unexpected op_type in OpIntents")
	unknownUnitType               = errors.New("unknown unit type")
	unknownDependencyType         = errors.New("unknown dependency type")
	ErrNoAccountBaseProvided      = errors.New("uip.AccountBase must be provided")
	ErrNoBlockChainGetterProvided = errors.New("uip.BlockChainGetter must be provided")
	ErrNotEnoughParamInformation  = errors.New("not enough param information")
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
