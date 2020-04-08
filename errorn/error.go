package errorn

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uip"
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
)

type ValueTypeNotFound struct {
	ValueType string
}

func (e ValueTypeNotFound) Error() string {
	return fmt.Sprintf("value type %v not found", e.ValueType)
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
