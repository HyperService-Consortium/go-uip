package uiptypes


type TypeIDUnderlyingType = uint16
type TypeID TypeIDUnderlyingType
type Contract = []byte
type Pos = []byte
type Desc = []byte

type Storage interface {

	// blockID + color decide an only transaction on chain with id chainID
	GetTransactionProof(chainID ChainID, blockID BlockID, color []byte) (MerkleProof, error)
	GetStorageAt(chainID ChainID, typeID TypeID, contract Contract, post Pos, description Desc) (interface{}, error)
}

type StorageGetter interface {
	GetBlockStorage(ChainID) Router
}
