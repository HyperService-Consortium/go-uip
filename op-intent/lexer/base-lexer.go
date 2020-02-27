package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/standard"
)

type BaseLexer struct {}

func (l *BaseLexer) checkContractInvocation(intent *InvokeIntent) (i *InvokeIntent, err error) {


	if intent.Src == nil {
		return nil, errorn.NewFieldNotFound("src")
	}

	if intent.Dst == nil {
		return nil, errorn.NewFieldNotFound("src")
	}
	//if err = standard.CheckValidHexString(intent.Dst); err != nil {
	//	return nil, errorn.NewInvalidFieldError(err).Desc(errorn.AtOpIntentField{Field: "dst"})
	//}

	if len(intent.Amount) == 0 {
		intent.Amount = "00"
	}
	if err = standard.CheckValidHexString(intent.Amount); err != nil {
		return nil, errorn.NewInvalidFieldError(err).Desc(errorn.AtOpIntentField{Field: "amount"})
	}

	if len(intent.FuncName) == 0 {
		return nil, errorn.NewFieldNotFound("function name")
	}

	return intent, nil
}