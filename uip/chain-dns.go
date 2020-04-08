package uip

type ChainIDUnderlyingType = uint64
type ChainID = ChainIDUnderlyingType
type ChainTypeUnderlyingType = uint64
type ChainType ChainTypeUnderlyingType

type ChainInfo interface {
	GetChainType() ChainType
	GetChainHost() string
}

type ChainDNS interface {
	GetChainInfo(chainID ChainID) (chainInfo ChainInfo, err error)
}
