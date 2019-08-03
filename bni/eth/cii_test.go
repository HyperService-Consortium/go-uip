package bni

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	TransType "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	valuetype "github.com/Myriad-Dreamin/go-uip/const/value_type"
	types "github.com/Myriad-Dreamin/go-uip/types"
)

type testdata struct {
	Constant interface{} `json:"constant"`
}

func TestContractInvocationDataABI(t *testing.T) {
	meta := new(types.ContractInvokeMeta)
	meta.FuncName = "sam"
	meta.Params = make([]types.RawParams, 3, 3)
	v1, err := json.Marshal(testdata{Constant: []uint32{2, 3, 4}})
	meta.Params[0] = types.RawParams{Type: "uint32[]", Value: v1}
	v2, err := json.Marshal(testdata{Constant: []byte{1, 2, 3}})
	meta.Params[1] = types.RawParams{Type: "bytes", Value: v2}
	v3, err := json.Marshal(testdata{Constant: "0x7f49b5c4c1cae9ea898f856ea4c2e10f3d5a3456"})
	meta.Params[2] = types.RawParams{Type: "address", Value: v3}
	res, err := ContractInvocationDataABI(meta, nil)
	if err != nil {
		t.Error("SZHSB", err)
	}
	dst := make([]byte, 1000)
	hex.Encode(dst, res)
	fmt.Println(string(dst))
}

type getter struct {
}

func (g *getter) Get([]byte) ([]byte, error) {
	return []byte("true"), nil
}

func TestDataTransaction(t *testing.T) {
	meta := new(types.ContractInvokeMeta)
	meta.FuncName = "baz"
	meta.Params = make([]types.RawParams, 2, 2)
	v0, err := json.Marshal(testdata{Constant: 2})
	if err != nil {
		t.Error(err)
	}
	meta.Params[0] = types.RawParams{Type: "uint32", Value: v0}
	// v1, err := json.Marshal(testdata{Constant: "NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN"})
	// if err != nil {
	// 	t.Error(err)
	// }
	meta.Params[1] = types.RawParams{Type: "bool", Value: []byte(`{"contract":"1234567812345678123456781234567812345678", "pos":"0x0", "field":"aut"}`)}
	tx := new(types.TransactionIntent)
	tx.Src = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	tx.Dst = []byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	tx.Meta, err = json.Marshal(meta)
	if err != nil {
		t.Error(err)
	}
	tx.Amt = "0"
	tx.TransType = TransType.ContractInvoke
	tx.ChainID = 1
	b, err := new(BN).Translate(tx, new(getter))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}

func TestGetStorageAt(t *testing.T) {
	b, _ := hex.DecodeString("1234567812345678123456781234567812345678")
	fmt.Println(new(BN).GetStorageAt(1, valuetype.Bool, b, []byte{1}, []byte("some varible")))
	fmt.Println(new(BN).GetStorageAt(1, valuetype.Uint256, b, []byte{1}, []byte("some varible")))
}

/*
9d2206ab
0000000000000000000000000000000000000000000000000000000000000002
0000000000000000000000000000000000000000000000000000000000000040
000000000000000000000000000000000000000000000000000000000000006c
4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e
4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e
4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e4e
4e4e4e4e4e4e4e4e4e4e4e4e0000000000000000000000000000000000000000
*/
