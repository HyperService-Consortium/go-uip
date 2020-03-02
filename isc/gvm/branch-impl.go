package gvm

import "fmt"

type Goto interface {
	GetGotoIndexGVMI() uint64
}

type SetState interface {
	GetRefNameGVMI() string
	GetRightHandStatementGVMI() VTok
}

type Condition interface {
	GetGotoConditionGVMI() VTok
}

type ConditionGoto interface {
	Condition
	Goto
}

type ConditionSetState interface {
	Condition
	SetState
}


type _Goto struct{}

func (_ *_Goto) Exec(g *ExecCtx, i Instruction) error {
	g.PC = uint64(i.(Goto).GetGotoIndexGVMI())
	return nil
}


type _ConditionGoto struct{}

func (_ *_ConditionGoto) Exec(g *ExecCtx, i Instruction) error {
	inst := i.(ConditionGoto)
	v, err := EvalG(g, inst.GetGotoConditionGVMI())
	if err != nil {
		return err
	}
	if v.GetGVMType() != RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}
	if v.Unwrap().(bool) {
		g.PC = uint64(inst.GetGotoIndexGVMI())
		return nil
	}
	g.PC++
	return nil
}

type _ConditionSetState struct{}

func (_ *_ConditionSetState) Exec(g *ExecCtx, i Instruction) error {
	inst := i.(ConditionSetState)
	v, err := EvalG(g, inst.GetGotoConditionGVMI())
	if err != nil {
		return err
	}
	if v.GetGVMType() != RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}

	if v.Unwrap().(bool) {
		k, err := EvalG(g, inst.GetRightHandStatementGVMI())
		if err != nil {
			return err
		}
		err = g.Save(inst.GetRefNameGVMI(), k)
		if err != nil {
			return err
		}
	}

	g.PC++
	return nil
}

type _SetState struct{}

func (_ *_SetState) Exec(g *ExecCtx, i Instruction) error {
	inst := i.(SetState)
	k, err := EvalG(g, inst.GetRightHandStatementGVMI())
	if err != nil {
		return err
	}
	err = g.Save(inst.GetRefNameGVMI(), k)
	if err != nil {
		return err
	}

	g.PC++
	return nil
}

