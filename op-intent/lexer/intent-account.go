package lexer

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/uip"
	"strconv"
	"strings"
)

func AccountUnmarshalResult(entityKey string, content document.Document) (_ lexer_types.Account, err error) {
	if content.IsArray() {
		return nil, errorn.NewInvalidDocumentType("want string or object")
	}
	if !content.IsObject() {
		x := content.String()
		if strings.HasPrefix(x, "0x") {
			b, err := lexer_types.DecodeAddress(x[2:])
			if err != nil {
				return nil, errorn.NewDecodeAddressError(err)
			}
			return &lexer_types.RawAccount{Address: b}, nil
		}
		return &lexer_types.NameAccount{Name: x}, nil
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
			return &lexer_types.NamespacedNameAccount{Name: v.String(), ChainID: domain}, nil
		} else {
			return &lexer_types.NameAccount{Name: v.String()}, nil
		}
	}

	v = content.Get(FieldKeyAddress)
	if v.Exists() {
		x := v.String()
		if strings.HasPrefix(x, "0x") {
			x = x[2:]
		}
		b, err := lexer_types.DecodeAddress(x)
		if err != nil {
			return nil, errorn.NewDecodeAddressError(err)
		}
		if domain != 0 {
			return &lexer_types.NamespacedRawAccount{Address: b, ChainID: domain}, nil
		} else {
			return &lexer_types.RawAccount{Address: b}, nil
		}
	}

	return nil, errorn.NewInvalidFieldError(errorn.UnknownAccount)
}
