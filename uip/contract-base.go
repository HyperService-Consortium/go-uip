package uip


type ContractBase interface {
	ContractBase() ContractBase

	GetChainID(contractAddress []byte) (cid ChainIDUnderlyingType, err error)
}

