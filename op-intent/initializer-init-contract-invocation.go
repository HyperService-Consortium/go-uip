package opintent

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/standard"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type BaseContractInvocationOpIntent struct {
	Src      *RawAccountInfo `json:"invoker"`       // key
	Dst      string          `json:"contract_addr"` // key
	Code     []byte          `json:"contract_code"` // key
	FuncName string          `json:"func"`
	Params   []uip.RawParam  `json:"parameters"`
	Amount   string          `json:"amount"` // option
	Meta     json.RawMessage `json:"meta"`
}

func (ier *Initializer) initContractInvocation(info *RawIntent, content []byte) (
	sub interface{}, err error) {
	var intent = new(BaseContractInvocationOpIntent)
	sub = intent
	err = ier.unmarshal(content, intent)
	if err != nil {
		return
	}

	if intent.Src == nil {
		return nil, newFieldNotFound("src")
	}
	if err = standard.CheckValidHexString(intent.Dst); err != nil {
		return nil, newInvalidFieldError(err).Desc(AtOpIntentField{"dst"})
	}

	if len(intent.Amount) == 0 {
		intent.Amount = "00"
	}
	if err = standard.CheckValidHexString(intent.Amount); err != nil {
		return nil, newInvalidFieldError(err).Desc(AtOpIntentField{"amount"})
	}

	if len(intent.FuncName) == 0 {
		return nil, newFieldNotFound("function name")
	}
	return
}
