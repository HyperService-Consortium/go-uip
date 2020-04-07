package internal

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
)

func incPC(g *gvm.ExecCtx) error {
	g.PC++
	return nil
}

func cond(g *gvm.ExecCtx, c gvm.VTok,
	ifFunc func(g *gvm.ExecCtx) error, elseFunc func(g *gvm.ExecCtx) error) error {
	v, err := c.Eval(g)
	if err != nil {
		return err
	}
	if uip.TypeIDUnderlyingType(v.GetGVMType()) != value_type.Bool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
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
