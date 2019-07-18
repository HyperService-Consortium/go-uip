package opintent

import "encoding/json"

type hexstring = string

type RawAccountInfo struct {
	ChainId uint64 `json:"domain"`
	Name    string `json:"user_name"`
}

type RawParams struct {
	Type  string          `json:"Type"`
	Value json.RawMessage `json:"Value"`
}
