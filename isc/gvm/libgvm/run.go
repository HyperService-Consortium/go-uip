package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

func Step(g abstraction.Machine, fn string, pc uint64) (uint64, error) {
	f, inst, err := fetch(g, fn, pc)
	if err != nil {
		return 0, err
	}
	c := abstraction.ExecCtx{Machine: g, Function: f, PC: pc}
	return c.PC, inst.Exec(&c)
}

func Continue(g abstraction.Machine, fn string, pc uint64) (uint64, error) {
	f, err := g.GetFunction(fn)
	c := abstraction.ExecCtx{Machine: g, Function: f, PC: pc}
	for ; err == nil; err = Iter(&c) {
	}

	return c.PC, err
}

//TrapCallFunc
func Run(g abstraction.Machine, fn string) (uint64, error) {
	err := pushFrame(g, fn)
	var pc uint64 = 0
	for err == nil {
		pc, err = Continue(g, fn, pc)
		if err != nil {
			return 0, err
		}

		if err == OutOfRange {
			pc, fn, err = popFrame(g)
		} else if trap, ok := err.(Trap); ok {
			err = trap.DoTrap(g, &fn, &pc)
		}
	}
	return pc, err
}

func Iter(g *abstraction.ExecCtx) (err error) {
	if err = validate(g); err != nil {
		return
	}
	inst, err := g.Fetch(g.PC)
	if err != nil {
		return err
	}
	return inst.Exec(g)
}

func validate(g *abstraction.ExecCtx) error {
	if g.PC >= uint64(g.Len()) {
		return OutOfRange
	}
	return nil
}
