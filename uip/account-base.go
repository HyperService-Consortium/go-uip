package uip

type AccountBase interface {
	// sign
	AccountBase() AccountBase

	Get(name string, chainId ChainIDUnderlyingType) (Account, error)
	GetRelay(domain ChainIDUnderlyingType) (Account, error)
	GetTransactionProofType(chainId ChainIDUnderlyingType) (MerkleProofType, error)
}
