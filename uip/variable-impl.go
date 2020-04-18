package uip

import "github.com/Myriad-Dreamin/gvm"

type VariableImpl struct {
	Type  TypeID
	Value interface{}
}

func (v VariableImpl) GetGVMType() gvm.RefType {
	return gvm.RefType(v.Type)
}

func (v VariableImpl) Unwrap() interface{} {
	return v.Value
}

func (v VariableImpl) Encode() ([]byte, error) {
	panic("implement me")
}

func (v VariableImpl) GetType() TypeID {
	return v.Type
}

func (v VariableImpl) GetValue() interface{} {
	return v.Value
}
