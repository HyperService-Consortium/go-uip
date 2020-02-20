package opintent

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Initializer struct {
	degPool          *DegreePool
	accountBase      uip.AccountBase
	contractBase uip.ContractBase
	largerThanLarger uint32
	chainGetter uip.BlockChainGetter
}

func NewInitializer(options ...interface{}) (*Initializer, error) {
	var ier = &Initializer{}
	for i := range options {
		switch op := options[i].(type) {
		case uip.AccountBase:
			ier.accountBase = op
		case uip.ContractBase:
			ier.contractBase = op
		case uip.BlockChainGetter:
			ier.chainGetter = op
		}
	}
	if ier.degPool == nil {
		ier.degPool = newDegreePool()
	}
	if ier.accountBase == nil {
		return nil, ErrNoAccountBaseProvided
	}
	if ier.contractBase == nil {
		ier.contractBase = uip.ContractBaseNilImpl{}
	}
	if ier.chainGetter == nil {
		return nil, ErrNoBlockChainGetterProvided
	}
	return ier, nil
}