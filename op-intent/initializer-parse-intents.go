package opintent

import (
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type RawIntentsI interface {
	Len() int
	GetRawIntent(idx int) RawIntentI
}

type RawIntentI interface {
	GetType() trans_type.Type
	GetSub() interface{}
}

func (ier *Initializer) ParseIntents(rawIntents RawIntentsI) (intents TxIntentsImpl, err error) {
	var addition []uip.TxIntentI
	intents, addition = make(TxIntentsImpl, 0, rawIntents.Len()), make([]uip.TxIntentI, 0, 1)

	for idx := 0; idx < rawIntents.Len(); idx++ {
		addition, err = ier.ParseIntent(rawIntents.GetRawIntent(idx))
		if err != nil {
			return nil, err.(*ParseError).Desc(AtOpIntentsPos{Pos: idx})
		}
		intents = append(intents, addition...)
	}
	return
}

func (ier *Initializer) ParseIntent(rawIntent RawIntentI) (intents []uip.TxIntentI, err error) {
	switch rawIntent.GetType() {
	case trans_type.Payment:
		if intents, err = ier.parsePayment(rawIntent); err != nil {
			return nil, err
		}

	case trans_type.ContractInvoke:
		if intents, err = ier.parseContractInvocation(rawIntent); err != nil {
			return nil, err
		}

	default:
		return nil, newInvalidFieldError(invalidOpType)
	}
	return
}
