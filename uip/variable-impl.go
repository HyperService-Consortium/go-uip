package uip

type VariableImpl struct {
	Type  TypeID
	Value interface{}
}

func (v VariableImpl) GetType() TypeID {
	return v.Type
}

func (v VariableImpl) GetValue() interface{} {
	return v.Value
}
