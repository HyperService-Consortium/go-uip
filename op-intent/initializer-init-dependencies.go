package opintent

import (
	"github.com/tidwall/gjson"
)

// the LeftName intent is before RightName intent
type RawDependency struct {
	Src string
	Dst string
}

func (r *RawDependency) GetSrc() string {
	return r.Src
}

func (r *RawDependency) GetDst() string {
	return r.Dst
}

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
	var res, left, right, dep gjson.Result
	for idx, dependency := range rawDeps {
		res = gjson.ParseBytes(dependency)
		if left = res.Get("left"); !left.Exists() {
			return nil, newFieldNotFound("left")
		}
		deps.dependencies[idx].Src = left.String()
		if right = res.Get("right"); !right.Exists() {
			return nil, newFieldNotFound("right")
		}
		deps.dependencies[idx].Dst = right.String()
		if dep = res.Get("dep"); !dep.Exists() {
			return nil, newFieldNotFound("dep")
		}

		switch dep.String() {
		case "before":
		case "after":
			deps.dependencies[idx].Src, deps.dependencies[idx].Dst =
				deps.dependencies[idx].Dst, deps.dependencies[idx].Src
		default:
			return nil, newInvalidFieldError(unknownDependencyType)
		}
	}
	return
}
