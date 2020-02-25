package lexer

import (
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
)

type PaymentIntent struct {
	*IntentImpl
	Src        Account           `json:"src"`    // key
	Dst        Account           `json:"dst"`    // key
	Amount     string            `json:"amount"` // key
	UnitString string            `json:"unit"`   // optional
	Meta       document.Document `json:"meta"`

	Unit UnitType.Type `json:"-"`
}

func (intent *PaymentIntent) UnmarshalJSON(b []byte) error {
	r, err := document.NewGJSONDocument(b)
	if err != nil {
		return err
	}
	return intent.UnmarshalResult(r)
}

func (intent *PaymentIntent) UnmarshalResult(content document.Document) (err error) {
	src := content.Get(FieldOpIntentsSrc)
	if !src.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsSrc)
	}
	dst := content.Get(FieldOpIntentsDst)
	if !dst.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsDst)
	}
	amount := content.Get(FieldOpIntentsAmount)
	if !amount.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsAmount)
	}
	unit := content.Get(FieldOpIntentsUnit)
	if !unit.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsUnit)
	}
	meta := content.Get(FieldOpIntentsMeta)
	intent.Src, err = AccountUnmarshalResult(FieldOpIntentsUserName, src)
	if err != nil {
		return
	}

	intent.Dst, err = AccountUnmarshalResult(FieldOpIntentsUserName, dst)
	if err != nil {
		return
	}

	intent.Amount = amount.String()
	intent.UnitString = unit.String()
	intent.Meta = meta
	return
}
