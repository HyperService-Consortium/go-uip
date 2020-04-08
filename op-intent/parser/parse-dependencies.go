package parser

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
)

// the LeftName intent is before RightName intent
type Dependency struct {
	Src int64
	Dst int64
}

type DependenciesInfo struct {
	Dependencies []Dependency
}

func (ier *Parser) InitDependencies(
	rawDeps lexer.RawDependenciesI, nameMap map[string]int) (
	deps *DependenciesInfo, err error) {
	deps = &DependenciesInfo{
		Dependencies: make([]Dependency, rawDeps.Len()),
	}

	for idx := range deps.Dependencies {
		rawDep := rawDeps.GetDependencies(idx)

		n := rawDep.GetSrc()
		if i, ok := nameMap[n]; !ok {
			return nil, errorn.NewOpNameNotFound(n)
		} else {
			deps.Dependencies[idx].Src = int64(i)
		}

		n = rawDep.GetDst()
		if i, ok := nameMap[n]; !ok {
			return nil, errorn.NewOpNameNotFound(n)
		} else {
			deps.Dependencies[idx].Dst = int64(i)
		}
	}
	return
}
