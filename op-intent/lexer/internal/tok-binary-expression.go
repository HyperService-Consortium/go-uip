package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/serial"
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

type DeterminedBinaryExpression struct {
	Type  value_type.Type `json:"type"`
	Sign  sign_type.Type  `json:"sign"`
	Left  token.Param     `json:"left"`
	Right token.Param     `json:"right"`
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

func (b DeterminedBinaryExpression) Unmarshal(r io.Reader, i *uip.VTok, err *error) {
	serial.Read(r, &b.Type, err)
	serial.Read(r, &b.Sign, err)
	DecodeVTok(r, &b.Left, err)
}

func (b DeterminedBinaryExpression) GetLeftTok() gvm.VTok {
	return b.Left
}

func (b DeterminedBinaryExpression) Eval(g *gvm.ExecCtx) (gvm.Ref, error) {
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

func (b DeterminedBinaryExpression) GetRightTok() gvm.VTok {
	return b.Right
}

func (b DeterminedBinaryExpression) GetGVMTok() gvm.TokType {
	return token.BinaryExpression
}

func (b DeterminedBinaryExpression) GetGVMType() gvm.RefType {
	return gvm.RefType(b.Type)
}
func (b DeterminedBinaryExpression) GetType() token.Type {
	return token.BinaryExpression
}

func (b DeterminedBinaryExpression) GetSign() sign_type.Type {
	return b.Sign
}

func (b DeterminedBinaryExpression) GetLeft() token.Param {
	return b.Left
}

func (b DeterminedBinaryExpression) GetRight() token.Param {
	return b.Right
}

func (b DeterminedBinaryExpression) GetParamType() value_type.Type {
	return b.Type
}

func (b BinaryExpression) GetParamType() value_type.Type {
	return b.Type
}

func (b BinaryExpression) Determine(f InstantiateAccountF) (_ token.Param, err error) {
	var db DeterminedBinaryExpression
	db.Left, err = b.Left.Determine(f)
	if err != nil {
		return nil, err
	}

	db.Right, err = b.Right.Determine(f)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
