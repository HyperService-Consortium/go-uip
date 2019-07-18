package opintent

import "encoding/json"

type TransactionIntent struct {
	TransType uint64    `json:"trans_type"`
	Src       []byte    `json:"src"`
	Dst       []byte    `json:"dst"`
	Meta      []byte    `json:"meta"`
	Amt       hexstring `json:"amt"`
	ChainId   uint64    `json:"chain_id"`
}

func (tx *TransactionIntent) Bytes() []byte {
	b, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return b
}

type MerkleProofProposal struct {
	// 0 = transaction proof, 1 = data proof
	DescriptionType uint8
	// 0 = simpleMerkleTreeUsingSha256, 1 = simpleMerkleTreeUsingSha512, 2 = merklePatriciaTrieUsingKeccak256
	MerkleProofType   uint16
	ValueType         uint16
	SourceDescription []byte
}

// type TransactionIntentInterface interface {
// 	GetSrc() []byte
// 	GetDst() []byte
// 	GetChainId() uint64
// 	GetSrcAccount() types.Account
// 	GetDstAccount() types.Account
// 	GetMeta() []byte
// 	Init() error
// 	Send() error
// 	Call() error
// }

// type EthTransactionIntent struct {
// 	TransactionIntent
// }
