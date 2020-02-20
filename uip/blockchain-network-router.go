package uip

import "time"

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
