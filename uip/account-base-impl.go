package uip

import "errors"

type AccountBaseImpl struct {

}

func (a AccountBaseImpl) AccountBase() AccountBase {
	return a
}

func (AccountBaseImpl) Get(name string, chainId uint64) (Account, error) {
	return nil, errors.New("op intent: empty provider")
}

func (AccountBaseImpl) GetRelay(domain uint64) (Account, error) {
	return nil, errors.New("op intent: empty provider")
}

func (AccountBaseImpl) GetTransactionProofType(chainId uint64) (MerkleProofType, error) {
	return 0, errors.New("op intent: empty provider")
}