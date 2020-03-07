package libgvm

import "github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"

type Trap interface {
	DoTrap(g abstraction.Machine, fn *string, pc *uint64) error
}

type TrapCallFunc struct {
	newFn string
}

func (c TrapCallFunc) DoTrap(g abstraction.Machine, fn *string, pc *uint64) {

}
