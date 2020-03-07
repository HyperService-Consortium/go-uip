package libgvm

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

func incPC(g *abstraction.ExecCtx) error {
	g.PC++
	return nil
}

type Goto struct {
	Index uint64 `json:"goto"`
}

func (inst *Goto) Exec(g *abstraction.ExecCtx) error {
	g.PC = inst.Index
	return nil
}

func cond(g *abstraction.ExecCtx, c abstraction.VTok,
	ifFunc func(g *abstraction.ExecCtx) error, elseFunc func(g *abstraction.ExecCtx) error) error {
	v, err := c.Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != RefBool {
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

type ConditionGoto struct {
	Goto
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionGoto) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.Goto.Exec, incPC)
}

type SetState struct {
	Target          string           `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetState) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	err = g.Save(G.Target, k)
	if err != nil {
		return err
	}

	g.PC++
	return nil
}

type ConditionSetState struct {
	SetState
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetState) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetState.Exec, incPC)
}

type SetLocal struct {
	Target          string           `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetLocal) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	g.This[G.Target] = k
	g.PC++
	return nil
}

type ConditionSetLocals struct {
	SetLocal
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetLocals) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetLocal.Exec, incPC)
}

type SetParentLocals struct {
	Target          string           `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetParentLocals) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	g.Parent[G.Target] = k
	g.PC++
	return nil
}

type ConditionSetParentLocals struct {
	SetParentLocals
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetParentLocals) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetParentLocals.Exec, incPC)
}

type SetFuncReturn struct {
	Target          int              `json:"target"`
	RightExpression abstraction.VTok `json:"expression"`
}

func (G SetFuncReturn) Exec(g *abstraction.ExecCtx) error {
	k, err := G.RightExpression.Eval(g)
	if err != nil {
		return err
	}
	g.Parent[FuncReturnName(g, G.Target)] = k
	g.PC++
	return nil
}

type ConditionSetFuncReturn struct {
	SetFuncReturn
	Condition abstraction.VTok `json:"condition"`
}

func (inst *ConditionSetFuncReturn) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.SetFuncReturn.Exec, incPC)
}

type CallFunc = trapCallFunc

type ConditionCallFunc struct {
	CallFunc
	Condition abstraction.VTok `json:"condition"`
}

func (inst ConditionCallFunc) Exec(g *abstraction.ExecCtx) error {
	return cond(g, inst.Condition, inst.CallFunc.Exec, incPC)
}
