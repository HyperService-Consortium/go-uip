package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

type PackPrototype interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte) (interface{}, error)
}

func Encode(variable abstraction.Ref) ([]byte, error) { return EncodeG(instance, variable) }

func DecodeR(x []byte, t abstraction.RefType) (interface{}, error) { return DecodeGR(instance, x, t) }

func EncodeG(g abstraction.GVM, variable abstraction.Ref) ([]byte, error) {
	return variable.Encode()
}

func DecodeGR(g abstraction.GVM, x []byte, t abstraction.RefType) (interface{}, error) {
	return g.DecodeRef(t, x)
}
