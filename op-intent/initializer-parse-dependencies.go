package opintent

// the LeftName intent is before RightName intent
type Dependency struct {
	Src int64
	Dst int64
}

type DependenciesInfo struct {
	dependencies []Dependency
}

type RawDependenciesI interface {
	Len() int
	GetDependencies(i int) RawDependencyI
}

type RawDependencyI interface {
	GetSrc() string
	GetDst() string
}

func (ier *Initializer) ParseDependencies(
	rawDeps RawDependenciesI, nameMap map[string]int) (
	deps *DependenciesInfo, err error) {
	deps = &DependenciesInfo{
		dependencies: make([]Dependency, rawDeps.Len()),
	}

	for idx, _ := range deps.dependencies {
		rawDep := rawDeps.GetDependencies(idx)

		n := rawDep.GetSrc()
		if i, ok := nameMap[n]; !ok {
			return nil, newOpNameNotFound(n)
		} else {
			deps.dependencies[idx].Src = int64(i)
		}

		n = rawDep.GetDst()
		if i, ok := nameMap[n]; !ok {
			return nil, newOpNameNotFound(n)
		} else {
			deps.dependencies[idx].Dst = int64(i)
		}
	}
	return
}
