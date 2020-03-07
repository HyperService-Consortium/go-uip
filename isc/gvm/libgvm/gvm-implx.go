package libgvm

import "github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"

type impl = GVM
type GVMeX struct {
	impl
	g *Mem
}

func NewGVM() (*GVMeX, error) {
	g, err := NewMem()
	if err != nil {
		return nil, err
	}
	return &GVMeX{
		impl: GVM{Machine: g},
		g:    g,
	}, nil
}

func (i *GVMeX) AddFunction(fn string, instructions []abstraction.Instruction) error {
	if i.g.InstSet == nil {
		i.g.InstSet = make(map[string]abstraction.Function)
	}

	i.g.InstSet[fn] = FunctionImpl(instructions)
	return nil
}
