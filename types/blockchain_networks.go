package types

type chainID = uint64
type rawTransaction = []byte
type information = []byte
type address = []byte
type KVGetter interface {
	Get([]byte) ([]byte, error)
}

type Router interface {
	RouteRaw(chainID, rawTransaction) (information, error)
	// Route(*TransactionIntent, provedData) (information, error)

	MustWithSigner() bool
	RouteWithSigner(Signer) Router
}

type Translator interface {
	Translate(*TransactionIntent, KVGetter) (rawTransaction, error)
}

type Checker interface {
	CheckAddress(address) bool
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
