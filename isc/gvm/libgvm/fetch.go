package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

func fetch(g abstraction.GVM, mainFunc string, pc uint64) (f abstraction.Function, inst abstraction.Instruction, err error) {

	f, err = g.GetFunction(mainFunc)
	if err != nil {
		return nil, nil, err
	}
	inst, err = f.Fetch(pc)
	if err != nil {
		return nil, nil, err
	}
	return
}
