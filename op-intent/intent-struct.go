package opintent


type BaseOpIntent struct {
	Name         string `json:"name"`
	OpTypeString string `json:"op_type"`
}

type ContractInvokeMeta struct {
	Code     []byte     `json:"contract_code"` // key
	FuncName string     `json:"func"`
	Params   []RawParam `json:"parameters"`
	Meta     []byte     `json:"meta"`
}
