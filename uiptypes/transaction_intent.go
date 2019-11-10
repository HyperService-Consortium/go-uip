package uiptypes

import "encoding/json"

type hexstring = string

type TransactionIntent struct {
	TransType uint64    `json:"trans_type"`
	Src       []byte    `json:"src"`
	Dst       []byte    `json:"dst"`
	Meta      []byte    `json:"meta"`
	Amt       hexstring `json:"amt"`
	ChainID   uint64    `json:"chain_id"`
}

func (tx *TransactionIntent) Bytes() []byte {
	b, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return b
}

type BaseOpIntent struct {
	Name   string `json:"name"`
	OpType string `json:"op_type"`
}

type BasePaymentOpIntent struct {
	Src        *RawAccountInfo `json:"src"`    // key
	Dst        *RawAccountInfo `json:"dst"`    // key
	Amount     hexstring       `json:"amount"` // key
	UnitString string          `json:"unit"`   // optional
	Meta       []byte          `json:"meta"`
}

type BaseContractInvocationOpIntent struct {
	Src      *RawAccountInfo `json:"invoker"`       // key
	Dst      []byte          `json:"contract_addr"` // key
	Code     []byte          `json:"contract_code"` // key
	FuncName string          `json:"func"`
	Params   []RawParams     `json:"parameters"`
	Meta     []byte          `json:"meta"`
}

type ContractInvokeMeta struct {
	Code     []byte      `json:"contract_code"` // key
	FuncName string      `json:"func"`
	Params   []RawParams `json:"parameters"`
	Meta     []byte      `json:"meta"`
}

type RawAccountInfo struct {
	ChainId uint64 `json:"domain"`
	Name    string `json:"user_name"`
}

type RawParams struct {
	Type  string          `json:"Type"`
	Value json.RawMessage `json:"Value"`
}

type MerkleProofProposal struct {
	Tid uint64
	// 0 = transaction proof, 1 = data proof
	DescriptionType MerkleProofProposalType
	// 0 = simpleMerkleTreeUsingSha256, 1 = simpleMerkleTreeUsingSha512, 2 = merklePatriciaTrieUsingKeccak256
	MerkleProofType   MerkleProofType
	ValueType         TypeID
	SourceDescription []byte
}
