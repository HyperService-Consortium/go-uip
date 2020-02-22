package opintent

import (
	"encoding/json"
)

type BaseIfStatementOpIntent struct {
	If        []json.RawMessage `json:"if"`        // key
	Condition json.RawMessage   `json:"condition"` // key
	Else      []json.RawMessage `json:"else"`      // option
}

func (ier *Initializer) initIfStatement(info *RawIntent, content []byte) (sub interface{}, err error) {
	var intent = new(BaseIfStatementOpIntent)
	sub = intent
	err = ier.unmarshal(content, intent)
	if err != nil {
		return
	}

	return
}
