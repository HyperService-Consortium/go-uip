package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer/internal"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

func EncodeVTok(w io.Writer, v uip.VTok, err *error) {
	internal.EncodeVTok(w, v, err)
}

func DecodeVTok(r io.Reader, v *uip.VTok, err *error) {
	internal.DecodeVTok(r, v, err)
}

func DecodeVTokWithType(r io.Reader, v *uip.VTok, t gvm.TokType, err *error) {
	internal.DecodeVTokWithType(r, v, t, err)
}

func ReadConstant(r io.Reader, v *uip.VTok, t gvm.RefType, err *error) {
	internal.ReadConstant(r, v, t, err)
}
