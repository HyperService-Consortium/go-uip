package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer/internal"
	"github.com/HyperService-Consortium/go-uip/uip"
	"strconv"
	"strings"
)

type Account = internal.Account

type RawAccount = internal.RawAccount
type NameAccount = internal.NameAccount
type NamespacedAccount = internal.NamespacedNameAccount
type NamespacedRawAccount = internal.NamespacedRawAccount

func NewNamespacedRawAccount(a uip.Account) Account {
	return internal.NewNamespacedRawAccount(a)
}

func AccountUnmarshalResult(entityKey string, content document.Document) (_ Account, err error) {
	if content.IsArray() {
		return nil, errorn.NewInvalidDocumentType("want string or object")
	}
	if !content.IsObject() {
		x := content.String()
		if strings.HasPrefix(x, "0x") {
			b, err := DecodeAddress(x[2:])
			if err != nil {
				return nil, errorn.NewDecodeAddressError(err)
			}
			return &RawAccount{Address: b}, nil
		}
		return &NameAccount{Name: x}, nil
	}
	var domain uip.ChainIDUnderlyingType
	v := content.Get(FieldOpIntentsDomain)
	if v.Exists() {
		domain, err = strconv.ParseUint(v.String(), 10, 64)
		if err != nil {
			return nil, errorn.NewDecodeDomainError(err)
		}
	}
	v = content.Get(entityKey)
	if v.Exists() {
		if domain != 0 {
			return &NamespacedAccount{Name: v.String(), ChainID: domain}, nil
		} else {
			return &NameAccount{Name: v.String()}, nil
		}
	}

	v = content.Get(FieldKeyAddress)
	if v.Exists() {
		x := v.String()
		if strings.HasPrefix(x, "0x") {
			x = x[2:]
		}
		b, err := DecodeAddress(x)
		if err != nil {
			return nil, errorn.NewDecodeAddressError(err)
		}
		if domain != 0 {
			return &NamespacedRawAccount{Address: b, ChainID: domain}, nil
		} else {
			return &RawAccount{Address: b}, nil
		}
	}

	return nil, errorn.NewInvalidFieldError(errorn.UnknownAccount)
}
