package internal

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	gvm_type "github.com/Myriad-Dreamin/gvm/libgvm/gvm-type"
)

func incPC(g *gvm.ExecCtx) error {
	g.PC++
	return nil
}

func cond(g *gvm.ExecCtx, c gvm.VTok,
	ifFunc func(g *gvm.ExecCtx) error, elseFunc func(g *gvm.ExecCtx) error) error {
	v, err := c.Eval(g)
	if err != nil {
		return errorn.NewEvalError(err)
	}
	if uip.TypeIDUnderlyingType(v.GetGVMType()) != value_type.Bool {
		return errorn.NewRuntimeTypeAssertionError(gvm_type.RefBool, v)
	}
	if v.Unwrap().(bool) {
		if ifFunc != nil {
			return ifFunc(g)
		}
	} else if elseFunc != nil {
		return elseFunc(g)
	}

	return nil
}
