package opintent

import (
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (ier *Initializer) parsePayment(i lexer.Intent) (intents []uip.TxIntentI, err error) {
	paymentIntent := i.(*lexer.PaymentIntent)
	var srcInfo, dstInfo uip.Account
	var intent uip.TxIntentI
	srcInfo, err = ier.accountBase.Get(
		paymentIntent.Src.Name, paymentIntent.Src.ChainId)
	if err != nil {
		return nil, errorn.NewGetAccountFailed(err).Desc(errorn.AtOpIntentField{"src"})
	}
	dstInfo, err = ier.accountBase.Get(
		paymentIntent.Dst.Name, paymentIntent.Dst.ChainId)
	if err != nil {
		return nil, errorn.NewGetAccountFailed(err).Desc(errorn.AtOpIntentField{"dst"})
	}

	if intent, err = ier.genPayment(srcInfo, nil, paymentIntent.Amount, paymentIntent.Meta, paymentIntent.Unit); err != nil {
		return nil, errorn.NewGenPaymentError(err).Desc(errorn.AtOpIntentField{"src"})
	}
	intents = append(intents, intent)
	if intent, err = ier.genPayment(nil, dstInfo, paymentIntent.Amount, paymentIntent.Meta, paymentIntent.Unit); err != nil {
		return nil, errorn.NewGenPaymentError(err).Desc(errorn.AtOpIntentField{"dst"})
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
	src uip.Account, dst uip.Account, amt string, meta document.Document, ut UnitType.Type,
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

	m ,err := meta.RawBytes()
	if err != nil {
		return nil, err
	}
	tx.SetIntent(&TransactionIntent{
		Src:       src.GetAddress(),
		Dst:       dst.GetAddress(),
		TransType: trans_type.Payment,
		Amt:       amt,
		Meta:      m,
		ChainID:   dst.GetChainId(),
	})

	translator, err := ier.chainGetter.GetTranslator(dst.GetChainId())
	if err != nil {
		return nil, errorn.NewGetTranslatorError(err).Desc(errorn.AtChainID{dst.GetChainId()})
	}

	tx, err = translator.ParseTransactionIntent(tx)
	if err != nil {
		return nil, errorn.NewParseTransactionIntentError(err)
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
