package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/Myriad-Dreamin/gvm"
)

type StateVariable struct {
	Type     value_type.Type `json:"type"`
	Contract Account         `json:"contract"`
	Pos      []byte          `json:"pos"`
	Field    []byte          `json:"field"`
}

func (s StateVariable) GetGVMTok() gvm.TokType {
	return convGVMTokType(token.StateVariable)
}

func (s StateVariable) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	//return g.Load(string(p.Field), p.GetGVMType())
	// todo
	panic("todo")
}

func (s StateVariable) GetGVMType() gvm.RefType {
	return gvm.RefType(s.Type)
}

func (s StateVariable) GetType() token.Type {
	return token.StateVariable
}

func (s StateVariable) GetParamType() value_type.Type {
	return s.Type
}

func (s StateVariable) GetContract() token.Token {
	return s.Contract
}

func (s StateVariable) GetPos() []byte {
	return s.Pos
}

func (s StateVariable) GetField() []byte {
	return s.Field
}

func (s StateVariable) Determine(f InstantiateAccountF) (Param, error) {
	a, err := f(s.Contract)
	if err != nil {
		return nil, err
	}
	s.Contract = NewNamespacedRawAccount(a)
	return s, nil
}
