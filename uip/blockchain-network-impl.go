package uip

type BlockChainInterfaceImpl struct {
	Signer Signer
}

func (bn BlockChainInterfaceImpl) CheckAddress(address []byte) error {
	return nil
}

func (bn BlockChainInterfaceImpl) MustWithSigner() bool {
	return false
}

// func (bn *BlockChainInterfaceImpl) LockSigner() bool {
//
// }

func Route(newBn BlockChainInterface, intent TransactionIntent, kvGetter Storage) ([]byte, error) {
	onChainTransaction, err := newBn.Translate(intent, kvGetter)
	if err != nil {
		return nil, err
	}
	return newBn.RouteRaw(intent.GetChainID(), onChainTransaction)
}
