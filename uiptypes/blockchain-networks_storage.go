package uiptypes


type TypeIDUnderlyingType = uint16
type TypeID TypeIDUnderlyingType
type ContractAddress = []byte
type Pos = []byte
type Desc = []byte

type Storage interface {

	// blockID + color decide an only transaction on chain with id chainID
	GetTransactionProof(chainID ChainID, blockID BlockID, color []byte) (MerkleProof, error)
	GetStorageAt(chainID ChainID, typeID TypeID, contractAddress ContractAddress, pos Pos, description Desc) (interface{}, error)
}

type StorageGetter interface {
	GetBlockStorage(ChainID) Router
}
