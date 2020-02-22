package opintent

import (
	"github.com/HyperService-Consortium/go-uip/standard"
)

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

func (ier *Initializer) initContractInvocationR(info *RawIntent, content ResultI) (
	sub interface{}, err error) {
	var intent = new(BaseContractInvocationOpIntent)
	sub = intent
	err = intent.UnmarshalResult(content)
	if err != nil {
		return nil, err
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
