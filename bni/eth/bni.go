package bni

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"regexp"
	"strconv"

	"golang.org/x/crypto/sha3"

	"github.com/Myriad-Dreamin/go-ethabi"

	ethclient "github.com/Myriad-Dreamin/go-ves/net/eth_client"

	TransType "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	opintent "github.com/Myriad-Dreamin/go-uip/op-intent"
	types "github.com/Myriad-Dreamin/go-uip/types"
	gjson "github.com/tidwall/gjson"
)

type BN struct {
	signer types.Signer
}

func (bn *BN) MustWithSigner() bool {
	return true
}

func (bn *BN) RouteWithSigner(signer types.Signer) types.Router {
	nbn := new(BN)
	nbn.signer = signer
	return nbn
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
		//_ = meta
		return ContractInvocationDataABI(&meta, kvs)
		//return nil, errors.New("todo")
	default:
		return nil, errors.New("cant translate")
	}
}

func (bn *BN) CheckAddress(addr []byte) bool {
	return len(addr) == 32
}

func ContractInvocationDataABI(meta *types.ContractInvokeMeta, provedData map[string][]byte) ([]byte, error) {

	abiencoder := ethabi.NewEncoder()
	//Encodes(descs []string, vals []interface{})
	keccak := sha3.NewLegacyKeccak256()
	var descs []string
	var vals []interface{} = make([]interface{}, 0)
	var funcsig string = meta.FuncName + "("
	var err error
	for id, param := range meta.Params {
		t := param.Type
		if t == "address payable" || t == "contract" {
			t = "address"
		}
		funcsig += t
		descs = append(descs, t)
		constant := gjson.Get(string(param.Value), "constant")
		if !constant.Exists() {
			constant = gjson.Parse(string(provedData[gjson.Get(string(param.Value), "field").String()]))
		}
		vals, err = appendVal(vals, t, constant)
		if err != nil {
			return nil, err
		}
		if id == len(meta.Params)-1 {
			funcsig += ")"
		} else {
			funcsig += ","
		}
	}
	keccak.Write([]byte(funcsig))
	result := keccak.Sum([]byte(""))[0:4]
	if len(descs) > 0 {
		result2, err := abiencoder.Encodes(descs, vals)
		if err != nil {
			return nil, err
		}
		result = append(result, result2...)
	}
	return result, nil
}

func appendVal(vals []interface{}, t string, rawval gjson.Result) ([]interface{}, error) {
	typeRegex := regexp.MustCompile("([a-zA-Z]+)(([0-9]+)(x([0-9]+))?)?")
	matches := typeRegex.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return vals, fmt.Errorf("invalid type '%v'", t)
	}
	parsedType := matches[0]
	var varSize int
	if len(parsedType[3]) > 0 {
		var err error
		varSize, err = strconv.Atoi(parsedType[2]) //ParseInt(sparsedType[2], 10, 0) //strconv.Atoi()
		if err != nil {
			return vals, fmt.Errorf("abi: error parsing variable size: %v", err)
		}
	} else {
		if parsedType[0] == "uint" || parsedType[0] == "int" {
			// this should fail because it means that there's something wrong with
			// the abi type (the compiler should always format it to the size...always)
			return vals, fmt.Errorf("unsupported arg type: %s", t)
		}
	}
	switch varType := parsedType[1]; varType {
	case "int":
		switch varSize {
		case 8:
			vals = append(vals, int8(rawval.Int()))
		case 16:
			vals = append(vals, int16(rawval.Int()))
		case 32:
			vals = append(vals, int32(rawval.Int()))
		case 64:
			vals = append(vals, int64(rawval.Int()))
		case 256:
			str := rawval.String()
			val, ok := big.NewInt(0).SetString(str, 10)
			if !ok {
				return vals, fmt.Errorf("Invalid int value")
			}
			vals = append(vals, val)
		}
	case "uint":
		switch varSize {
		case 8:
			vals = append(vals, uint8(rawval.Int()))
		case 16:
			vals = append(vals, uint16(rawval.Int()))
		case 32:
			vals = append(vals, uint32(rawval.Int()))
		case 64:
			vals = append(vals, uint64(rawval.Int()))
		case 256:
			str := rawval.String()
			val, ok := big.NewInt(0).SetString(str, 10)
			if !ok {
				return vals, fmt.Errorf("Invalid uint value")
			}
			vals = append(vals, val)
		}
	case "bool":
		vals = append(vals, rawval.Bool())
	case "string":
		vals = append(vals, rawval.String())
	default:
		return vals, fmt.Errorf("unsupported arg type: %s", t)
	}
	//vals = append(vals, ori["constant"])
	return vals, nil
}
