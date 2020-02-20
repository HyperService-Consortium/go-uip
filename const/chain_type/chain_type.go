package ChainType

import "github.com/HyperService-Consortium/go-uip/uip"

type Type = uip.ChainType

const (
	Unassigned Type = iota
	Ethereum
	TendermintNSB
)
