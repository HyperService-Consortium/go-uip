package libgvm

import "github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"

func saveFrame(g *abstraction.ExecCtx) error {
	err := setPC(g, g.Depth, g.PC)
	if err != nil {
		return err
	}
	err = setFN(g, g.Depth, g.FN)
	if err != nil {
		return err
	}
	err = saveLocals(g, g.Depth, g.This)
	if err != nil {
		return err
	}
	return nil
}

func loadFrame(g *abstraction.ExecCtx) (err error) {
	g.FN, err = GetFN(g, g.Depth)
	if err != nil {
		return err
	}
	g.PC, err = GetPC(g, g.Depth)
	if err != nil {
		return err
	}
	g.This = g.Parent
	if g.Depth > 0 {
		g.Parent, err = loadLocals(g, g.Depth-1)
		if err != nil {
			return err
		}
	} else {
		g.Parent = nil
	}
	g.Function, err = g.GetFunction(g.FN)
	if err != nil {
		return err
	}
	return nil
}

func deleteFrame(g *abstraction.ExecCtx) error {
	err := deletePC(g, g.Depth)
	if err != nil {
		return err
	}
	err = deleteFN(g, g.Depth)
	if err != nil {
		return err
	}
	err = deleteLocals(g, g.Depth)
	if err != nil {
		return err
	}
	return nil
}

func pushFrame(g *abstraction.ExecCtx, fn string) error {
	err := saveFrame(g)
	if err != nil {
		return err
	}
	g.Depth++
	g.FN, g.PC, g.Parent, g.This = fn, 0, g.This, make(abstraction.Locals)
	g.Function, err = g.GetFunction(fn)
	if err != nil {
		return err
	}
	return setCurrentState(g)
}

func popFrame(g *abstraction.ExecCtx) (err error) {
	if g.Depth == 0 {
		return StopUnderFlow
	}
	err = deleteFrame(g)
	if err != nil {
		return err
	}

	g.Depth--
	err = loadFrame(g)
	if err != nil {
		return err
	}

	return setCurrentState(g)
}
