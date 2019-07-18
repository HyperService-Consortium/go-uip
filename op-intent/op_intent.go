package opintent

import (
	"encoding/json"

	types "github.com/Myriad-Dreamin/go-uip/types"
)

type OpIntentInitializer struct {
}

func NewOpIntentInitializer() *OpIntentInitializer {
	return new(OpIntentInitializer)
}

type BaseOpIntent struct {
	Name   string `json:"name"`
	OpType string `json:"op_type"`
}

func (ier *OpIntentInitializer) InitOpIntent(opIntents types.OpIntents) (transactionIntents []*TransactionIntent, err error) {
	contents, rawDependencies := opIntents.GetContents(), opIntents.GetDependencies()
	var intent BaseOpIntent
	var rtx [][]*TransactionIntent
	var proposals [][]*MerkleProofProposal
	var tx []*TransactionIntent
	var proposal []*MerkleProofProposal
	for _, content := range contents {
		err = json.Unmarshal(content, &intent)
		if err != nil {
			return nil, err
		}
		switch intent.OpType {
		case "Payment":
			if tx, proposal, err = ier.InitPaymentOpIntent(intent.Name, content); err != nil {
				return nil, err
			} else {
				rtx = append(rtx, tx)
				proposals = append(proposals, proposal)
			}

		case "ContractInvocation":
			if tx, proposal, err = ier.InitContractInvocationOpIntent(intent.Name, content); err != nil {
				return nil, err
			} else {
				rtx = append(rtx, tx)
				proposals = append(proposals, proposal)
			}
			// if tx, err = ier.InitContractInvocationOpIntent(intent.Name, intent.SubIntent); err != nil {
			// 	return nil, err
			// } else {
			// 	rtx = append(rtx, tx)
			// }

		default:
			return nil, invalidOpType
		}
	}
	if err = ier.TopologicalSort(rtx, rawDependencies); err != nil {
		return nil, err
	}
	for _, rt := range rtx {
		transactionIntents = append(transactionIntents, rt...)
	}
	return
}
