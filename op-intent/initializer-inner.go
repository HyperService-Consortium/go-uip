package opintent

import (
	"encoding/hex"
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
)

func (ier *Initializer) unmarshal(b []byte, v interface{}) (err error) {
	err = json.Unmarshal(b, v)
	if err != nil {
		return errorn.NewUnmarshalError(err)
	}
	return
}

func (ier *Initializer) marshal(v interface{}) (b []byte, err error) {
	b, err = json.Marshal(v)
	if err != nil {
		return nil, errorn.NewMarshalError(err)
	}
	return
}
//marshal

func (ier *Initializer) decodeHex(src string) (b []byte, err error) {
	b, err = hex.DecodeString(src)
	if err != nil {
		return nil, errorn.NewUnmarshalError(err)
	}
	return
}
