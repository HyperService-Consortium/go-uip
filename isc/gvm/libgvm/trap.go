package libgvm

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"strconv"
)

type Trap interface {
	DoTrap(g *abstraction.ExecCtx) error
}

type trapCallFunc struct {
	FN    string             `json:"fn"`
	Left  []string           `json:"left"`
	Right []abstraction.VTok `json:"right"`
}

func (c trapCallFunc) Error() string {
	return fmt.Sprintf("trap calling: %v", c.FN)
}

type FuncParam struct {
	T abstraction.RefType
	K int
}

func (f FuncParam) GetGVMTok() abstraction.TokType {
	return TokFuncParam
}

func (f FuncParam) GetGVMType() abstraction.RefType {
	return f.T
}

func (f FuncParam) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	return GetParam(g, f.K), nil
}

func GetParam(g *abstraction.ExecCtx, k int) abstraction.Ref {
	return g.This[strconv.Itoa(k)]
}

func FuncParamName(k int) string {
	return strconv.Itoa(k)
}

func FuncReturnName(g *abstraction.ExecCtx, k int) string {
	return g.This["_gvm_return"+strconv.Itoa(k)].Unwrap().(string)
}

func GetReturn(g *abstraction.ExecCtx, k int) abstraction.Ref {
	return g.Parent[g.This["_gvm_return"+strconv.Itoa(k)].Unwrap().(string)]
}

func (c trapCallFunc) Exec(g *abstraction.ExecCtx) error {
	g.PC++
	return c
}

func (c trapCallFunc) DoTrap(g *abstraction.ExecCtx) (err error) {
	var refs = make([]abstraction.Ref, len(c.Right))
	for l := range c.Right {
		refs[l], err = c.Right[l].Eval(g)
		if err != nil {
			return err
		}
	}
	err = pushFrame(g, c.FN)
	if err != nil {
		return err
	}
	for l := range c.Right {
		g.This[strconv.Itoa(l)] = refs[l]
	}
	for l := range c.Left {
		g.This["_gvm_return"+strconv.Itoa(l)] = String(c.Left[l])
	}
	return
}
