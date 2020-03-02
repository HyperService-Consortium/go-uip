package gvm

import (
	"errors"
	"fmt"
)

type Mem struct {
	Base
	Context map[string]Ref
	InstSet map[string]Function
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
		g: g,
	}, nil
}


var instance GVM = new(Mem)
func (b *Mem) Save(field string, k Ref) error {
	if b.Context == nil {
		b.Context = make(map[string]Ref)
	}

	b.Context[field] = k
	return nil
}

var Undefined = ref{t: Unknown}

func (b *Mem) Load(field string, t RefType) (Ref, error) {
	if b.Context == nil {
		b.Context = make(map[string]Ref)
	}

	if r, ok := b.Context[field] ; ok {
		if r.GetGVMType() != t {
			return nil, fmt.Errorf("type error: %v -\\> %v", r.GetGVMType(), t)
		}
		return r, nil
	} else {
		return Undefined, nil
	}
}

type FunctionImpl []Instruction
type pureLazyInstruction struct {
	Instruction
}

func (p pureLazyInstruction) DeserializeGVMI() (Instruction, error) {
	return p.Instruction, nil
}

func (f FunctionImpl) LazyFetch(pc uint64) (LazyInstruction, error) {
	return pureLazyInstruction{f[pc]}, nil
}

func (f FunctionImpl) Fetch(pc uint64) (Instruction, error) {
	return f[pc], nil
}

func (f FunctionImpl) Len() int {
	return len(f)
}

func (b *Mem) GetFunction(functionName string) (Function, error) {
	if b.InstSet == nil {
		b.InstSet = make(map[string]Function)
	}
	f, ok := b.InstSet[functionName]
	if !ok {
		return nil, errors.New("not found")
	}
	return f, nil
}

