package bn_storage

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/standard"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/gvm"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/tidwall/gjson"
	"math/big"
	"reflect"
	"testing"
)

type Kv struct {
	K string
	V interface{}
}

type GJSONAssertion struct {
	kvs []Kv
}

func (g GJSONAssertion) AssertBytes(object []byte) (err error) {
	for _, assertKeyValue := range g.kvs {
		k, v := assertKeyValue.K, assertKeyValue.V
		if err = g.compare(gjson.GetBytes(object, k), v); err != nil {
			return fmt.Errorf("compared failed on %v, assertion error %v", k, err)
		}
	}
	return
}

var int64T = reflect.TypeOf(int64(1))

func (g GJSONAssertion) compare(bytes gjson.Result, v interface{}) error {
	t := reflect.TypeOf(v)
	switch bytes.Type {
	case gjson.Null:
		if v != nil {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.False:
		if t.Kind() != reflect.Bool || v != false {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.True:
		if t.Kind() != reflect.Bool || v != true {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.Number:
		if !t.ConvertibleTo(int64T) ||
			reflect.ValueOf(v).Convert(int64T).Int() != bytes.Int() {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.String:
		if t.Kind() != reflect.String || v != bytes.String() {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.JSON:
		return fmt.Errorf("not basic comparable data")
	default:
		panic("unknown g-json type")
	}
	return nil
}

//noinspection GoUnusedExportedFunction
func GJSONWant(kvs ...Kv) GJSONAssertion {
	return GJSONAssertion{kvs: kvs}
}

type MockBNIStorage struct {
	Data []MockData
}

func (m MockBNIStorage) GetTransactionProof(_ uip.ChainID, _ uip.BlockID, _ []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (m MockBNIStorage) GetStorageAt(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	for _, d := range m.Data {
		if d.ChainID == chainID && d.TypeID == typeID &&
			bytes.Equal(d.ContractAddress, contractAddress) &&
			bytes.Equal(d.Pos, pos) &&
			bytes.Equal(d.Description, description) {
			return d.V, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockBNIStorage) insertMockData(data []MockData) {
	m.Data = append(m.Data, data...)
}

type bNIStorageTestSet struct {
	s uip.Storage
}

type MockData struct {
	ChainID         uip.ChainID
	TypeID          uip.TypeID
	ContractAddress uip.ContractAddress
	Pos             []byte
	Description     []byte
	V               uip.Variable
}

type MockValue struct {
	T value_type.Type
	V interface{}
}

func (m MockValue) GetGVMType() gvm.RefType {
	return (gvm.RefType)(m.T)
}

func (m MockValue) Unwrap() interface{} {
	return m.V
}

func (m MockValue) Encode() ([]byte, error) {
	panic("implement me")
}

func (m MockValue) GetType() uip.TypeID {
	return m.T
}

func (m MockValue) GetValue() interface{} {
	return m.V
}

func (m MockValue) IsValid() bool {
	return standard.IsValidStorageValue(m.V, m.T)
}

func (b bNIStorageTestSet) MockingData() []MockData {
	return []MockData{
		{
			ChainID:         2,
			TypeID:          value_type.Int64,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int64, int64(10)},
		},
		{
			ChainID:         2,
			TypeID:          value_type.Int32,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int32, int32(11)},
		},
		{
			ChainID:         2,
			TypeID:          value_type.Int128,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int128, bigInt3},
		},
		{
			ChainID:         2,
			TypeID:          value_type.Int256,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int256, bigInt3},
		},
	}
}

func (b bNIStorageTestSet) RunTests(t *testing.T) {
	t.Run("testGetInt32", b.testGetInt32)
	t.Run("testGetInt64", b.testGetInt64)
	t.Run("testGetInt128", b.testGetInt128)
	t.Run("testGetInt256", b.testGetInt256)
}

func assertType(l *testing.T, x uip.Variable, t value_type.Type, k reflect.Kind) bool {
	l.Helper()
	if x.GetType() != t {
		l.Fatal("bad type")
		return false
	}
	v0 := x.GetValue()
	v := reflect.ValueOf(v0)
	if v.Type().Kind() != k {
		l.Fatal("bad value type")
		return false
	}
	return true
}

func assertTypeOf(l *testing.T, x uip.Variable, t value_type.Type, r reflect.Type) bool {
	l.Helper()
	if x.GetType() != t {
		l.Fatal("bad type")
		return false
	}
	v0 := x.GetValue()
	v := reflect.ValueOf(v0)
	if v.Type() != r {
		l.Fatal("bad value type")
		return false
	}
	return true
}

func (b bNIStorageTestSet) testGetInt32(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int32, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertType(t, x, value_type.Int32, reflect.Int32) {
		return
	}
	if x.GetValue().(int32) != 11 {
		t.Fatal("bad value")
	}
}

func (b bNIStorageTestSet) testGetInt64(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int64, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertType(t, x, value_type.Int64, reflect.Int64) {
		return
	}
	if x.GetValue().(int64) != 10 {
		t.Fatal("bad value")
	}
}

var bigInt3 = big.NewInt(3)

func (b bNIStorageTestSet) testGetInt128(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int128, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertTypeOf(t, x, value_type.Int128, reflect.TypeOf(bigInt3)) {
		return
	}
	if x.GetValue().(*big.Int).Cmp(bigInt3) != 0 {
		t.Fatal("bad value")
	}
}

func (b bNIStorageTestSet) testGetInt256(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int256, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertTypeOf(t, x, value_type.Int256, reflect.TypeOf(bigInt3)) {
		return
	}
	if x.GetValue().(*big.Int).Cmp(bigInt3) != 0 {
		t.Fatal("bad value")
	}
}
