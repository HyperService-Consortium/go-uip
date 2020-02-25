package opintent

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
)

type mcs struct{}
type ts struct {}
type _serializer struct {
	TransactionIntent ts
	Meta              struct {
		Contract mcs
	}
}

var Serializer = _serializer{}

func (mcs) Unmarshal(b []byte, meta *lexer.ContractInvokeMeta) error {
	return json.Unmarshal(b, meta)
}

func (mcs) Marshal(meta *lexer.ContractInvokeMeta) ([]byte, error) {
	return json.Marshal(meta)
}

func (ts) Unmarshal(b []byte, meta *TransactionIntent) error {
	return json.Unmarshal(b, meta)
}

func (ts) Marshal(meta *TransactionIntent) ([]byte, error) {
	return json.Marshal(meta)
}

