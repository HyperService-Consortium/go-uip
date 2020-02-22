package opintent

type RawParam struct {
	Type  string  `json:"Type"`
	Value ResultI `json:"Value"`
}

func (r *RawParam) UnmarshalJSON(b []byte) (err error) {
	c, err := NewGJSONResult(b)
	if err != nil {
		return
	}
	return r.UnmarshalResult(c)
}

func (r *RawParam) UnmarshalResult(content ResultI) (err error) {
	r.Value, content = content.Get(FieldOpIntentsValue), content.Get(FieldOpIntentsType)
	if !r.Value.Exists() {
		return newFieldNotFound(FieldOpIntentsValue)
	}
	if !content.Exists() {
		return newFieldNotFound(FieldOpIntentsType)
	}
	r.Type = content.String()
	return
}


func initParamsR(i ResultI) (params []RawParam, err error) {
	if i.Exists() && !i.IsArray() {
		return nil, newInvalidFieldError(ErrTypeError).Desc(AtOpIntentField{Field: FieldOpIntentsParameters})
	}
	rawParams := i.Array()
	params = make([]RawParam, rawParams.Len())
	for i := 0; i < rawParams.Len(); i++ {
		err = params[i].UnmarshalResult(rawParams.Index(i))
		if err != nil {
			return nil, err.(*ParseError).Desc(AtOpIntentParameterPos{i})
		}
	}
	return
}
