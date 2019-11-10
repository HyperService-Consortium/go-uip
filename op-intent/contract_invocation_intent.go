package opintent

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	merkleprooftype "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	"github.com/HyperService-Consortium/go-uip/const/merkleproof_proposal_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"

	"github.com/tidwall/gjson"
)

func (ier *OpIntentInitializer) InitContractInvocationOpIntent(
	name string,
	subIntent json.RawMessage,
) (txs []*TransactionIntent, requiringMerkleProof []*MerkleProofProposal, err error) {
	var invokeIntent uiptypes.BaseContractInvocationOpIntent
	err = json.Unmarshal(subIntent, &invokeIntent)
	var tx *TransactionIntent
	if err != nil {
		return
	}
	if invokeIntent.Src == nil {
		return nil, nil, initializeError("src")
	}
	if len(invokeIntent.Dst) == 0 {
		return nil, nil, initializeError("dst")
	}
	if len(invokeIntent.FuncName) == 0 {
		return nil, nil, initializeError("function name")
	}
	var dstAddr []byte
	dstAddr, err = hex.DecodeString(invokeIntent.Dst)
	if err != nil {
		return
	}

	var srcInfo uiptypes.Account
	srcInfo, err = ier.accountProvider.Get(invokeIntent.Src.Name, invokeIntent.Src.ChainId)
	if err != nil {
		return
	}
	if tx, err = func() (*TransactionIntent, error) {
		var meta uiptypes.ContractInvokeMeta
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
			Dst:       dstAddr,
			Meta:      b,
			// temporary = 1000
			Amt:     "03e8",
			ChainID: uiptypes.ChainIDUnderlyingType(srcInfo.GetChainId()),
		}
		return tx, nil
	}(); err != nil {
		return
	}
	txs = append(txs, tx)
	var proposal []*MerkleProofProposal
	proposal, err = ier.parseContractInvokeProof(&invokeIntent)
	requiringMerkleProof = append(requiringMerkleProof, proposal...)

	return
}

func (ier *OpIntentInitializer) parseContractInvokeProof(intent *uiptypes.BaseContractInvocationOpIntent) (proposal []*MerkleProofProposal, err error) {
	var b []byte
	var txp transactionProofSourceDescription
	txp.ChainID = intent.Src.ChainId
	b, err = json.Marshal(&txp)
	if err != nil {
		return
	}
	var merkleproofType merkleprooftype.Type
	merkleproofType, err = ier.accountProvider.GetTransactionProofType(intent.Src.ChainId)
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
