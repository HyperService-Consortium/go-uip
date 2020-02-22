package opintent

import UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"

type BasePaymentOpIntent struct {
	Src        *RawAccountInfo `json:"src"`    // key
	Dst        *RawAccountInfo `json:"dst"`    // key
	Amount     string          `json:"amount"` // key
	UnitString string          `json:"unit"`   // optional
	Meta       ResultI         `json:"meta"`

	Unit UnitType.Type `json:"-"`
}

func (intent *BasePaymentOpIntent) UnmarshalJSON(b []byte) error {
	r, err := NewGJSONResult(b)
	if err != nil {
		return err
	}
	return intent.UnmarshalResult(r)
}

func (intent *BasePaymentOpIntent) UnmarshalResult(content ResultI) (err error) {
	src := content.Get(FieldOpIntentsSrc)
	if !src.Exists() {
		return newFieldNotFound(FieldOpIntentsSrc)
	}
	dst := content.Get(FieldOpIntentsDst)
	if !dst.Exists() {
		return newFieldNotFound(FieldOpIntentsDst)
	}
	amount := content.Get(FieldOpIntentsAmount)
	if !amount.Exists() {
		return newFieldNotFound(FieldOpIntentsAmount)
	}
	unit := content.Get(FieldOpIntentsUnit)
	if !unit.Exists() {
		return newFieldNotFound(FieldOpIntentsUnit)
	}
	meta := content.Get(FieldOpIntentsMeta)
	intent.Src = new(RawAccountInfo)
	err = intent.Src.UnmarshalResult(src)
	if err != nil {
		return
	}

	intent.Dst = new(RawAccountInfo)
	err = intent.Dst.UnmarshalResult(dst)
	if err != nil {
		return
	}

	intent.Amount = amount.String()
	intent.UnitString = unit.String()
	intent.Meta = meta
	return
}
