package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
	"strconv"
	"strings"
)

type Account = token.Token

type NamespacedAccount struct {
	Name string `json:"name"`
	ChainID uip.ChainIDUnderlyingType `json:"chain_id"`
}

func (n NamespacedAccount) GetName() string {
	return n.Name
}

func (n NamespacedAccount) GetChainID() uip.ChainIDUnderlyingType {
	return n.ChainID
}

func (n NamespacedAccount) GetType() token.Type {
	return token.NamespacedNameAccount
}

type NameAccount struct {
	Name string `json:"name"`
}

func (n NameAccount) GetName() string {
	return n.Name
}

func (n NameAccount) GetType() token.Type {
	return token.NameAccount
}

type NamespacedRawAccount struct {
	Address []byte `json:"address"`
	ChainID uip.ChainIDUnderlyingType `json:"chain_id"`
}

func (n NamespacedRawAccount) GetChainId() uip.ChainID {
	return n.ChainID
}

func (n NamespacedRawAccount) GetAddress() uip.Address {
	return n.Address
}

func NewNamespacedRawAccount(a uip.Account) Account {
	return &NamespacedRawAccount{
		Address: a.GetAddress(),
		ChainID: a.GetChainId(),
	}
}

func (n NamespacedRawAccount) GetType() token.Type {
	return token.NamespacedRawAccount
}

type RawAccount struct {
	Address []byte `json:"address"`
}

func (n RawAccount) GetAddress() []byte {
	return n.Address
}

func (n RawAccount) GetType() token.Type {
	return token.RawAccount
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
			return &RawAccount{Address:b}, nil
		}
		return &NameAccount{Name:x}, nil
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

