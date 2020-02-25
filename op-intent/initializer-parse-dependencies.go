package opintent

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
)

// the LeftName intent is before RightName intent
type Dependency struct {
	Src int64
	Dst int64
}

type DependenciesInfo struct {
	dependencies []Dependency
}

func (ier *Initializer) InitDependencies(
	rawDeps lexer.RawDependenciesI, nameMap map[string]int) (
	deps *DependenciesInfo, err error) {
	deps = &DependenciesInfo{
		dependencies: make([]Dependency, rawDeps.Len()),
	}

	for idx, _ := range deps.dependencies {
		rawDep := rawDeps.GetDependencies(idx)

		n := rawDep.GetSrc()
		if i, ok := nameMap[n]; !ok {
			return nil, errorn.NewOpNameNotFound(n)
		} else {
			deps.dependencies[idx].Src = int64(i)
		}

		n = rawDep.GetDst()
		if i, ok := nameMap[n]; !ok {
			return nil, errorn.NewOpNameNotFound(n)
		} else {
			deps.dependencies[idx].Dst = int64(i)
		}
	}
	return
}
