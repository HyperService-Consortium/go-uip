package libgvm

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

type Trap interface {
	DoTrap(g *abstraction.ExecCtx) error
}

type TrapCallFunc struct {
	NewFn string
}

func (c TrapCallFunc) Error() string {
	return fmt.Sprintf("trap calling: %v", c.NewFn)
}

func (c TrapCallFunc) DoTrap(g *abstraction.ExecCtx) error {
	return pushFrame(g, c.NewFn)
}
