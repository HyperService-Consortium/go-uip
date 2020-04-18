package uip

import "github.com/Myriad-Dreamin/gvm"

type Variable interface {
	gvm.Ref
	GetType() TypeID
	GetValue() interface{}
}
