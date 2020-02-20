package opintent

import (
	"encoding/hex"
	"encoding/json"
)

func (ier *Initializer) unmarshal(b []byte, v interface{}) (err error) {
	err = json.Unmarshal(b, v)
	if err != nil {
		return newUnmarshalError(err)
	}
	return
}

func (ier *Initializer) marshal(v interface{}) (b []byte, err error) {
	b, err = json.Marshal(v)
	if err != nil {
		return nil, newMarshalError(err)
	}
	return
}
//marshal

func (ier *Initializer) decodeHex(src string) (b []byte, err error) {
	b, err = hex.DecodeString(src)
	if err != nil {
		return nil, newUnmarshalError(err)
	}
	return
}
