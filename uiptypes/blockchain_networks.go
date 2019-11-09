package uiptypes

type KVGetter interface {
	Get([]byte) ([]byte, error)
}

type BlockChainInterface interface {
	Router
	Translator
	Checker
	Storage
}

type BlockChainGetter interface {
	CheckerGetter
	TranslatorGetter
	RouterGetter
	StorageGetter
	GetBlockChainInterface(ChainID) BlockChainInterface
}