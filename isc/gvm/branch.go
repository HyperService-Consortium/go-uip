package gvm

import (
	"errors"
)

func getFunc(g GVM, mainFunc string, pc uint64) (f Function, inst LazyInstruction, err error) {

	f, err = g.GetFunction(mainFunc)
	if err != nil {
		return nil, nil, err
	}
	inst, err = f.LazyFetch(pc)
	if err != nil {
		return nil, nil, err
	}
	return
}

func InitPC(g GVM, mainFunc string, pc uint64) (uint64, error) {
	f, inst, err := getFunc(g, mainFunc, pc)
	if err != nil {
		return 0, err
	}
	switch inst.GetGVMIType() {
	case InstGoto, InstConditionSetState,
	InstConditionGoto, InstSetState:
		c := ExecCtx{GVM: g, Function: f, PC: pc}
		return c.PC, c._iter(inst)
	default:
		return pc, nil
	}
}

func NextPC(g GVM, fn string, pc uint64) (uint64, error) {
	f, inst, err := getFunc(g, fn, pc)
	if err != nil {
		return 0, err
	}
	c := ExecCtx{GVM: g, Function: f, PC: pc}
	return c.PC, c._iter(inst)
}

func ConsumePC(g GVM, fn string, pc uint64) (uint64, error) {

	f, err := g.GetFunction(fn)
	c := ExecCtx{GVM: g, Function: f, PC: pc}
	for ; err == nil; err = Iter(&c) {}
	return c.PC, err
}

func GetCurrentFnThread(g GVM) (string, error) {

}

func GetCurrentPC(g GVM) (int64, error) {

}

func pushThread(g GVM, fn string) error {

}

func popThread(g GVM) error {

}

func Run(g GVM) (uint64, error) {

}

func Iter(g *ExecCtx)  (err error) {
	if err = g.validate(); err != nil {
		return
	}
	inst, err := g.LazyFetch(g.PC)
	if err != nil {
		return err
	}
	return g._iter(inst)
}

type ExecCtx struct {
	GVM
	Function
	PC uint64
}

var Trap = errors.New("gvm trapped")

var OutOfRange = errors.New("gvm stopped")


func (g *ExecCtx) _iter(instruction LazyInstruction) error {
	if p, err := g.GetInstPrototype(instruction.GetGVMIType()); err != nil {
		return err
	} else {
		i, err := instruction.DeserializeGVMI()
		if err != nil {
			return err
		}
		return p.Exec(g, i)
	}
}

func (g *ExecCtx) validate() error {
	if g.PC >= uint64(g.Len()) {
		return OutOfRange
	}
	return nil
}
