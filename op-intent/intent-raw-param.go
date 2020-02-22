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

	r.Type = content.Get(FieldOpIntentsType).String()
	r.Value = content.Get(FieldOpIntentsValue)
	return
}



func initParamsR(i ResultI) (params []RawParam, err error) {
	rawParams := i.Array()
	params = make([]RawParam, rawParams.Len())
	for i := 0; i < rawParams.Len(); i++ {
		err = params[i].UnmarshalResult(rawParams.Index(i))
		if err != nil {
			return
		}
	}
	return
}
