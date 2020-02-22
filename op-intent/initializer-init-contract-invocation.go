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

	if err = checkContractInvocationFields(intent); err != nil {
		return nil, err
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
	if err = checkContractInvocationFields(intent); err != nil {
		return nil, err
	}
	return
}

func checkContractInvocationFields(intent *BaseContractInvocationOpIntent) (err error) {


	if intent.Src == nil {
		return newFieldNotFound("src")
	}
	if err = standard.CheckValidHexString(intent.Dst); err != nil {
		return newInvalidFieldError(err).Desc(AtOpIntentField{"dst"})
	}

	if len(intent.Amount) == 0 {
		intent.Amount = "00"
	}
	if err = standard.CheckValidHexString(intent.Amount); err != nil {
		return newInvalidFieldError(err).Desc(AtOpIntentField{"amount"})
	}

	if len(intent.FuncName) == 0 {
		return newFieldNotFound("function name")
	}
	return nil
}
