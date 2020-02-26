package parser

import (
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
)


func (ier * Parser) _queryAccount(m lexer.AccountMap, a lexer.Account) (uip.Account, error) {
	switch a.GetType() {
	case token.NamespacedNameAccount:
		a := a.(token.NamespacedNameAccountI)
		if c, ok := m[a.GetName()]; ok && c != nil {
			if r, ok := c[a.GetChainID()]; ok {
				return r, nil
			}
		}
		r, err := ier.AccountBase.Get(a.GetName(), a.GetChainID())
		if err != nil {
			return nil, errorn.NewGetAccountFailed(err)
		}
		return r, nil
	case token.NameAccount:
		a := a.(token.NameAccountI)
		if c, ok := m[a.GetName()]; ok && c != nil {
			if r, ok := c[0]; ok {
				return r, nil
			} else {
				if ier.contextChainID == 0 {
					return nil, errorn.NewNoDeterminedChainID()
				}
				if r, ok := c[ier.contextChainID]; ok {
					return r, nil
				}
			}
		}
		return nil, errorn.NewAccountNotFound(a.GetName(), 0x20200202)
	case token.NamespacedRawAccount:
		return a.(token.NamespacedRawAccountI), nil
	case token.RawAccount:
		if ier.contextChainID == 0 {
			return nil, errorn.NewNoDeterminedChainID()
		}
		return &uip.AccountImpl{
			ChainId: ier.contextChainID,
			Address: a.(token.RawAccountI).GetAddress(),
		}, nil
	default:
		return nil, errorn.NewAccountTypeNotFound(int(a.GetType()))
	}
}
func (ier * Parser) queryAccount(a lexer.Account) (uip.Account, error) {
	return ier._queryAccount(ier.Program.AccountMapping, a)
}


func (ier * Parser) queryContract(a lexer.Account) (uip.Account, error) {
	return ier._queryAccount(ier.Program.ContractMapping, a)
}

func (ier * Parser) parsePayment(i lexer.Intent) (intents []uip.TxIntentI, err error) {
	paymentIntent := i.(*lexer.PaymentIntent)
	var srcInfo, dstInfo uip.Account
	var intent uip.TxIntentI

	srcInfo, err = ier.queryAccount(paymentIntent.Src)
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "src"})
	}

	dstInfo, err = ier.queryAccount(paymentIntent.Dst)
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "dst"})
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

func (ier * Parser) genPayment(
	src uip.Account, dst uip.Account, amt string, meta document.Document, ut UnitType.Type,
) (tx uip.TxIntentI, err error) {
	if src == nil {
		if src, err = ier.AccountBase.GetRelay(dst.GetChainId()); err != nil {
			return
		}
	} else {
		if dst, err = ier.AccountBase.GetRelay(src.GetChainId()); err != nil {
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

	translator, err := ier.ChainGetter.GetTranslator(dst.GetChainId())
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
