package bni

import (
	"errors"
	"strings"

	chainType "github.com/Myriad-Dreamin/go-uip/const/chain_type"
)

type ChainInfo struct {
	Host      string
	ChainType uint64
}

func decoratePrefix(hexs string) string {
	if strings.HasPrefix(hexs, "0x") {
		return hexs
	} else {
		return "0x" + hexs
	}
}

func (c *ChainInfo) GetHost() string {
	return c.Host
}

func (c *ChainInfo) GetChainType() uint64 {
	return c.ChainType
}

type ChainInfoInterface interface {
	GetHost() string
	GetChainType() uint64
}

func SearchChainId(domain uint64) (ChainInfoInterface, error) {
	switch domain {
	case 0:
		return nil, errors.New("nil domain is not allowed")
	case 1: // ethereum chain 1
		return &ChainInfo{
			Host:      "127.0.0.1:8545",
			ChainType: chainType.Ethereum,
		}, nil
	case 2: // ethereum chain 2
		return &ChainInfo{
			Host:      "127.0.0.1:8545",
			ChainType: chainType.Ethereum,
		}, nil
	case 3: // tendermint chain 1
		return &ChainInfo{
			Host:      "47.254.66.11:26657",
			ChainType: chainType.TendermintNSB,
		}, nil
	case 4: // ethereum chain 1
		return &ChainInfo{
			Host:      "47.251.2.73:26657",
			ChainType: chainType.TendermintNSB,
		}, nil
	default:
		return nil, errors.New("not found")
	}
}
