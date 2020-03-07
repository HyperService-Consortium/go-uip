package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

type GVM struct {
	abstraction.Machine
}

func (i *GVM) Continue() error {
	return Continue(i.Machine)
}

func (i *GVM) Step() error {
	return Step(i.Machine)
}

func (i *GVM) Run(fn string) error {
	return Run(i.Machine, fn)
}
