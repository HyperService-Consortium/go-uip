package uip

import (
	"errors"
)

type ContractBaseNilImpl struct {
}

func (a ContractBaseNilImpl) ContractBase() ContractBase {
	return a
}

func (ContractBaseNilImpl) GetChainID(contractAddress []byte) (ChainIDUnderlyingType, error) {
	return 0, errors.New("contract base is not provided")
}
