package opintent

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

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
	v3, err := json.Marshal(testdata{Constant: 88888})
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
