package opintent

type RawDependencies struct {
	dependencies []RawDependency
}

func (r RawDependencies) Len() int {
	return len(r.dependencies)
}

func (r RawDependencies) GetDependencies(i int) RawDependencyI {
	return &r.dependencies[i]
}

func (ier *Initializer) InitDependencies(rawDeps [][]byte) (deps *RawDependencies, err error) {
	deps = &RawDependencies{
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

func (ier *Initializer) InitDependenciesR(source ResultI) (deps *RawDependencies, err error) {
	if source.Exists() && !source.IsArray() {
		return nil, newInvalidFieldError(ErrTypeError).Desc(AtOpIntentField{Field: FieldOpIntentsParameters})
	}
	rawDeps := source.Array()

	deps = &RawDependencies{
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

