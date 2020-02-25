package lexer


type ContractInvokeMeta struct {
	Code     []byte      `json:"contract_code"` // key
	FuncName string      `json:"func"`
	Params   []ParamImpl `json:"parameters"`
	Meta     []byte      `json:"meta"`
}
