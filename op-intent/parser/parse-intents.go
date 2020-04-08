package parser

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type RawIntentsI interface {
	Len() int
	GetRawIntent(idx int) lexer.Intent
}

func (ier *Parser) ParseIntents(rawIntents RawIntentsI) (intents TxIntentsImpl, err error) {
	intents, err = ier.parseIntents(rawIntents)
	if err != nil {
		return nil, err
	}
	return ier.fillIndex(intents)
}

func (ier *Parser) parseIntents(rawIntents RawIntentsI) (intents TxIntentsImpl, err error) {
	var addition []uip.TxIntentI
	intents, addition = make(TxIntentsImpl, 0, rawIntents.Len()), make([]uip.TxIntentI, 0, 1)

	for idx := 0; idx < rawIntents.Len(); idx++ {
		addition, err = ier.ParseIntent(rawIntents.GetRawIntent(idx))
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentsPos{Pos: idx})
		}
		intents = append(intents, addition...)
	}
	return
}

func (ier *Parser) ParseIntent(rawIntent lexer.Intent) (intents []uip.TxIntentI, err error) {
	switch rawIntent.GetType() {
	case token_type.Pay:
		if intents, err = ier.parsePayment(rawIntent.(*lexer.PaymentIntent)); err != nil {
			return nil, err
		}

	case token_type.Invoke:
		if intents, err = ier.parseContractInvocation(rawIntent.(*lexer.InvokeIntent)); err != nil {
			return nil, err
		}

	case token_type.If:
		if intents, err = ier.parseIf(rawIntent.(*lexer.IfIntent)); err != nil {
			return nil, err
		}

	case token_type.Loop:
		if intents, err = ier.parseLoop(rawIntent.(*lexer.LoopIntent)); err != nil {
			return nil, err
		}

	default:
		return nil, errorn.NewInvalidFieldError(errorn.InvalidOpType)
	}
	return
}
