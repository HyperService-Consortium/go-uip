package opintent

import (
	"encoding/json"

	merkleproof_proposal_type "github.com/Myriad-Dreamin/go-uip/const/merkleproof_proposal_type"
	merkleproof_type "github.com/Myriad-Dreamin/go-uip/const/merkleproof_type"
	trans_type "github.com/Myriad-Dreamin/go-uip/const/trans_type"
	value_type "github.com/Myriad-Dreamin/go-uip/const/value_type"
	types "github.com/Myriad-Dreamin/go-uip/types"
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
