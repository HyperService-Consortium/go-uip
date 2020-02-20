package uip

type Variable interface {
	GetType() TypeID
	GetValue() interface{}
}
