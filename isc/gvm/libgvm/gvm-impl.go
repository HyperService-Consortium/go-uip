package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

type GVM struct {
	abstraction.Machine
}

func (i *GVM) Continue(fn string, pc uint64) (uint64, error) {
	return Continue(i.Machine, fn, pc)
}

func (i *GVM) Step(fn string, pc uint64) (uint64, error) {
	return Step(i.Machine, fn, pc)
}

func (i *GVM) Run(fn string) (uint64, error) {
	return Run(i.Machine, fn)
}
