package lexer_types

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
	"io"
)

type BinaryExpression struct {
	Type  value_type.Type `json:"type"`
	Sign  sign_type.Type  `json:"sign"`
	Left  Param           `json:"left"`
	Right Param           `json:"right"`
}

func (b BinaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(b.Type)
}

func (b BinaryExpression) Determine(f DetermineContext) (_ token_types.Param, err error) {
	var db DeterminedBinaryExpression
	db.Left, err = b.Left.Determine(f)
	if err != nil {
		return nil, err
	}

	db.Right, err = b.Right.Determine(f)
	if err != nil {
		return nil, err
	}
	db.Type = b.Type
	db.Sign = b.Sign

	return &db, nil
}

type DeterminedBinaryExpression struct {
	Type  value_type.Type   `json:"type"`
	Sign  sign_type.Type    `json:"sign"`
	Left  token_types.Param `json:"left"`
	Right token_types.Param `json:"right"`
}

func (b DeterminedBinaryExpression) GetSign() sign_type.Type {
	return b.Sign
}

func (b DeterminedBinaryExpression) GetLeft() token_types.Param {
	return b.Left
}

func (b DeterminedBinaryExpression) GetRight() token_types.Param {
	return b.Right
}

func (b DeterminedBinaryExpression) GetGVMTok() gvm.TokType {
	return token_type.BinaryExpression
}

func (b DeterminedBinaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(b.Type)
}

func (b DeterminedBinaryExpression) Marshal(w io.Writer, err *error) {
	if *err != nil {
		return
	}
	serial.Write(w, b.Type, err)
	serial.Write(w, b.Sign, err)
	EncodeVTok(w, b.Left, err)
	EncodeVTok(w, b.Right, err)
}

func (b *DeterminedBinaryExpression) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	serial.Read(r, &b.Type, err)
	serial.Read(r, &b.Sign, err)
	DecodeVTok(r, &b.Left, err)
	DecodeVTok(r, &b.Right, err)
	*i = b
}

func (b DeterminedBinaryExpression) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
	l, err := b.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	r, err := b.Right.Eval(g)
	if err != nil {
		return nil, err
	}
	return BiCalc(l, r, gvm_type.SignType(b.Sign))
}

func BiCalc(l gvm.Ref, r gvm.Ref, signType gvm_type.SignType) (gvm.Ref, error) {
	if IsGVMNative(l) && IsGVMNative(r) {
		//todo: impl bitOp
		return gvm_type.BiCalc(l, r, signType)
	}
	return nil, errorn.NewRuntimeBiCalcError(l.GetGVMType(), r.GetGVMType(), signType)
}

func IsGVMNative(l gvm.Ref) bool {
	//todo upstream
	return gvm_type.IsStandardRefType(l.GetGVMType()) || l.GetGVMType() == gvm_type.RefUnknown
}
