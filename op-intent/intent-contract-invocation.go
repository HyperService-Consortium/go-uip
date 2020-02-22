package opintent

import "encoding/base64"

type BaseContractInvocationOpIntent struct {
	Src      *RawAccountInfo `json:"invoker"`       // key
	Dst      string          `json:"contract_addr"` // key
	Code     []byte          `json:"contract_code"` // key
	FuncName string          `json:"func"`
	Params   []RawParam      `json:"parameters"`
	Amount   string          `json:"amount"` // option
	Meta     ResultI         `json:"meta"`
}

func (intent *BaseContractInvocationOpIntent) UnmarshalJSON(b []byte) error {
	r, err := NewGJSONResult(b)
	if err != nil {
		return err
	}
	return intent.UnmarshalResult(r)
}

func (intent *BaseContractInvocationOpIntent) UnmarshalResult(content ResultI) (err error) {

	invoker := content.Get(FieldOpIntentsInvoker)
	if !invoker.Exists() {
		return newFieldNotFound(FieldOpIntentsInvoker)
	}
	contractAddr := content.Get(FieldOpIntentsContractAddr)
	if !contractAddr.Exists() {
		return newFieldNotFound(FieldOpIntentsContractAddr)
	}
	contractCode := content.Get(FieldOpIntentsContractCode)
	fn := content.Get(FieldOpIntentsFunc)
	if !fn.Exists() {
		return newFieldNotFound(FieldOpIntentsFunc)
	}
	parameters := content.Get(FieldOpIntentsParameters)
	if !parameters.Exists() {
		return newFieldNotFound(FieldOpIntentsParameters)
	}
	amount := content.Get(FieldOpIntentsAmount)
	if !amount.Exists() {
		return newFieldNotFound(FieldOpIntentsAmount)
	}
	meta := content.Get(FieldOpIntentsMeta)

	intent.Src = new(RawAccountInfo)
	err = intent.Src.UnmarshalResult(invoker)
	if err != nil {
		return
	}
	intent.Dst = contractAddr.String()
	intent.Code, err = base64.StdEncoding.DecodeString(contractCode.String())
	if err != nil {
		return
	}
	intent.FuncName = fn.String()
	intent.Params, err = initParamsR(parameters)
	if err != nil {
		return
	}
	intent.Amount = amount.String()
	intent.Meta = meta

	return
}

