package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"strconv"
)

func formatUintSuffix(v uint64) string {
	return strconv.FormatUint(v, 16)
}

const (
	fCurrentPC = "_gvm_cpc"
	fPCPrefix  = "_gvm_pc"
	fCurrentFN = "_gvm_cfn"
	fFNPrefix  = "_gvm_fn"
	fDepth     = "_gvm_dep"
)

// GetPC return the value of PC at depth `dep`
func GetPC(g abstraction.GVM, dep uint64) (uint64, error) {
	r, err := g.Load(fPCPrefix+formatUintSuffix(dep), RefUint64)
	if err != nil {
		return 0, err
	}
	return r.Unwrap().(uint64), nil
}

// GetCurrentPC return the value of current PC
func GetCurrentPC(g abstraction.GVM) (uint64, error) {
	r, err := g.Load(fCurrentPC, RefUint64)
	if err != nil {
		return 0, nil
	}
	return r.Unwrap().(uint64), nil
}

// GetFN return the function name at depth `dep`
func GetFN(g abstraction.GVM, dep uint64) (string, error) {
	r, err := g.Load(fFNPrefix+formatUintSuffix(dep), RefString)
	if err != nil {
		return "", err
	}
	return r.Unwrap().(string), nil
}

// GetCurrentFN return the value of current function name
func GetCurrentFN(g abstraction.GVM) (string, error) {
	r, err := g.Load(fCurrentFN, RefString)
	if err != nil {
		return "", nil
	}
	return r.Unwrap().(string), nil
}

// GetCurrentDepth return the value of current depth
func GetCurrentDepth(g abstraction.GVM) (uint64, error) {
	r, err := g.Load(fDepth, RefUint64)
	if err != nil {
		return 0, nil
	}
	return r.Unwrap().(uint64), nil
}

func setPC(g abstraction.GVM, dep, pc uint64) error {
	return g.Save(fPCPrefix+formatUintSuffix(dep), Uint64(pc))
}

func setCurrentPC(g abstraction.GVM, pc uint64) error {
	return g.Save(fCurrentPC, Uint64(pc))
}

func setFN(g abstraction.GVM, dep uint64, fn string) error {
	return g.Save(fFNPrefix+formatUintSuffix(dep), String(fn))
}

func setCurrentFN(g abstraction.GVM, fn string) error {
	return g.Save(fCurrentFN, String(fn))
}

func setCurrentDepth(g abstraction.GVM, dep uint64) error {
	return g.Save(fDepth, Uint64(dep))
}

func fetch(g abstraction.GVM, mainFunc string, pc uint64) (f abstraction.Function, inst abstraction.Instruction, err error) {

	f, err = g.GetFunction(mainFunc)
	if err != nil {
		return nil, nil, err
	}
	inst, err = f.Fetch(pc)
	if err != nil {
		return nil, nil, err
	}
	return
}
