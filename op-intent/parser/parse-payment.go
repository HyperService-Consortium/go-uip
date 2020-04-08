package parser

import (
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	UnitType "github.com/HyperService-Consortium/go-uip/const/unit_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (ier *Parser) _queryAccount(m lexer_types.AccountMap, a lexer_types.Account, contextChainID uip.ChainIDUnderlyingType) (uip.Account, error) {
	switch a.GetType() {
	case token_type.NamespacedNameAccount:
		a := a.(token_types.NamespacedNameAccountI)
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
	case token_type.NameAccount:
		a := a.(token_types.NameAccountI)
		if c, ok := m[a.GetName()]; ok && c != nil {
			if r, ok := c[0]; ok {
				return r, nil
			} else {
				if contextChainID == 0 {
					return nil, errorn.NewNoDeterminedChainID()
				}
				if r, ok := c[contextChainID]; ok {
					return r, nil
				}
			}
		}
		return nil, errorn.NewAccountNotFound(a.GetName(), 0x20200202)
	case token_type.NamespacedRawAccount:
		return a.(token_types.NamespacedRawAccountI), nil
	case token_type.RawAccount:
		if contextChainID == 0 {
			return nil, errorn.NewNoDeterminedChainID()
		}
		return &uip.AccountImpl{
			ChainId: contextChainID,
			Address: a.(token_types.RawAccountI).GetAddress(),
		}, nil
	default:
		return nil, errorn.NewAccountTypeNotFound(int(a.GetType()))
	}
}

func (ier *Parser) QueryAccount(a lexer_types.Account) (uip.Account, error) {
	return ier._queryAccount(ier.Program.AccountMapping, a, 0)
}

func (ier *Parser) QueryContract(a lexer_types.Account) (uip.Account, error) {
	return ier._queryAccount(ier.Program.ContractMapping, a, 0)
}

func (ier *Parser) queryAccountWCtx(a lexer_types.Account, contextChainID uip.ChainIDUnderlyingType) (uip.Account, error) {
	return ier._queryAccount(ier.Program.AccountMapping, a, contextChainID)
}

func (ier *Parser) queryContractWCtx(a lexer_types.Account, contextChainID uip.ChainIDUnderlyingType) (uip.Account, error) {
	return ier._queryAccount(ier.Program.ContractMapping, a, contextChainID)
}

func (ier *Parser) parsePayment(paymentIntent *lexer.PaymentIntent) (intents []uip.TxIntentI, err error) {

	var srcInfo, dstInfo uip.Account
	var intent uip.TxIntentI

	srcInfo, err = ier.QueryAccount(paymentIntent.Src)
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "src"})
	}

	dstInfo, err = ier.QueryAccount(paymentIntent.Dst)
	if err != nil {
		return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentField{Field: "dst"})
	}

	if intent, err = ier.genPayment(
		paymentIntent.GetName()+".cna", srcInfo, nil, paymentIntent.Amount, paymentIntent.Meta, paymentIntent.Unit); err != nil {
		return nil, errorn.NewGenPaymentError(err).Desc(errorn.AtOpIntentField{Field: "src"})
	}
	intents = append(intents, intent)
	if intent, err = ier.genPayment(
		paymentIntent.GetName()+".cnb", nil, dstInfo, paymentIntent.Amount, paymentIntent.Meta, paymentIntent.Unit); err != nil {
		return nil, errorn.NewGenPaymentError(err).Desc(errorn.AtOpIntentField{Field: "dst"})
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

func (ier *Parser) genPayment(
	name string, src uip.Account, dst uip.Account, amt string, meta document.Document, ut UnitType.Type,
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

	m, err := meta.RawBytes()
	if err != nil {
		return nil, err
	}

	tx = newIntent(&instruction.TransactionIntent{
		Src:       src.GetAddress(),
		Dst:       dst.GetAddress(),
		TransType: trans_type.Payment,
		Amt:       amt,
		Meta:      m,
		ChainID:   dst.GetChainId(),
	}, name)

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
