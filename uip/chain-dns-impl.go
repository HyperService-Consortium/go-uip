package uip

import "errors"

type ChainDNSImpl map[ChainIDUnderlyingType]ChainInfo

func (c ChainDNSImpl) GetChainInfo(cid ChainID) (ci ChainInfo, err error) {
	ci, ok := c[cid]
	if !ok {
		return nil, errors.New("not found")
	}
	return ci, nil
}

func (c ChainDNSImpl) SetChainInfo(cid ChainID, ci ChainInfo) {
	c[cid] = ci
}


