package opintent

import (
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Account = lexer_types.Account
type LexParam = lexer_types.Param
type DeterminedParam = lexer_types.DeterminedParam
type DetermineContext = lexer_types.DetermineContext

type RawAccount = lexer_types.RawAccount
type NameAccount = lexer_types.NameAccount
type NamespacedRawAccount = lexer_types.NamespacedRawAccount
type NamespacedNameAccount = lexer_types.NamespacedNameAccount
type FullAccount = lexer_types.FullAccount

//noinspection GoUnusedExportedFunction
func NewNamespacedRawAccount(a uip.Account) *NamespacedRawAccount {
	return NewNamespacedRawAccount(a)
}

type BinaryExpression = lexer_types.BinaryExpression
type LocalStateVariable = lexer_types.LocalStateVariable
type StateVariable = lexer_types.StateVariable
type UnaryExpression = lexer_types.UnaryExpression

type Uint8 = lexer_types.Uint8
type Uint16 = lexer_types.Uint16
type Uint32 = lexer_types.Uint32
type Uint64 = lexer_types.Uint64
type Uint128 = lexer_types.Uint128
type Uint256 = lexer_types.Uint256
type Int8 = lexer_types.Int8
type Int16 = lexer_types.Int16
type Int32 = lexer_types.Int32
type Int64 = lexer_types.Int64
type Int128 = lexer_types.Int128
type Int256 = lexer_types.Int256

type String = lexer_types.String
type Bytes = lexer_types.Bytes
type Bool = lexer_types.Bool
type Unknown = lexer_types.Unknown

var Undefined = lexer_types.Undefined

type ChainMap = lexer_types.ChainMap
type AccountMap = lexer_types.AccountMap

func BuildAccountMap(accounts []FullAccount) (res AccountMap, err error) {
	return lexer_types.BuildAccountMap(accounts)
}

func DecodeAddress(src string) ([]byte, error) {
	return lexer_types.DecodeAddress(src)
}

func DecodeContractPos(src string) ([]byte, error) {
	return lexer_types.DecodeContractPos(src)
}
