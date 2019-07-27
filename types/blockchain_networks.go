package types

type chainID = uint64

type Router interface {
	RouteRaw(uint64, []byte) ([]byte, error)
	Route(uint64, []byte) ([]byte, error)
}

type Translator interface {
	Translate(*TransactionIntent, map[string][]byte) ([]byte, error)
}

type Checker interface {
	CheckAddress([]byte) bool
}

type BlockChainInterface interface {
	Router
	Translator
	Checker
}

type CheckerGetter interface {
	GetChecker(chainID) Checker
}

type TranslatorGetter interface {
	GetTranslator(chainID) Translator
}

type RouterGetter interface {
	GetRouter(chainID) Router
}

type BlockChainGetter interface {
	CheckerGetter
	TranslatorGetter
	RouterGetter
	GetBlockChainInterface(chainID) BlockChainInterface
}
