package opintent

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

func (r *RawDependency) UnmarshalResult(res ResultI) error {
	var left, right, dep ResultI
	if left = res.Get("left"); !left.Exists() {
		return newFieldNotFound("left")
	}
	r.Src = left.String()
	if right = res.Get("right"); !right.Exists() {
		return newFieldNotFound("right")
	}
	r.Dst = right.String()
	if dep = res.Get("dep"); !dep.Exists() {
		return newFieldNotFound("dep")
	}

	switch dep.String() {
	case "before":
	case "after":
		r.Src, r.Dst =
			r.Dst, r.Src
	default:
		return newInvalidFieldError(unknownDependencyType)
	}
	return nil
}
