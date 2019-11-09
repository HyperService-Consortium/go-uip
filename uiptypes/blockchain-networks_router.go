package uiptypes

import "time"

type RouteOptionTimeout time.Duration

type Router interface {
	RouteRaw(ChainID, RawTransaction) (
		transactionReceipt TransactionReceipt, err error)

	// blockID + color decide an only transaction on chain
	WaitForTransact(chainID ChainID, transactionReceipt TransactionReceipt,
		options ...interface{}) (blockID BlockID, color []byte, err error)

	MustWithSigner() (yesOrNot bool)
	RouteWithSigner(signer Signer) (Router, error)
}

type RouterGetter interface {
	GetRouter(ChainID) Router
}
