package types

import "time"

type chainID = uint64
type rawTransaction = []byte
type address = []byte
type KVGetter interface {
	Get([]byte) ([]byte, error)
}

type WaitOption struct {
	Timeout time.Duration
}

func (w *WaitOption) WithTimeout(t time.Duration) *WaitOption {
	w.Timeout = t
	return w
}

func NewWaitOption() *WaitOption {
	return &WaitOption{
		Timeout: 10 * time.Second,
	}
}

type receipt = []byte
type info = []byte
type additional = []byte
type blockID = []byte
type Router interface {
	RouteRaw(chainID, rawTransaction) (info, error)
	RouteRawTransaction(chainID, rawTransaction) (receipt, error)
	WaitForTransact(chainID, receipt, *WaitOption) (blockID, additional, error)
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

type typeID = uint16
type contract = []byte
type pos = []byte
type desc = []byte
type Storage interface {
	GetTransactionProof(chainID, blockID, additional) (MerkleProof, error)
	GetStorageAt(chainID, typeID, contract, pos, desc) (interface{}, error)
}

type BlockChainInterface interface {
	Router
	Translator
	Checker
	Storage
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

type StorageGetter interface {
	GetBlockStorage(chainID) Router
}

type BlockChainGetter interface {
	CheckerGetter
	TranslatorGetter
	RouterGetter
	StorageGetter
	GetBlockChainInterface(chainID) BlockChainInterface
}
