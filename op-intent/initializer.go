package opintent

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type Initializer struct {
	degPool          *DegreePool
	accountBase      uip.AccountBase
	contractBase     uip.ContractBase
	chainGetter      uip.BlockChainGetter
	largerThanLarger uint32
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
		return nil, errorn.ErrNoAccountBaseProvided
	}
	if ier.contractBase == nil {
		ier.contractBase = uip.ContractBaseNilImpl{}
	}
	if ier.chainGetter == nil {
		return nil, errorn.ErrNoBlockChainGetterProvided
	}
	return ier, nil
}
