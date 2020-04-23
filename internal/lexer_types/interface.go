package lexer_types

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"io"
	"strings"
)

type Account interface {
	token_types.Token
	uip.Serializable
	Unmarshal(r io.Reader, i *Account, err *error)
}

type Param interface {
	GetGVMType() gvm.RefType
	Determine(f DetermineContext) (token_types.Param, error)
}

type DeterminedParam interface {
	token_types.Param
	Determine(f DetermineContext) (token_types.Param, error)
}

type DetermineContext interface {
	QueryContract(a Account) (uip.Account, error)
	QueryAccount(a Account) (uip.Account, error)
}

func decodeHex(src string) (b []byte, err error) {
	b, err = hex.DecodeString(src)
	if err != nil {
		return nil, errorn.NewDecodeHexError(err)
	}
	return b, nil
}

func DecodeAddress(src string) ([]byte, error) {
	if strings.HasPrefix(src, "0x") {
		return decodeHex(src[2:])
	}
	return decodeHex(src)
}

func DecodeContractPos(src string) ([]byte, error) {
	return DecodeAddress(src)
}

func DecodeContractAddress(src string) ([]byte, error) {
	return DecodeAddress(src)
}
