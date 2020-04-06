package internal

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/Myriad-Dreamin/gvm"
)

func convertGVMTokType(t token.Type) gvm.TokType {
	return gvm.TokType(t - token.Constant)
}
