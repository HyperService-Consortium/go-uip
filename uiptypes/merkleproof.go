package uiptypes

type MerkleProofType uint16
type RootHash = []byte
type Proof = []byte
type MerkleProofKey = []byte
type MerkleProofValue = []byte

// MerkleProof is the merkle proof that proves the existence of key-value pair on the  merkle tree
type MerkleProof interface {
	GetType() MerkleProofType
	GetRootHash() RootHash
	GetProof() Proof
	GetKey() MerkleProofKey
	GetValue() MerkleProofValue
}

// BlockCheck is the merkle proof that proves the existence of merkle tree (root hash) on the block
type BlockCheck interface {
	GetChainId() ChainID
	GetBlockId() BlockID
	GetRootTag() uint8
	GetRootHash() RootHash
}
