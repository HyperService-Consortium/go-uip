package opintent

import (
	"encoding/json"
	"errors"

	types "github.com/HyperService-Consortium/go-uip/uiptypes"

	merkleprooftype "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	merkleproof_proposal_type "github.com/HyperService-Consortium/go-uip/const/merkleproof_proposal_type"
	trans_type "github.com/HyperService-Consortium/go-uip/const/trans_type"
	unit_type "github.com/HyperService-Consortium/go-uip/const/unit_type"
	value_type "github.com/HyperService-Consortium/go-uip/const/value_type"
)

func (ier *OpIntentInitializer) InitPaymentOpIntent(
	name string,
	subIntent json.RawMessage,
) (txs []*TransactionIntent, requiringMerkleProof []*MerkleProofProposal, err error) {
	var paymentIntent types.BasePaymentOpIntent
	err = json.Unmarshal(subIntent, &paymentIntent)
	var tx *TransactionIntent
	var proposal []*MerkleProofProposal
	if err != nil {
		return
	}
	if paymentIntent.Src == nil {
		return nil, nil, initializeError("src")
	}
	if paymentIntent.Dst == nil {
		return nil, nil, initializeError("dst")
	}
	if len(paymentIntent.Amount) == 0 {
		return nil, nil, initializeError("amount")
	}
	t, ok := unit_type.Mapping[paymentIntent.UnitString]
	if !ok {
		return nil, nil, errors.New("unknown unit type")
	}
	var srcInfo, dstInfo types.Account
	srcInfo, err = ier.accountProvider.Get(paymentIntent.Src.Name, paymentIntent.Src.ChainId)
	if err != nil {
		return
	}
	dstInfo, err = ier.accountProvider.Get(paymentIntent.Dst.Name, paymentIntent.Dst.ChainId)
	if err != nil {
		return
	}
	if tx, proposal, err = ier.genPayment(srcInfo, nil, paymentIntent.Amount, paymentIntent.Meta, t); err != nil {
		return
	}
	requiringMerkleProof = append(requiringMerkleProof, proposal...)
	txs = append(txs, tx)
	if tx, proposal, err = ier.genPayment(nil, dstInfo, paymentIntent.Amount, paymentIntent.Meta, t); err != nil {
		return
	}
	requiringMerkleProof = append(requiringMerkleProof, proposal...)
	txs = append(txs, tx)
	// cinfo, err = SearchChainId(paymentIntent.Src.ChainId)
	// if err != nil {
	// 	return nil, err
	// }
	// var processor ProcessorInterface
	// switch cinfo.GetChainType() {
	// case chain_type.Ethereum:
	// 	processor = eth_processor.GetProcessor()
	// default:
	// 	return nil, errors.New("unsupport chain_type")
	// }
	// if !processor.CheckAddress(paymentIntent.Src.)
	return
}

// type PaymentMeta struct {
// 	OpType string `json:"op_type"`
// }

// var pm = []byte(`{"op_type": "transact"}`)

type transactionProofSourceDescription struct {
	ChainID uint64 `json:"chain_id"`
}

func (ier *OpIntentInitializer) genPayment(
	src types.Account, dst types.Account, amt string, meta []byte, utid unit_type.Type,
) (tx *TransactionIntent, proposal []*MerkleProofProposal, err error) {
	if src == nil {
		if src, err = ier.accountProvider.GetRelay(dst.GetChainId()); err != nil {
			return nil, nil, err
		}
	} else {
		if dst, err = ier.accountProvider.GetRelay(src.GetChainId()); err != nil {
			return nil, nil, err
		}
	}

	tx = &TransactionIntent{
		Src:       src.GetAddress(),
		Dst:       dst.GetAddress(),
		TransType: trans_type.Payment,
		Amt:       amt,
		Meta:      meta,
		ChainID:   dst.GetChainId(),
	}

	var txp transactionProofSourceDescription
	txp.ChainID = dst.GetChainId()
	b, err := json.Marshal(&txp)
	if err != nil {
		return nil, nil, err
	}
	var merkleproofType merkleprooftype.Type
	merkleproofType, err = ier.accountProvider.GetTransactionProofType(dst.GetChainId())
	if err != nil {
		return nil, nil, err
	}
	proposal = append(proposal, &MerkleProofProposal{
		DescriptionType:   merkleproof_proposal_type.TransactionProof,
		MerkleProofType:   merkleproofType,
		ValueType:         value_type.Bytes,
		SourceDescription: b,
	})
	return
}
