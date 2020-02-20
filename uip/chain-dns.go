package uip

type ChainInfo interface {
	GetChainType() ChainType
	GetChainHost() string
}

type ChainDNS interface {
	GetChainInfo(chainID ChainID) (chainInfo ChainInfo, err error)
}
