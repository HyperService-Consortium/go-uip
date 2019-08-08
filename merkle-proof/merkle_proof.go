package merkleproof

import (
	"encoding/json"

	merkleprooftype "github.com/Myriad-Dreamin/go-uip/const/merkle-proof-type"
	types "github.com/Myriad-Dreamin/go-uip/types"
)

type MerkleProof struct {
	Type     uint16 `json:"mtype"`
	RootHash []byte `json:"roothash"`
	Proof    []byte `json:"proof"`
	Key      []byte `json:"key"`
	Value    []byte `json:"value"`
}

func (mp *MerkleProof) GetType() uint16 {
	return mp.Type
}

func (mp *MerkleProof) GetProof() []byte {
	return mp.Proof
}

func (mp *MerkleProof) GetRootHash() []byte {
	return mp.RootHash
}

func (mp *MerkleProof) GetKey() []byte {
	return mp.Key
}

func (mp *MerkleProof) GetValue() []byte {
	return mp.Value
}

func NewMPTUsingKeccak256(proof [][]byte, key, value []byte) types.MerkleProof {
	p, _ := json.Marshal(proof)
	return &MerkleProof{
		Type:     merkleprooftype.MerklePatriciaTrieUsingKeccak256,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}

func NewSecureMPTUsingKeccak256(proof [][]byte, key, value []byte) types.MerkleProof {
	p, _ := json.Marshal(proof)
	return &MerkleProof{
		Type:     merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}

func NewSimpleMekrleTreeUsingSha256(proof [][]byte, key, value []byte) types.MerkleProof {
	p, _ := json.Marshal(proof)
	return &MerkleProof{
		Type:     merkleprooftype.SimpleMerkleTreeUsingSha256,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}

func NewSimpleMekrleTreeUsingSha512(proof [][]byte, key, value []byte) types.MerkleProof {
	p, _ := json.Marshal(proof)
	return &MerkleProof{
		Type:     merkleprooftype.SimpleMerkleTreeUsingSha512,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}
