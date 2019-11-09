package base_blockchain_interface

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type BlockChainInterface struct {
	Signer uiptypes.Signer
}

func (bn BlockChainInterface) CheckAddress(address []byte) bool {
	return true
}

func (bn BlockChainInterface) MustWithSigner() bool {
	return false
}

// func (bn *BlockChainInterface) LockSigner() bool {
//
// }

func Route(newBn uiptypes.BlockChainInterface, intent *uiptypes.TransactionIntent, kvGetter uiptypes.KVGetter) ([]byte, error) {
	// todo
	onChainTransaction, err := newBn.Translate(intent, kvGetter)
	if err != nil {
		return nil, err
	}
	return newBn.RouteRaw(intent.ChainID, onChainTransaction)
}
