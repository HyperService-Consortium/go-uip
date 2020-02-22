package opintent

type RawDependenciesInfo struct {
	dependencies []RawDependency
}

func (r RawDependenciesInfo) Len() int {
	return len(r.dependencies)
}

func (r RawDependenciesInfo) GetDependencies(i int) RawDependencyI {
	return &r.dependencies[i]
}

func (ier *Initializer) InitDependencies(rawDeps [][]byte) (deps *RawDependenciesInfo, err error) {
	deps = &RawDependenciesInfo{
		dependencies: make([]RawDependency, len(rawDeps)),
	}
	var res ResultI
	for idx, dependency := range rawDeps {
		res, err = NewGJSONResult(dependency)
		if err != nil {
			return
		}
		err = deps.dependencies[idx].UnmarshalResult(res)
		if err != nil {
			return
		}
	}
	return
}

func (ier *Initializer) InitDependenciesR(source ResultI) (deps *RawDependenciesInfo, err error) {
	rawContents := source.Get(FieldOpIntentsDependencies)
	if ! rawContents.Exists() {
		return nil, newFieldNotFound(FieldOpIntentsDependencies)
	}
	rawDeps := rawContents.Array()

	deps = &RawDependenciesInfo{
		dependencies: make([]RawDependency, rawDeps.Len()),
	}
	for idx := 0; idx < rawDeps.Len(); idx++ {
		err = deps.dependencies[idx].UnmarshalResult(rawDeps.Index(idx))
		if err != nil {
			return
		}
	}
	return
}

