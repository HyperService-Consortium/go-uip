package bni

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"

	ethclient "github.com/Myriad-Dreamin/go-ves/net/eth_client"

	TransType "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	opintent "github.com/Myriad-Dreamin/go-uip/op-intent"
	types "github.com/Myriad-Dreamin/go-uip/types"
)

type BN struct {
}

func (bn *BN) RouteRaw(destination uint64, payload []byte) ([]byte, error) {
	ci, err := SearchChainId(destination)
	if err != nil {
		return nil, err
	}
	return ethclient.Do((&url.URL{Scheme: "http", Host: ci.GetHost(), Path: "/"}).String(), payload)
}

func (bn *BN) Route(intent *types.TransactionIntent, kvs map[string][]byte) ([]byte, error) {
	// todo
	onChainTransaction, err := bn.Translate(intent, kvs)
	if err != nil {
		return nil, err
	}
	return bn.RouteRaw(intent.ChainID, onChainTransaction)
}

func (bn *BN) Translate(
	intent *opintent.TransactionIntent,
	kvs map[string][]byte,
) ([]byte, error) {
	switch intent.TransType {
	case TransType.Payment:
		return json.Marshal(map[string]interface{}{
			"from":  decoratePrefix(hex.EncodeToString(intent.Src)),
			"to":    decoratePrefix(hex.EncodeToString(intent.Dst)),
			"value": decoratePrefix(intent.Amt),
		})
	case TransType.ContractInvoke:
		var meta types.ContractInvokeMeta
		err := json.Unmarshal(intent.Meta, &meta)
		if err != nil {
			return nil, err
		}
		_ = meta

		return nil, errors.New("todo")
	default:
		return nil, errors.New("cant translate")
	}
}

func (bn *BN) CheckAddress(addr []byte) bool {
	return len(addr) == 32
}
