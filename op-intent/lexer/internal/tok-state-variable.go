package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type StateVariable struct {
	Type     value_type.Type `json:"type"`
	Contract Account         `json:"contract"`
	Pos      []byte          `json:"pos"`
	Field    []byte          `json:"field"`
}

func (s StateVariable) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, s.Type, err)
	EncodeAccount(w, s.Contract, err)
	serial.Write(w, s.Pos, err)
	serial.Write(w, s.Field, err)
}

func (s *StateVariable) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &s.Type, err)
	DecodeAccount(r, &s.Contract, err)
	serial.Read(r, &s.Pos, err)
	serial.Read(r, &s.Field, err)
	*i = s
}

func (s StateVariable) GetGVMTok() gvm.TokType {
	return token.StateVariable
}

func (s StateVariable) GetGVMType() gvm.RefType {
	return gvm.RefType(s.Type)
}

func (s StateVariable) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	//return g.Load(string(p.Field), p.GetGVMType())
	// todo
	panic("todo")
}

func (s *StateVariable) Determine(c DetermineContext) (token.Param, error) {
	a, err := c.QueryContract(s.Contract)
	if err != nil {
		return nil, err
	}
	s.Contract = NewNamespacedRawAccount(a)
	return s, nil
}
