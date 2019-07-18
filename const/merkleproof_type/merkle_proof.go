package merkle_proof

type Type = uint16

const (
	Invalid Type = 0 + iota
	SimpleMerkleTreeUsingSha256
	SimpleMerkleTreeUsingSha512
	MerklePatriciaTrieUsingKeccak256
	SecureMerklePatriciaTrieUsingKeccak256
)
