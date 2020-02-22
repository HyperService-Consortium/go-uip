package uip

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


type MerkleProofProposal struct {
	Tid uint64
	// 0 = transaction proof, 1 = data proof
	DescriptionType MerkleProofProposalType
	// 0 = simpleMerkleTreeUsingSha256, 1 = simpleMerkleTreeUsingSha512, 2 = merklePatriciaTrieUsingKeccak256
	MerkleProofType   MerkleProofType
	ValueType         TypeID
	SourceDescription []byte
}

type TxIntentI interface {
	GetIntent() *TransactionIntent
	SetIntent(*TransactionIntent)

	GetProposals() []MerkleProofProposal
	SetProposals([]MerkleProofProposal)
}

