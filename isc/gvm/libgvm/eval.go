package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

func Eval(p abstraction.VTok) (abstraction.Ref, error) {
	return EvalG(instance, p)
}

func EvalG(g abstraction.GVM, p abstraction.VTok) (abstraction.Ref, error) {
	return p.Eval(g)
}
