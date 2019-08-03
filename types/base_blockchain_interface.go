package types

type BaseBlockChainInterface struct {
	Signer Signer
}

func (bn BaseBlockChainInterface) CheckAddress(address []byte) bool {
	return true
}

func (bn BaseBlockChainInterface) MustWithSigner() bool {
	return false
}

// func (bn *BaseBlockChainInterface) LockSigner() bool {
//
// }

func Route(newBn BlockChainInterface, intent *TransactionIntent, kvGetter KVGetter) ([]byte, error) {
	// todo
	onChainTransaction, err := newBn.Translate(intent, kvGetter)
	if err != nil {
		return nil, err
	}
	return newBn.RouteRaw(intent.ChainID, onChainTransaction)
}
