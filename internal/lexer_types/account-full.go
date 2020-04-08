package lexer_types

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type FullAccount struct {
	Name    string
	ChainID uip.ChainIDUnderlyingType
	Address []byte
}

func (a *FullAccount) GetChainId() uip.ChainID {
	return a.ChainID
}

func (a *FullAccount) GetAddress() uip.Address {
	return a.Address
}

const (
	FieldOpIntentsDomain = "domain"
	FieldKeyAddress      = "address"
)

func (a *FullAccount) UnmarshalResult(nameKey string, content document.Document) (err error) {
	v := content.Get(nameKey)
	if !v.Exists() {
		return errorn.NewFieldNotFound(nameKey)
	}
	a.Name = v.String()
	v = content.Get(FieldOpIntentsDomain)
	if !v.Exists() {
		return errorn.NewFieldNotFound(FieldOpIntentsDomain)
	}
	a.ChainID = v.Uint()
	if a.ChainID == 0 {
		return errorn.NewInvalidFieldError(errors.New("a.ChainID cant not be zero"))
	}
	v = content.Get(FieldKeyAddress)
	if !v.Exists() {
		return errorn.NewFieldNotFound(FieldKeyAddress)
	}
	a.Address, err = DecodeAddress(v.String())
	if err != nil {
		return err
	}
	return nil
}
