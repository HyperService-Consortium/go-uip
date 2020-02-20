package uip


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