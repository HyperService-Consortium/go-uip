package opintent

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

type BaseOpIntent struct {
	Name         string `json:"name"`
	OpTypeString string `json:"op_type"`
}

func (ier *Initializer) Parse(
	opIntents uip.OpIntents,
) (txIntents TxIntents, err error) {
	initializedIntents, err := ier.InitContents(opIntents.GetContents())
	if err != nil {
		return nil, err
	}

	initializedDependencies, err := ier.InitDependencies(
		opIntents.GetDependencies())
	if err != nil {
		return nil, err
	}

	parsedDependencies, err := ier.ParseDependencies(
		initializedDependencies, initializedIntents.nameMap)
	if err != nil {
		return nil, err
	}

	intents, err := ier.ParseIntents(initializedIntents)
	if err != nil {
		return nil, err
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(intents, parsedDependencies.dependencies); err != nil {
		return nil, newSortError(err)
	}

	return intents, nil
}
