package bni

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	TransType "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	types "github.com/Myriad-Dreamin/go-uip/types"
)

type testdata struct {
	Constant interface{} `json:"constant"`
}

func TestContractInvocationDataABI(t *testing.T) {
	meta := new(types.ContractInvokeMeta)
	meta.FuncName = "baz"
	meta.Params = make([]types.RawParams, 4, 4)
	v1, err := json.Marshal(testdata{Constant: 69})
	meta.Params[0] = types.RawParams{Type: "uint32", Value: v1}
	v2, err := json.Marshal(testdata{Constant: true})
	meta.Params[1] = types.RawParams{Type: "bool", Value: v2}
	v3, err := json.Marshal(testdata{Constant: "8888888"})
	meta.Params[2] = types.RawParams{Type: "int256", Value: v3}
	v4, err := json.Marshal(testdata{Constant: "SZHNN"})
	meta.Params[3] = types.RawParams{Type: "string", Value: v4}
	res, err := ContractInvocationDataABI(meta, nil)
	if err != nil {
		t.Error("SZHSB", err)
	}
	dst := make([]byte, 1000)
	hex.Encode(dst, res)
	fmt.Println(string(dst))
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
	v1, err := json.Marshal(testdata{Constant: "NNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN"})
	if err != nil {
		t.Error(err)
	}
	meta.Params[1] = types.RawParams{Type: "string", Value: v1}
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
	b, err := new(BN).Translate(tx, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
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
