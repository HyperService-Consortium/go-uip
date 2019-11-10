package ChainType

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type Type = uiptypes.ChainType

const (
	Unassigned Type = iota
	Ethereum
	TendermintNSB
)
