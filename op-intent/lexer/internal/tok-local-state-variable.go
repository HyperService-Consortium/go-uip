package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type LocalStateVariable struct {
	Type  value_type.Type `json:"type"`
	Pos   []byte          `json:"pos"`
	Field []byte          `json:"field"`
}

func (l LocalStateVariable) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, l.Type, err)
	serial.Write(w, l.Pos, err)
	serial.Write(w, l.Field, err)
}

func (l LocalStateVariable) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &l.Type, err)
	serial.Read(r, &l.Pos, err)
	serial.Read(r, &l.Field, err)
	*i = l
}

func (l LocalStateVariable) GetGVMTok() gvm.TokType {
	return token.LocalStateVariable
}

func (l LocalStateVariable) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	return g.Load(string(l.Field), l.GetGVMType())
}

func (l LocalStateVariable) GetGVMType() gvm.RefType {
	return gvm.RefType(l.Type)
}

func (l LocalStateVariable) Determine(_ InstantiateAccountF) (token.Param, error) {
	return l, nil
}
