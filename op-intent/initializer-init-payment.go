package opintent

import (
	"encoding/json"
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/standard"
)

type BasePaymentOpIntent struct {
	Src        *RawAccountInfo `json:"src"`    // key
	Dst        *RawAccountInfo `json:"dst"`    // key
	Amount     string          `json:"amount"` // key
	UnitString string          `json:"unit"`   // optional
	Meta       json.RawMessage `json:"meta"`

	Unit UnitType.Type `json:"-"`
}

func (ier *Initializer) initPayment(info *RawIntent, content []byte) (sub interface{}, err error) {
	var intent = new(BasePaymentOpIntent)
	sub = intent
	err = ier.unmarshal(content, intent)
	if err != nil {
		return
	}
	if intent.Src == nil {
		return nil, newFieldNotFound("src")
	}
	if intent.Dst == nil {
		return nil, newFieldNotFound("dst")
	}


	// cinfo, err = SearchChainId(paymentIntent.Src.ChainId)
	// if err != nil {
	// 	return nil, err
	// }
	// var processor ProcessorInterface
	// switch cinfo.GetChainType() {
	// case chain_type.Ethereum:
	// 	processor = eth_processor.GetProcessor()
	// default:
	// 	return nil, errors.New("unsupport chain_type")
	// }
	// if !processor.CheckAddress(paymentIntent.Src.)

	if err = standard.CheckValidHexString(intent.Amount); err != nil {
		return nil, newInvalidFieldError(err).Desc(AtOpIntentField{"amount"})
	}

	t, ok := UnitType.Mapping[intent.UnitString]
	if !ok {
		return nil, newInvalidFieldError(unknownUnitType).Desc(AtOpIntentField{"unit"})
	}

	intent.Unit = t
	return
}
