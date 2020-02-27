package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
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

func (r *RawDependency) UnmarshalResult(res document.Document) error {
	var sub document.Document
	if sub = res.Get(FieldKeyLeft); !sub.Exists() {
		return errorn.NewFieldNotFound(FieldKeyLeft)
	}
	r.Src = sub.String()
	if sub = res.Get(FieldKeyRight); !sub.Exists() {
		return errorn.NewFieldNotFound(FieldKeyRight)
	}
	r.Dst = sub.String()

	switch res.Get(FieldDependencyDep).String() {
	case "before", "":
	case "after":
		r.Src, r.Dst =
			r.Dst, r.Src
	default:
		return errorn.NewInvalidFieldError(errorn.UnknownDependencyType)
	}
	return nil
}
