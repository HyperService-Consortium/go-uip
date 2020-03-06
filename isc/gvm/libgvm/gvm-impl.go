package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

type Impl struct {
	abstraction.GVM
}

func (i *Impl) Decode(x []byte, t abstraction.RefType) (abstraction.Ref, error) {
	return i.GVM.DecodeRef(t, x)
}

func (i *Impl) Continue(fn string, pc uint64) (uint64, error) {
	return Continue(i.GVM, fn, pc)
}

func (i *Impl) Step(fn string, pc uint64) (uint64, error) {
	return Step(i.GVM, fn, pc)
}

type impl = Impl
type ImplX struct {
	impl
	g *Mem
}

func (i *ImplX) AddFunction(fn string, instructions []abstraction.Instruction) error {
	if i.g.InstSet == nil {
		i.g.InstSet = make(map[string]abstraction.Function)
	}

	i.g.InstSet[fn] = FunctionImpl(instructions)
	return nil
}
