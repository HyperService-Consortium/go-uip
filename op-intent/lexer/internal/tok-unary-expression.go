package internal

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

// Param.type == value_type.Bool
type UnaryExpression struct {
	Type value_type.Type `json:"type"`
	Sign sign_type.Type  `json:"sign"`
	Left Param           `json:"left"`
}

func (u UnaryExpression) GetGVMTok() gvm.TokType {
	return convertGVMTokType(token.UnaryExpression)
}

func (u UnaryExpression) Marshal(w io.Writer, err *error) {
	panic("implement me")
}

func (u UnaryExpression) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	panic("implement me")
}

func (u UnaryExpression) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	l, err := u.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	switch u.GetSign() {
	case sign_type.LNot:
		return gvm_type.LNot(l)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", u.GetSign())
	}
}

func (u UnaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(u.Type)
}

func (u UnaryExpression) GetType() token.Type {
	return token.UnaryExpression
}

func (u UnaryExpression) GetSign() sign_type.Type {
	return u.Sign
}

func (u UnaryExpression) GetLeft() token.Param {
	return u.Left
}

func (u UnaryExpression) GetParamType() value_type.Type {
	return u.Type
}

func (u UnaryExpression) Determine(f InstantiateAccountF) (_ Param, err error) {
	u.Left, err = u.Left.Determine(f)
	if err != nil {
		return nil, err
	}
	return u, nil
}
