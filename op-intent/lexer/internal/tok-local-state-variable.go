package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/Myriad-Dreamin/gvm"
)

type LocalStateVariable struct {
	Type  value_type.Type `json:"type"`
	Pos   []byte          `json:"pos"`
	Field []byte          `json:"field"`
}

func (l LocalStateVariable) GetGVMTok() gvm.TokType {
	return convGVMTokType(token.LocalStateVariable)
}

func (l LocalStateVariable) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	return g.Load(string(l.Field), l.GetGVMType())
}

func (l LocalStateVariable) GetGVMType() gvm.RefType {
	return gvm.RefType(l.Type)
}

func (l LocalStateVariable) GetType() token.Type {
	return token.LocalStateVariable
}

func (l LocalStateVariable) GetParamType() value_type.Type {
	return l.Type
}

func (l LocalStateVariable) GetPos() []byte {
	return l.Pos
}

func (l LocalStateVariable) GetField() []byte {
	return l.Field
}

func (l LocalStateVariable) Determine(f InstantiateAccountF) (Param, error) {
	return l, nil
}
