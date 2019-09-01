package chaininfo

import (
	"encoding/hex"
	"errors"

	chain_type "github.com/HyperService-Consortium/go-uip/const/chain_type"
	merkleprooftype "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	types "github.com/HyperService-Consortium/go-uip/types"
	"github.com/HyperService-Consortium/go-uip/types/account"
)

type ChainInfo struct {
	Host      string
	ChainType uint64
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
			ChainType: chain_type.Ethereum,
		}, nil
	case 2: // ethereum chain 2
		return &ChainInfo{
			Host:      "127.0.0.1:8545",
			ChainType: chain_type.Ethereum,
		}, nil
	case 3: // tendermint chain 1
		return &ChainInfo{
			Host:      "47.251.2.73:26657", //"47.254.66.11",
			ChainType: chain_type.TendermintNSB,
		}, nil
	case 4: // tendermint chain 2
		return &ChainInfo{
			Host:      "47.251.2.73:26657",
			ChainType: chain_type.TendermintNSB,
		}, nil
	case 5: // tendermint chain 3
		return &ChainInfo{
			Host:      "47.88.101.45:26657",
			ChainType: chain_type.TendermintNSB,
		}, nil
	default:
		return nil, errors.New("not found")
	}
}

func TempGetRelay(domain uint64) (types.Account, error) {
	switch domain {
	case 0:
		return nil, errors.New("nil domain is not allowed")
	case 1: // ethereum chain 1
		b, err := hex.DecodeString("0ac45f1e6b8d47ac4c73aee62c52794b5898da9f")
		return &account.PureAccount{
			ChainId: 1,
			Address: b,
		}, err
	case 2: // ethereum chain 2
		b, err := hex.DecodeString("d051a43d3ea62afff3632bca3d5abf68bc6fd737")
		return &account.PureAccount{
			ChainId: 2,
			Address: b,
		}, err
	case 3: // tendermint chain 1
		b, err := hex.DecodeString("2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")
		return &account.PureAccount{
			ChainId: 3,
			Address: b,
		}, err
	case 4: // tendermint chain 2
		b, err := hex.DecodeString("2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")
		return &account.PureAccount{
			ChainId: 4,
			Address: b,
		}, err
	default:
		return nil, errors.New("not found")
	}
}

func TempSearchAccount(name string, chainId uint64) (types.Account, error) {
	switch name {
	case "":
		return nil, errors.New("nil name is not allowed")
	case "a1": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 1: // ethereum chain 1
			b, err := hex.DecodeString("d051a43d3ea62afff3632bca3d5abf68bc6fd737")
			return &account.PureAccount{
				ChainId: 1,
				Address: b,
			}, err
		case 2: // ethereum chain 2
			b, err := hex.DecodeString("93334ae4b2d42ebba8cc7c797bfeb02bfb3349d6")
			return &account.PureAccount{
				ChainId: 2,
				Address: b,
			}, err
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("604bdd2dd4b7e1b761e2ac96db99bb2bda386bb0d075b51a8f49c5103ebaa985")
			return &account.PureAccount{
				ChainId: 3,
				Address: b,
			}, err
		case 4: // ethereum chain 1
			b, err := hex.DecodeString("604bdd2dd4b7e1b761e2ac96db99bb2bda386bb0d075b51a8f49c5103ebaa985")
			return &account.PureAccount{
				ChainId: 4,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a2": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 1: // ethereum chain 1
			b, err := hex.DecodeString("47a1cdb6594d6efed3a6b917f2fbaa2bbcf61a2e")
			return &account.PureAccount{
				ChainId: 1,
				Address: b,
			}, err
		case 2: // ethereum chain 2
			b, err := hex.DecodeString("981739a13593980763de3353340617ef16da6354")
			return &account.PureAccount{
				ChainId: 2,
				Address: b,
			}, err
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("cfe900c7a56f87882f0e18e26851bce7b7e61ebeca6c4b235fa360d627dfac63")
			return &account.PureAccount{
				ChainId: 3,
				Address: b,
			}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("4f7a1b3d9f2f8f3e2c7e7729bc873fc55e607e47309941391a7a82673e563887")
			return &account.PureAccount{
				ChainId: 4,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a3": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
			return &account.PureAccount{
				ChainId: 3,
				Address: b,
			}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
			return &account.PureAccount{
				ChainId: 4,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a4": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("2333ffffffffffffffff2333ffffffffffffffff2333ffffffffffffffffffff")
			return &account.PureAccount{
				ChainId: 3,
				Address: b,
			}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("2333ffffffffffffffff2333ffffffffffffffff2333ffffffffffffffffffff")
			return &account.PureAccount{
				ChainId: 4,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	default:
		return nil, errors.New("not registered account")
	}
}

func GetTransactionProofType(chainId uint64) (uint16, error) {
	// ethereum chain 1
	switch chainId {
	case 0:
		return merkleprooftype.Invalid, errors.New("nil domain is not allowed")
	case 1: // ethereum chain 1
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 2: // ethereum chain 2
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 3: // tendermint chain 1
		return merkleprooftype.MerklePatriciaTrieUsingKeccak256, nil
	case 4: // tendermint chain 1
		return merkleprooftype.MerklePatriciaTrieUsingKeccak256, nil
	default:
		return merkleprooftype.Invalid, errors.New("not found")
	}
}
