package uip

import (
	"errors"
)

type AccountBaseNilImpl struct {
}

func (a AccountBaseNilImpl) AccountBase() AccountBase {
	return a
}

func (AccountBaseNilImpl) Get(name string, chainId uint64) (Account, error) {
	return nil, errors.New("op intent: empty provider")
}

func (AccountBaseNilImpl) GetRelay(domain uint64) (Account, error) {
	return nil, errors.New("op intent: empty provider")
}

func (AccountBaseNilImpl) GetTransactionProofType(chainId uint64) (MerkleProofType, error) {
	return 0, errors.New("op intent: empty provider")
}
