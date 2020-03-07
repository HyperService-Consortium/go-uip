package libgvm

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

type Mem struct {
	Context map[string]abstraction.Ref
	InstSet map[string]abstraction.Function
}

func (g *Mem) CreateRef(t abstraction.RefType, v interface{}) abstraction.Ref {
	return CreateRef(t, v)
}

func (g *Mem) DecodeRef(t abstraction.RefType, r []byte) (abstraction.Ref, error) {
	return DecodeRef(t, r)
}

func NewMem() (*Mem, error) {
	return new(Mem), nil
}

func (g *Mem) Save(field string, k abstraction.Ref) error {
	if g.Context == nil {
		g.Context = make(map[string]abstraction.Ref)
	}

	g.Context[field] = k
	return nil
}

func (g *Mem) Load(field string, t abstraction.RefType) (abstraction.Ref, error) {
	if g.Context == nil {
		g.Context = make(map[string]abstraction.Ref)
	}

	if r, ok := g.Context[field]; ok {
		if r.GetGVMType() != t {
			return nil, fmt.Errorf("type error: %v -\\> %v", r.GetGVMType(), t)
		}
		return r, nil
	} else {
		return Undefined, nil
	}
}

func (g *Mem) GetFunction(functionName string) (abstraction.Function, error) {
	if g.InstSet == nil {
		g.InstSet = make(map[string]abstraction.Function)
	}
	f, ok := g.InstSet[functionName]
	if !ok {
		return nil, errors.New("not found")
	}
	return f, nil
}
