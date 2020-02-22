package opintent

type RawAccountInfo struct {
	ChainId uint64 `json:"domain"`
	Name    string `json:"user_name"`
}

func (r *RawAccountInfo) UnmarshalResult(i ResultI) error {
	domain := i.Get(FieldOpIntentsDomain)
	if !domain.Exists() {
		return newFieldNotFound(FieldOpIntentsDomain)
	}
	name := i.Get(FieldOpIntentsUserName)
	if !name.Exists() {
		return newFieldNotFound(FieldOpIntentsUserName)
	}

	r.ChainId = domain.Uint()
	r.Name = name.String()

	return nil
}
