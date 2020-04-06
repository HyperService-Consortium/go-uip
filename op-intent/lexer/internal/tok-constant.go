package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type ConstantVariable struct {
	Type  value_type.Type `json:"type"`
	Const interface{}     `json:"constant"`
}

func (p ConstantVariable) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (p ConstantVariable) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	panic("implement me")
}

func (p ConstantVariable) Encode() ([]byte, error) {
	panic("todo")
}

func (p ConstantVariable) GetGVMTok() gvm.TokType {
	return convertGVMTokType(token.Constant)
}

func (p ConstantVariable) GetGVMType() gvm.RefType {
	return gvm.RefType(p.Type)
}

func (p ConstantVariable) Eval(_ *gvm.ExecCtx) (gvm.Ref, error) {
	return p, nil
}

func (p ConstantVariable) Unwrap() interface{} {

	return p.Const
}

func (p ConstantVariable) Determine(_ InstantiateAccountF) (Param, error) {
	return p, nil
}

func (p ConstantVariable) GetType() token.Type {
	return token.Constant
}

func (p ConstantVariable) GetParamType() value_type.Type {
	return p.Type
}
