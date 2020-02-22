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
	var sub ResultI
	if sub = res.Get(FieldDependencyLeft); !sub.Exists() {
		return newFieldNotFound(FieldDependencyLeft)
	}
	r.Src = sub.String()
	if sub = res.Get(FieldDependencyRight); !sub.Exists() {
		return newFieldNotFound(FieldDependencyRight)
	}
	r.Dst = sub.String()

	switch res.Get(FieldDependencyDep).String() {
	case "before", "":
	case "after":
		r.Src, r.Dst =
			r.Dst, r.Src
	default:
		return newInvalidFieldError(unknownDependencyType)
	}
	return nil
}
