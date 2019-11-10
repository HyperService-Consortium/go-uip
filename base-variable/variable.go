package base_variable

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type Variable struct {
	Type uiptypes.TypeID
	Value interface{}
}

func (v Variable) GetType() uiptypes.TypeID {
	return v.Type
}

func (v Variable) GetValue() interface{} {
	return v.Value
}
