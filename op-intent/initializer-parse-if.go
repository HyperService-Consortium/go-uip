package opintent

import "github.com/HyperService-Consortium/go-uip/uip"

func (ier *Initializer) parseIf(rawIntent RawIntentI) (intents []uip.TxIntentI, err error) {
	ifIntent := rawIntent.GetSub().(*BaseIfStatementOpIntent)
	ifIntent.If
	return
}