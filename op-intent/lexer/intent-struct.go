package lexer

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type ContractInvokeMeta struct {
	Code     []byte              `json:"contract_code"` // key
	FuncName string              `json:"func"`
	Params   []token_types.Param `json:"parameters"`
	Meta     []byte              `json:"meta"`
}

func UnmarshalContractInvokeMeta(b []byte, meta *ContractInvokeMeta) (err error) {
	var r = bytes.NewReader(b)
	serial.Read(r, &meta.Code, &err)
	serial.Read(r, &meta.Meta, &err)
	serial.Read(r, &meta.FuncName, &err)
	var paramsLength uint64
	serial.Read(r, &paramsLength, &err)
	if err != nil {
		return
	}
	meta.Params = make([]uip.VTok, paramsLength)
	for i := range meta.Params {
		lexer_types.DecodeVTok(r, &meta.Params[i], &err)
	}
	return
}

func MarshalContractInvokeMeta(meta *ContractInvokeMeta) (_ []byte, err error) {
	var w = bytes.NewBuffer(nil)
	serial.Write(w, meta.Code, &err)
	serial.Write(w, meta.Meta, &err)
	serial.Write(w, meta.FuncName, &err)
	serial.Write(w, uint64(len(meta.Params)), &err)
	for i := range meta.Params {
		lexer_types.EncodeVTok(w, meta.Params[i], &err)
	}
	return w.Bytes(), err
}
