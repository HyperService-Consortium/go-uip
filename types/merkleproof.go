package types

type merkleProofType = uint16
type proof = []byte
type key = []byte
type value = []byte

// MerkleProof is the merkle proof that proves the existence of key-value pair on the  merkle tree
type MerkleProof interface {
	GetType() uint16
	GetRootHash() []byte
	GetProof() []byte
	GetKey() []byte
	GetValue() []byte
}

// BlockCheck is the merkle proof that proves the existence of merkle tree (root hash) on the block
type BlockCheck interface {
	GetChainId() uint64
	GetBlockId() []byte
	GetRootTag() uint8
	GetRootHash() []byte
}
