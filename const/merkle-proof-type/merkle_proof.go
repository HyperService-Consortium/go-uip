package merkle_proof

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type Type = uiptypes.MerkleProofType

const (
	Invalid Type = 0 + iota
	SimpleMerkleTreeUsingSha256
	SimpleMerkleTreeUsingSha512
	MerklePatriciaTrieUsingKeccak256
	SecureMerklePatriciaTrieUsingKeccak256
)
