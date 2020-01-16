package uiptypes

type ChainInfo interface {
	GetChainType() ChainType
	GetChainHost() string
}

type ChainDNS interface {
	GetChainInfo(ChainID) (ChainInfo, error)
}
