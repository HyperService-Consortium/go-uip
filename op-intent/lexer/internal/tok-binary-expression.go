package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

type BinaryExpression struct {
	Type  value_type.Type `json:"type"`
	Sign  sign_type.Type  `json:"sign"`
	Left  Param           `json:"left"`
	Right Param           `json:"right"`
}

func (b BinaryExpression) GetLeftTok() gvm.VTok {
	return b.Left
}

func (b BinaryExpression) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	l, err := b.GetLeftTok().Eval(g)
	if err != nil {
		return nil, err
	}
	r, err := b.GetRightTok().Eval(g)
	if err != nil {
		return nil, err
	}
	return gvm_type.BiCalc(l, r, gvm_type.SignType(b.Sign))
}

func (b BinaryExpression) GetRightTok() gvm.VTok {
	return b.Right
}

func (b BinaryExpression) GetGVMTok() gvm.TokType {
	return convGVMTokType(token.BinaryExpression)
}

func (b BinaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(b.Type)
}
func (b BinaryExpression) GetType() token.Type {
	return token.BinaryExpression
}

func (b BinaryExpression) GetSign() sign_type.Type {
	return b.Sign
}

func (b BinaryExpression) GetLeft() token.Param {
	return b.Left
}

func (b BinaryExpression) GetRight() token.Param {
	return b.Right
}

func (b BinaryExpression) GetParamType() value_type.Type {
	return b.Type
}

func (b BinaryExpression) Determine(f InstantiateAccountF) (_ Param, err error) {
	b.Left, err = b.Left.Determine(f)
	if err != nil {
		return nil, err
	}

	b.Right, err = b.Right.Determine(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}
