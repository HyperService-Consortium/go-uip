package opintent

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/errorn"
)

func (ier *Initializer) unmarshal(b []byte, v interface{}) (err error) {
	err = json.Unmarshal(b, v)
	if err != nil {
		return errorn.NewUnmarshalError(err)
	}
	return
}
