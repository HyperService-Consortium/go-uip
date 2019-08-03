package bni

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/Myriad-Dreamin/go-ethabi"

	ethclient "github.com/Myriad-Dreamin/go-ves/net/eth_client"

	TransType "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	valuetype "github.com/Myriad-Dreamin/go-uip/const/value_type"
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

func (bn *BN) Route(intent *types.TransactionIntent, kvGetter types.KVGetter) ([]byte, error) {
	// todo
	onChainTransaction, err := bn.Translate(intent, kvGetter)
	if err != nil {
		return nil, err
	}
	return bn.RouteRaw(intent.ChainID, onChainTransaction)
}

func (bn *BN) Translate(
	intent *opintent.TransactionIntent,
	kvGetter types.KVGetter,
) ([]byte, error) {
	switch intent.TransType {
	case TransType.Payment:
		return json.Marshal(map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_sendTransaction",
			"params": []interface{}{
				map[string]interface{}{
					"from":  decoratePrefix(hex.EncodeToString(intent.Src)),
					"to":    decoratePrefix(hex.EncodeToString(intent.Dst)),
					"value": decoratePrefix(intent.Amt),
				},
			},
			"id": 1,
		})
	case TransType.ContractInvoke:
		var meta types.ContractInvokeMeta
		err := json.Unmarshal(intent.Meta, &meta)
		if err != nil {
			return nil, err
		}
		//_ = meta

		data, err := ContractInvocationDataABI(&meta, kvGetter)
		if err != nil {
			return nil, err
		}

		hexdata := hex.EncodeToString(data)
		// meta.FuncName

		return json.Marshal(map[string]interface{}{
			"jsonrpc": "2.0",
			"method":  "eth_sendTransaction",
			"params": []interface{}{
				map[string]interface{}{
					"from":  decoratePrefix(hex.EncodeToString(intent.Src)),
					"to":    decoratePrefix(hex.EncodeToString(intent.Dst)),
					"value": decoratePrefix(intent.Amt),
					"data":  decoratePrefix(hexdata),
				},
			},
			"id": 1,
		})
		//return nil, errors.New("todo")
	default:
		return nil, errors.New("cant translate")
	}
}

func (bn *BN) CheckAddress(addr []byte) bool {
	return len(addr) == 32
}

func (bn *BN) GetStorageAt(chainID uint64, typeID uint16, contractAddress []byte, pos []byte, desc []byte) (interface{}, error) {

	ci, err := SearchChainId(chainID)
	if err != nil {
		return nil, err
	}

	switch typeID {
	case valuetype.Bool:
		s, err := ethclient.NewEthClient(ci.GetHost()).GetStorageAt(contractAddress, pos, "latest")
		if err != nil {
			return nil, err
		}

		b, err := hex.DecodeString(s[2:])
		if err != nil {
			return nil, err
		}
		bs, err := ethabi.NewDecoder().Decodes([]string{"bool"}, b)
		return bs[0], nil
	case valuetype.Uint256:
		s, err := ethclient.NewEthClient(ci.GetHost()).GetStorageAt(contractAddress, pos, "latest")
		if err != nil {
			return nil, err
		}

		b, err := hex.DecodeString(s[2:])
		if err != nil {
			return nil, err
		}
		bs, err := ethabi.NewDecoder().Decodes([]string{"uint256"}, b)
		return bs[0], nil
	}

	return nil, nil
}

func ContractInvocationDataABI(meta *types.ContractInvokeMeta, kvGetter types.KVGetter) ([]byte, error) {

	abiencoder := ethabi.NewEncoder()
	//Encodes(descs []string, vals []interface{})
	keccak := sha3.NewLegacyKeccak256()
	var descs []string
	var vals []interface{} = make([]interface{}, 0)
	var funcsig string = meta.FuncName + "("
	//var err error
	for id, param := range meta.Params {
		t := param.Type
		if t == "address payable" || t == "contract" {
			t = "address"
		}
		funcsig += t
		descs = append(descs, t)
		constant := gjson.Get(string(param.Value), "constant")
		if !constant.Exists() {
			k := gjson.Get(string(param.Value), "field")
			if !k.Exists() {
				return nil, errors.New("field not found")
			}
			v, err := kvGetter.Get([]byte(k.String()))
			if err != nil {
				return nil, err
			}
			constant = gjson.ParseBytes(v)
		}
		val, err := appendVal(t, constant)
		if err != nil {
			return nil, err
		}
		fmt.Println("SHHHHH", val)
		vals = append(vals, val)
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

func appendSliceVal(brcnt int, t string, rawval gjson.Result) (interface{}, error) {
	var err error
	var ret interface{}
	i := strings.LastIndex(t, "[")
	// grab the last cell and create a type from there
	sliced := t[i:]
	// grab the slice size with regexp
	re := regexp.MustCompile("[0-9]+")
	intz := re.FindAllString(sliced, -1)
	var arr []gjson.Result = rawval.Array()
	if len(intz) == 1 {
		// is a array
		siz, err := strconv.Atoi(intz[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing variable size: %v", err)
		}
		if siz != len(arr) {
			return nil, fmt.Errorf("array length not match")
		}
	} else if len(intz) != 0 {
		return nil, fmt.Errorf("invalid formatting of array type")
	}
	/////////////////////////////////////////////////////////
	t = t[:i]
	typeRegex := regexp.MustCompile("([a-zA-Z]+)(([0-9]+)(x([0-9]+))?)?")
	matches := typeRegex.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("invalid type '%v'", t)
	}
	parsedType := matches[0]
	var varSize int
	if len(parsedType[3]) > 0 {
		var err error
		varSize, err = strconv.Atoi(parsedType[2]) //ParseInt(sparsedType[2], 10, 0) //strconv.Atoi()
		if err != nil {
			return nil, fmt.Errorf("error parsing variable size: %v", err)
		}
	} else {
		if parsedType[0] == "uint" || parsedType[0] == "int" {
			// this should fail because it means that there's something wrong with
			// the abi type (the compiler should always format it to the size...always)
			return nil, fmt.Errorf("unsupported arg type: %s", t)
		}
	}
	switch varType := parsedType[1]; varType {
	case "int":
		switch varSize {
		case 8:
			ret = make([]int8, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]int8)[id] = elem.(int8)
				if err != nil {
					return nil, err
				}
			}
		case 16:
			ret = make([]int16, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]int16)[id] = elem.(int16)
				if err != nil {
					return nil, err
				}
			}
		case 32:
			ret = make([]int32, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]int32)[id] = elem.(int32)
				if err != nil {
					return nil, err
				}
			}
		case 64:
			ret = make([]int64, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]int64)[id] = elem.(int64)
				if err != nil {
					return nil, err
				}
			}
		case 256:
			ret = make([]big.Int, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]big.Int)[id] = elem.(big.Int)
				if err != nil {
					return nil, err
				}
			}
		}
	case "uint":
		switch varSize {
		case 8:
			ret = make([]uint8, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]uint8)[id] = elem.(uint8)
				if err != nil {
					return nil, err
				}
			}
		case 16:
			ret = make([]uint16, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]uint16)[id] = elem.(uint16)
				if err != nil {
					return nil, err
				}
			}
		case 32:
			ret = make([]uint32, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]uint32)[id] = elem.(uint32)
				if err != nil {
					return nil, err
				}
			}
		case 64:
			ret = make([]uint64, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]uint64)[id] = elem.(uint64)
				if err != nil {
					return nil, err
				}
			}
		case 256:
			ret = make([]big.Int, len(arr))
			for id, obj := range arr {
				elem, err := appendVal(t, obj)
				ret.([]big.Int)[id] = elem.(big.Int)
				if err != nil {
					return nil, err
				}
			}
		}
	case "bool":
		ret = make([]bool, len(arr))
		for id, obj := range arr {
			elem, err := appendVal(t, obj)
			ret.([]bool)[id] = elem.(bool)
			if err != nil {
				return nil, err
			}
		}
	case "address":
		ret = make([][20]byte, len(arr))
		for id, obj := range arr {
			elem, err := appendVal(t, obj)
			ret.([][20]byte)[id] = elem.([20]byte)
			if err != nil {
				return nil, err
			}
		}
	case "string":
		ret = make([]string, len(arr))
		for id, obj := range arr {
			elem, err := appendVal(t, obj)
			ret.([]string)[id] = elem.(string)
			if err != nil {
				return nil, err
			}
		}
	case "bytes":
		ret = make([][]byte, len(arr))
		for id, obj := range arr {
			elem, err := appendVal(t, obj)
			ret.([][]byte)[id] = elem.([]byte)
			if err != nil {
				return nil, err
			}
		}
	default:
		return nil, fmt.Errorf("unsupported arg type: %s", t)
	}
	///////////////////////////////////////////
	fmt.Println("ARR", ret, reflect.TypeOf(ret).Elem())
	return ret, err
}

func appendVal(t string, rawval gjson.Result) (interface{}, error) {
	// check that array brackets are equal if they exist
	brcnt := strings.Count(t, "[")
	if brcnt != strings.Count(t, "]") {
		return nil, fmt.Errorf("invalid arg type in abi")
	}
	var err error
	var ret interface{}

	// if there are brackets, get ready to go into slice/array mode and
	// recursively create the type
	if brcnt != 0 {
		if brcnt != 1 {
			return nil, fmt.Errorf("unsupported arg type: %s", t)
		}
		return appendSliceVal(brcnt, t, rawval)
	}
	typeRegex := regexp.MustCompile("([a-zA-Z]+)(([0-9]+)(x([0-9]+))?)?")
	matches := typeRegex.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("invalid type '%v'", t)
	}
	parsedType := matches[0]
	var varSize int
	if len(parsedType[3]) > 0 {
		var err error
		varSize, err = strconv.Atoi(parsedType[2]) //ParseInt(sparsedType[2], 10, 0) //strconv.Atoi()
		if err != nil {
			return nil, fmt.Errorf("error parsing variable size: %v", err)
		}
	} else {
		if parsedType[0] == "uint" || parsedType[0] == "int" {
			// this should fail because it means that there's something wrong with
			// the abi type (the compiler should always format it to the size...always)
			return nil, fmt.Errorf("unsupported arg type: %s", t)
		}
	}
	switch varType := parsedType[1]; varType {
	case "int":
		switch varSize {
		case 8:
			ret = int8(rawval.Int())
		case 16:
			ret = int16(rawval.Int())
		case 32:
			ret = int32(rawval.Int())
		case 64:
			ret = int64(rawval.Int())
		case 256:
			str := rawval.String()
			val, ok := big.NewInt(0).SetString(str, 10)
			if !ok {
				return nil, fmt.Errorf("Invalid int value")
			}
			ret = val
		}
	case "uint":
		switch varSize {
		case 8:
			ret = uint8(rawval.Int())
		case 16:
			ret = uint16(rawval.Int())
		case 32:
			ret = uint32(rawval.Int())
		case 64:
			ret = uint64(rawval.Int())
		case 256:
			str := rawval.String()
			val, ok := big.NewInt(0).SetString(str, 10)
			if !ok {
				return nil, fmt.Errorf("Invalid int value")
			}
			ret = val
		}
	case "bool":
		ret = rawval.Bool()
	case "address":
		addr := rawval.String()
		retsli, err := hex.DecodeString(addr[2:])
		if err != nil {
			return nil, err
		}
		if len(retsli) != 20 {
			return nil, fmt.Errorf("invalid address value")
		}
		var rett [20]byte
		for i := 0; i < 20; i++ {
			rett[i] = retsli[i]
		}
		ret = rett
	case "string":
		ret = rawval.String()
	case "bytes":
		if varSize != 0 {
			return nil, fmt.Errorf("unsupported arg type: %s", t)
		}
		tmpret := []byte("\"" + rawval.String() + "\"")
		var retval []byte
		err = json.Unmarshal(tmpret, &retval)
		if err != nil {
			return nil, err
		}
		ret = retval
	default:
		return nil, fmt.Errorf("unsupported arg type: %s", t)
	}
	return ret, nil
}
