package lexer

import (
	"encoding/base64"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
)

type InvokeIntent struct {
	*IntentImpl
	Src      Account           `json:"invoker"`       // key
	Dst      Account           `json:"contract_addr"` // key
	Code     []byte            `json:"contract_code"` // key
	FuncName string            `json:"func"`
	Params   []Param           `json:"parameters"`
	Amount   string            `json:"amount"` // option
	Meta     document.Document `json:"meta"`
}

func (intent *InvokeIntent) UnmarshalJSON(b []byte) error {
	r, err := document.NewGJSONDocument(b)
	if err != nil {
		return err
	}
	return intent.UnmarshalDocument(r)
}

func (intent *InvokeIntent) UnmarshalDocument(content document.Document) (err error) {

	invoker := content.Get(FieldOpIntentsInvoker)
	if !invoker.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsInvoker)
	}
	contractAddr := content.Get(FieldOpIntentsContractAddr)
	if !contractAddr.Exists() {
		a, err := AccountUnmarshalResult("name", content.Get(FieldOpIntentsContract))
		if err != nil {
			return err
		}
		intent.Dst = a
	} else {
		b, err := DecodeAddress(contractAddr.String())
		if err != nil {
			return err
		}
		intent.Dst = RawAccount{Address: b}
	}
	contractCode := content.Get(FieldOpIntentsContractCode)
	fn := content.Get(FieldOpIntentsFunc)
	if !fn.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsFunc)
	}
	parameters := content.Get(FieldOpIntentsParameters)
	if !parameters.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsParameters)
	}
	amount := content.Get(FieldOpIntentsAmount)
	if !amount.Exists() {
		intent.Amount = "00"
	} else {
		intent.Amount = amount.String()
	}
	meta := content.Get(FieldOpIntentsMeta)

	intent.Src, err = AccountUnmarshalResult(FieldOpIntentsUserName, invoker)
	if err != nil {
		return
	}
	intent.Code, err = base64.StdEncoding.DecodeString(contractCode.String())
	if err != nil {
		return
	}
	intent.FuncName = fn.String()
	intent.Params, err = initParamsR(parameters)
	if err != nil {
		return
	}
	intent.Meta = meta

	return
}
