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

func (b2 *Mem) CreateRef(t abstraction.RefType, v interface{}) abstraction.Ref {
	panic("implement me")
}

func (b2 *Mem) DecodeRef(t abstraction.RefType, b []byte) (abstraction.Ref, error) {
	panic("implement me")
}

func NewMemGVM() (*Mem, error) {
	return new(Mem), nil
}

func NewGVM() (*ImplX, error) {
	g, err := NewMemGVM()
	if err != nil {
		return nil, err
	}
	return &ImplX{
		impl: Impl{g},
		g:    g,
	}, nil
}

var instance abstraction.GVM = new(Mem)

func (b *Mem) Save(field string, k abstraction.Ref) error {
	if b.Context == nil {
		b.Context = make(map[string]abstraction.Ref)
	}

	b.Context[field] = k
	return nil
}

type Unknown struct{}

func (u *Unknown) GetGVMType() abstraction.RefType {
	return RefUnknown
}

func (u *Unknown) Unwrap() interface{} {
	return nil
}

func (u *Unknown) Encode() ([]byte, error) {
	return nil, nil
}

var Undefined *Unknown

func (b *Mem) Load(field string, t abstraction.RefType) (abstraction.Ref, error) {
	if b.Context == nil {
		b.Context = make(map[string]abstraction.Ref)
	}

	if r, ok := b.Context[field]; ok {
		if r.GetGVMType() != t {
			return nil, fmt.Errorf("type error: %v -\\> %v", r.GetGVMType(), t)
		}
		return r, nil
	} else {
		return Undefined, nil
	}
}

func (b *Mem) GetFunction(functionName string) (abstraction.Function, error) {
	if b.InstSet == nil {
		b.InstSet = make(map[string]abstraction.Function)
	}
	f, ok := b.InstSet[functionName]
	if !ok {
		return nil, errors.New("not found")
	}
	return f, nil
}

type FunctionImpl []abstraction.Instruction

func (f FunctionImpl) Fetch(pc uint64) (abstraction.Instruction, error) {
	return f[pc], nil
}

func (f FunctionImpl) Len() int {
	return len(f)
}
