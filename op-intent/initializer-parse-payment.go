package opintent

import (
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (ier *Initializer) parsePayment(i RawIntentI) (intents []uip.TxIntentI, err error) {
	paymentIntent := i.GetSub().(*BasePaymentOpIntent)
	var srcInfo, dstInfo uip.Account
	var intent uip.TxIntentI
	srcInfo, err = ier.accountBase.Get(
		paymentIntent.Src.Name, paymentIntent.Src.ChainId)
	if err != nil {
		return nil, newGetAccountFailed(err).Desc(AtOpIntentField{"src"})
	}
	dstInfo, err = ier.accountBase.Get(
		paymentIntent.Dst.Name, paymentIntent.Dst.ChainId)
	if err != nil {
		return nil, newGetAccountFailed(err).Desc(AtOpIntentField{"dst"})
	}

	if intent, err = ier.genPayment(srcInfo, nil, paymentIntent.Amount, paymentIntent.Meta, paymentIntent.Unit); err != nil {
		return nil, newGenPaymentError(err).Desc(AtOpIntentField{"src"})
	}
	intents = append(intents, intent)
	if intent, err = ier.genPayment(nil, dstInfo, paymentIntent.Amount, paymentIntent.Meta, paymentIntent.Unit); err != nil {
		return nil, newGenPaymentError(err).Desc(AtOpIntentField{"dst"})
	}
	intents = append(intents, intent)

	return
}

// type PaymentMeta struct {
// 	OpTypeString string `json:"op_type"`
// }

// var pm = []byte(`{"op_type": "transact"}`)

type transactionProofSourceDescription struct {
	ChainID uint64 `json:"chain_id"`
}

func (ier *Initializer) genPayment(
	src uip.Account, dst uip.Account, amt string, meta []byte, ut UnitType.Type,
) (tx uip.TxIntentI, err error) {
	if src == nil {
		if src, err = ier.accountBase.GetRelay(dst.GetChainId()); err != nil {
			return
		}
	} else {
		if dst, err = ier.accountBase.GetRelay(src.GetChainId()); err != nil {
			return
		}
	}

	tx = new(TxIntentImpl)

	tx.SetIntent(&TransactionIntent{
		Src:       src.GetAddress(),
		Dst:       dst.GetAddress(),
		TransType: trans_type.Payment,
		Amt:       amt,
		Meta:      meta,
		ChainID:   dst.GetChainId(),
	})

	translator, err := ier.chainGetter.GetTranslator(dst.GetChainId())
	if err != nil {
		return nil, newGetTranslatorError(err).Desc(AtChainID{dst.GetChainId()})
	}

	tx, err = translator.ParseTransactionIntent(tx)
	if err != nil {
		return nil, newParseTransactionIntentError(err)
	}

	//if err = Parse

	//var txp transactionProofSourceDescription
	//txp.ChainID = dst.GetChainId()
	//b, err := json.marshal(&txp)
	//if err != nil {
	//	return
	//}
	//var merkleproofType merkle_proof.Type
	//merkleproofType, err = ier.accountBase.GetTransactionProofType(dst.GetChainId())
	//if err != nil {
	//	return
	//}
	//tx.proposals = append(tx.proposals, MerkleProofProposal{
	//	DescriptionType:   merkleproof_proposal_type.TransactionProof,
	//	MerkleProofType:   merkleproofType,
	//	ValueType:         value_type.Bytes,
	//	SourceDescription: b,
	//})
	return
}
