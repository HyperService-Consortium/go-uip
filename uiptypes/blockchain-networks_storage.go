package uiptypes


type TypeIDUnderlyingType = uint16
type TypeID TypeIDUnderlyingType
type ContractAddress = []byte

type Variable interface {
	GetType() TypeID
	GetValue() interface{}
}

type Storage interface {

	// blockID + color decide an only transaction on chain with id chainID
	GetTransactionProof(chainID ChainID, blockID BlockID, color []byte) (MerkleProof, error)
	GetStorageAt(chainID ChainID, typeID TypeID, contractAddress ContractAddress, pos []byte, description []byte) (Variable, error)
}

type StorageGetter interface {
	GetBlockStorage(chainID ChainID) (router Router)
}
