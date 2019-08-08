package opintent

import (
	"encoding/json"
	"errors"

	merkleprooftype "github.com/Myriad-Dreamin/go-uip/const/merkle-proof-type"
	merkleproof_proposal_type "github.com/Myriad-Dreamin/go-uip/const/merkleproof_proposal_type"
	trans_type "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	value_type "github.com/Myriad-Dreamin/go-uip/const/value_type"
	types "github.com/Myriad-Dreamin/go-uip/types"

	gjson "github.com/tidwall/gjson"
)

func (ier *OpIntentInitializer) InitContractInvocationOpIntent(
	name string,
	subIntent json.RawMessage,
) (txs []*TransactionIntent, requiringMerkleProof []*MerkleProofProposal, err error) {
	var invokeIntent types.BaseContractInvocationOpIntent
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
		var meta types.ContractInvokeMeta
		meta.Code = invokeIntent.Code
		meta.FuncName = invokeIntent.FuncName
		meta.Params = invokeIntent.Params
		meta.Meta = invokeIntent.Meta
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
			ChainID: srcInfo.GetChainId(),
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

func parseContractInvokeProof(intent *types.BaseContractInvocationOpIntent) (proposal []*MerkleProofProposal, err error) {
	var b []byte
	var txp transactionProofSourceDescription
	txp.ChainID = intent.Src.ChainId
	b, err = json.Marshal(&txp)
	if err != nil {
		return
	}
	var merkleproofType merkleprooftype.Type
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

	var intDesc value_type.Type
	for _, param := range intent.Params {
		if intDesc = value_type.FromString(param.Type); intDesc == value_type.Unknown {
			return nil, errors.New("unknown type: " + param.Type)
		}

		result := gjson.ParseBytes(param.Value)
		if !result.Get("constant").Exists() {
			if result.Get("contract").Exists() &&
				result.Get("pos").Exists() &&
				result.Get("field").Exists() {
				proposal = append(proposal, &MerkleProofProposal{
					DescriptionType:   merkleproof_proposal_type.DataProof,
					MerkleProofType:   merkleproofType,
					ValueType:         intDesc,
					SourceDescription: param.Value,
				})
			} else {
				return nil, errors.New("no enough info of source description")
			}
		}
	}
	return
}
