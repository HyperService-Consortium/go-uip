package uip

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"time"
)

type TypeIDUnderlyingType = value_type.Type
type TypeID = TypeIDUnderlyingType
type ContractAddress = []byte
type Address = []byte

type BlockChainInterface interface {
	Checker
	Translator
	Router
	Storage
}

type BlockChainGetter interface {
	CheckerGetter
	TranslatorGetter
	RouterGetter
	StorageGetter
	GetBlockChainInterface(ChainID) BlockChainInterface
}

type Checker interface {
	CheckAddress(Address) error
}

type CheckerGetter interface {
	GetChecker(ChainID) (checker Checker, err error)
}

type Translator interface {
	Translate(intent TransactionIntent, storage Storage) (rawTransaction RawTransaction, err error)

	// reflect.DeepEqual(Deserialize(rawTransaction.Serialize()), rawTransaction) == true
	Deserialize(raw []byte) (rawTransaction RawTransaction, err error)
	ParseTransactionIntent(intent TxIntentI) (TxIntentI, error)
}

type TranslatorGetter interface {
	GetTranslator(chainID ChainID) (intent Translator, err error)
}

type RouteOptionTimeout time.Duration

type Router interface {
	RouteRaw(chainID ChainID, rawTransaction RawTransaction) (
		transactionReceipt TransactionReceipt, err error)

	// blockID + color decide an only transaction on chain
	WaitForTransact(chainID ChainID, transactionReceipt TransactionReceipt,
		rOptions ...interface{}) (blockID BlockID, color []byte, err error)

	MustWithSigner() (yesOrNot bool)
	RouteWithSigner(signer Signer) (Router, error)
}

type RouterGetter interface {
	GetRouter(chainID ChainID) (router Router, err error)
}

type Storage interface {

	// blockID + color decide an only transaction on chain with id chainID
	GetTransactionProof(chainID ChainID, blockID BlockID, color []byte) (MerkleProof, error)
	GetStorageAt(chainID ChainID, typeID TypeID, contractAddress ContractAddress, pos []byte, description []byte) (Variable, error)
}

//
//Type     value_type.Type `json:"type"`
//Contract Account         `json:"contract"`
//Pos      []byte          `json:"pos"`
//Field    []byte          `json:"field"`
type StorageGetter interface {
	GetBlockStorage(chainID ChainID) (storage Storage, err error)
}
