package opintent

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

type testdata struct {
	Constant interface{} `json:"constant"`
}

func TestContractInvocationDataABI(t *testing.T) {
	meta := new(contractInvokeMeta)
	meta.FuncName = "baz"
	meta.Params = make([]RawParams, 4, 4)
	v1, err := json.Marshal(testdata{Constant: 69})
	meta.Params[0] = RawParams{Type: "uint32", Value: v1}
	v2, err := json.Marshal(testdata{Constant: true})
	meta.Params[1] = RawParams{Type: "bool", Value: v2}
	v3, err := json.Marshal(testdata{Constant: 88888})
	meta.Params[2] = RawParams{Type: "int256", Value: v3}
	v4, err := json.Marshal(testdata{Constant: "SZHNN"})
	meta.Params[3] = RawParams{Type: "string", Value: v4}
	res, err := ContractInvocationDataABI(meta, nil)
	if err != nil {
		t.Error("SZHSB", err)
	}
	dst := make([]byte, 1000)
	hex.Encode(dst, res)
	fmt.Println(string(dst))
}
