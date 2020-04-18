package lexer_types

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
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

func (u UnaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(u.Type)
}

type DeterminedUnaryExpression struct {
	Type value_type.Type   `json:"type"`
	Sign sign_type.Type    `json:"sign"`
	Left token_types.Param `json:"left"`
}

func (u DeterminedUnaryExpression) GetGVMTok() gvm.TokType {
	return token_type.UnaryExpression
}

func (u DeterminedUnaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(u.Type)
}

func (u DeterminedUnaryExpression) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, u.Type, err)
	serial.Write(w, u.Sign, err)
	EncodeVTok(w, u.Left, err)
}

func (u *DeterminedUnaryExpression) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	if *err != nil {
		return
	}
	serial.Read(r, &u.Type, err)
	serial.Read(r, &u.Sign, err)
	DecodeVTok(r, &u.Left, err)
	*i = u
}

func (u *DeterminedUnaryExpression) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	l, err := u.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	switch u.Sign {
	case sign_type.LNot:
		return gvm_type.LNot(l)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", u.Sign)
	}
}

func (u UnaryExpression) Determine(f DetermineContext) (_ token_types.Param, err error) {
	var du DeterminedUnaryExpression
	du.Left, err = u.Left.Determine(f)
	if err != nil {
		return nil, err
	}
	du.Sign = u.Sign
	du.Type = u.Type
	return &du, nil
}