package opintent

import (
	"encoding/json"
	"fmt"
	"math/big"
	"regexp"
	"strconv"

	"golang.org/x/crypto/sha3"

	"github.com/Myriad-Dreamin/go-ethabi"
	merkleproof_proposal_type "github.com/Myriad-Dreamin/go-uip/const/merkleproof_proposal_type"
	merkleproof_type "github.com/Myriad-Dreamin/go-uip/const/merkleproof_type"
	trans_type "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	value_type "github.com/Myriad-Dreamin/go-uip/const/value_type"
	types "github.com/Myriad-Dreamin/go-uip/types"
	"github.com/tidwall/gjson"
)

type BaseContractInvocationOpIntent struct {
	Src      *RawAccountInfo `json:"invoker"`       // key
	Dst      []byte          `json:"contract_addr"` // key
	Code     []byte          `json:"contract_code"` // key
	FuncName string          `json:"func"`
	Params   []RawParams     `json:"parameters"`
}

type contractInvokeMeta struct {
	Code     []byte      `json:"contract_code"` // key
	FuncName string      `json:"func"`
	Params   []RawParams `json:"parameters"`
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

func ContractInvocationDataABI(meta *contractInvokeMeta, provedData map[string][]byte) ([]byte, error) {

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

func (ier *OpIntentInitializer) InitContractInvocationOpIntent(
	name string,
	subIntent json.RawMessage,
) (txs []*TransactionIntent, requiringMerkleProof []*MerkleProofProposal, err error) {
	var invokeIntent BaseContractInvocationOpIntent
	err = json.Unmarshal(subIntent, &invokeIntent)
	var tx *TransactionIntent
	if err != nil {
		return
	}
	if invokeIntent.Src == nil {
		return nil, nil, initializeError("src")
	}
	if invokeIntent.Dst == nil {
		return nil, nil, initializeError("dst")
	}
	if len(invokeIntent.FuncName) == 0 {
		return nil, nil, initializeError("function name")
	}
	var srcInfo types.Account
	srcInfo, err = TempSearchAccount(invokeIntent.Src.Name, invokeIntent.Src.ChainId)
	if err != nil {
		return
	}
	if tx, err = func() (*TransactionIntent, error) {
		var meta contractInvokeMeta
		meta.Code = invokeIntent.Code
		meta.FuncName = invokeIntent.FuncName
		meta.Params = invokeIntent.Params
		var b, err = json.Marshal(&meta)
		if err != nil {
			return nil, err
		}

		tx := &TransactionIntent{
			TransType: trans_type.ContractInvoke,
			Src:       srcInfo.GetAddress(),
			Dst:       invokeIntent.Dst,
			Meta:      b,
			// temporary = 1000
			Amt:     "0x3e8",
			ChainId: srcInfo.GetChainId(),
		}
		return tx, nil
	}(); err != nil {
		return
	}
	txs = append(txs, tx)
	var proposal []*MerkleProofProposal
	proposal, err = parseContractInvokeProof(&invokeIntent)
	requiringMerkleProof = append(requiringMerkleProof, proposal...)

	return
}

func parseContractInvokeProof(intent *BaseContractInvocationOpIntent) (proposal []*MerkleProofProposal, err error) {
	var b []byte
	var txp transactionProofSourceDescription
	txp.ChainId = intent.Src.ChainId
	b, err = json.Marshal(&txp)
	if err != nil {
		return
	}
	var merkleproofType merkleproof_type.Type
	merkleproofType, err = getTransactionProofType(intent.Src.ChainId)
	if err != nil {
		return
	}

	proposal = append(proposal, &MerkleProofProposal{
		DescriptionType:   merkleproof_proposal_type.TransactionProof,
		MerkleProofType:   merkleproofType,
		ValueType:         value_type.Bytes,
		SourceDescription: b,
	})

	// todo: params
	for _, param := range intent.Params {
		_ = param
	}
	return
}
