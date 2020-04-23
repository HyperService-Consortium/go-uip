package lexer_types

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"io"
	"math/big"
	"testing"
)

type IdentQuery struct{}

func (i IdentQuery) QueryContract(a Account) (uip.Account, error) {
	return a.(uip.Account), nil
}

func (i IdentQuery) QueryAccount(a Account) (uip.Account, error) {
	return a.(uip.Account), nil
}

func TestSerialize(t *testing.T) {
	var x IdentQuery

	type testCase struct {
		name      string
		marshal   func(r io.Writer, b interface{})
		unmarshal func(r io.Reader) interface{}
		encoding  uip.Serializable
	}

	var inheritAccountTestCase = func(name string, encoding uip.Serializable) testCase {
		return testCase{
			name:      name,
			encoding:  encoding,
			marshal:   encodeAccount,
			unmarshal: unmarshalAccount,
		}
	}

	var inheritTokenTestCase = func(name string, encoding uip.Serializable) testCase {
		return testCase{
			name:      name,
			encoding:  encoding,
			marshal:   encodeToken,
			unmarshal: unmarshalToken,
		}
	}

	for _, tt := range []testCase{
		inheritAccountTestCase("name-account", NameAccount{Name: "a"}),
		inheritAccountTestCase("raw-account", RawAccount{Address: []byte("a")}),
		inheritAccountTestCase("namespaced-raw-account", NamespacedRawAccount{Address: []byte("a"), ChainID: 233}),
		inheritAccountTestCase("namespaced-name-account", NamespacedNameAccount{Name: "a", ChainID: 233}),
		inheritTokenTestCase("binary-expression", sugar.HandlerError((BinaryExpression{
			Type:  1,
			Sign:  1,
			Left:  Int64(3333),
			Right: Int64(233),
		}).Determine(nil)).(uip.Serializable)),
		inheritTokenTestCase("unary-expression", sugar.HandlerError((UnaryExpression{
			Type: 1,
			Sign: 1,
			Left: Int64(4444),
		}).Determine(nil)).(uip.Serializable)),
		inheritTokenTestCase("local-state-variable", sugar.HandlerError((LocalStateVariable{
			Type:  1,
			Pos:   []byte("pos"),
			Field: []byte("field"),
		}).Determine(nil)).(uip.Serializable)),
		inheritTokenTestCase("state-variable", sugar.HandlerError((&StateVariable{
			Type:     1,
			Pos:      []byte("pos"),
			Field:    []byte("field"),
			Contract: NamespacedRawAccount{Address: []byte("a"), ChainID: 233},
		}).Determine(x)).(uip.Serializable)),
		inheritTokenTestCase("uint8", Uint8(1)),
		inheritTokenTestCase("uint16", Uint16(1)),
		inheritTokenTestCase("uint32", Uint32(1)),
		inheritTokenTestCase("uint64", Uint64(1)),
		inheritTokenTestCase("uint128", (*Uint128)(big.NewInt(1))),
		inheritTokenTestCase("uint256", (*Uint256)(big.NewInt(1))),
		inheritTokenTestCase("int8", Int8(1)),
		inheritTokenTestCase("int16", Int16(1)),
		inheritTokenTestCase("int32", Int32(1)),
		inheritTokenTestCase("int64", Int64(1)),
		inheritTokenTestCase("int128", (*Int128)(big.NewInt(1))),
		inheritTokenTestCase("int256", (*Int256)(big.NewInt(1))),
		inheritTokenTestCase("string", String("a")),
		inheritTokenTestCase("bool", Bool(true)),
		inheritTokenTestCase("bytes", Bytes("a")),
	} {
		t.Run(tt.name, func(t *testing.T) {
			var b = bytes.NewBuffer(nil)
			tt.marshal(b, tt.encoding)
			assert.EqualValues(t, tt.encoding, tt.unmarshal(b))
			assert.EqualValues(t, 0, b.Len())
		})
	}

}

func encodeAccount(r io.Writer, b interface{}) {
	var err error
	EncodeAccount(r, b.(Account), &err)
	sugar.HandlerError0(err)
	return
}

func unmarshalAccount(r io.Reader) interface{} {
	var x Account
	var err error
	DecodeAccount(r, &x, &err)
	sugar.HandlerError0(err)
	return x
}

func encodeToken(r io.Writer, b interface{}) {
	var err error
	EncodeVTok(r, b.(uip.VTok), &err)
	sugar.HandlerError0(err)
	return
}

func unmarshalToken(r io.Reader) interface{} {
	var x uip.VTok
	var err error
	DecodeVTok(r, &x, &err)
	sugar.HandlerError0(err)
	return x
}
