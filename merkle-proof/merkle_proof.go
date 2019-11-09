package merkleproof

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/uiptypes"

	merkleprooftype "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
)

type MerkleProof struct {
	Type     uiptypes.MerkleProofType `json:"mtype"`
	RootHash []byte `json:"roothash"`
	Proof    []byte `json:"proof"`
	Key      []byte `json:"key"`
	Value    []byte `json:"value"`
}

func (mp *MerkleProof) GetType() uiptypes.MerkleProofType {
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

type MPTMerkleProof struct {
	RootHash  []byte   `json:"r"`
	HashChain [][]byte `json:"h"`
}

func NewMPTUsingKeccak256(proof [][]byte, key, value []byte) uiptypes.MerkleProof {
	p, _ := json.Marshal(&MPTMerkleProof{
		RootHash:  proof[0],
		HashChain: proof,
	})
	return &MerkleProof{
		Type:     merkleprooftype.MerklePatriciaTrieUsingKeccak256,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}

func NewSecureMPTUsingKeccak256(proof [][]byte, key, value []byte) uiptypes.MerkleProof {
	p, _ := json.Marshal(&MPTMerkleProof{
		RootHash:  proof[0],
		HashChain: proof,
	})
	return &MerkleProof{
		Type:     merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}

type SimpleMerkleProof struct {
	HashChain [][]byte `json:"h"`
}

func NewSimpleMekrleTreeUsingSha256(proof [][]byte, key, value []byte) uiptypes.MerkleProof {
	p, _ := json.Marshal(&SimpleMerkleProof{
		HashChain: proof,
	})
	return &MerkleProof{
		Type:     merkleprooftype.SimpleMerkleTreeUsingSha256,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}

func NewSimpleMekrleTreeUsingSha512(proof [][]byte, key, value []byte) uiptypes.MerkleProof {
	p, _ := json.Marshal(&SimpleMerkleProof{
		HashChain: proof,
	})
	return &MerkleProof{
		Type:     merkleprooftype.SimpleMerkleTreeUsingSha512,
		RootHash: proof[0],
		Proof:    p,
		Key:      key,
		Value:    value,
	}
}
