package parser

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/const/merkleproof_proposal_type"
	"github.com/HyperService-Consortium/go-uip/const/token_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/HyperService-Consortium/go-uip/internal/token_types"
	"github.com/HyperService-Consortium/go-uip/op-intent/instruction"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
)

func (ier *Parser) parseContractInvocation(invokeIntent *lexer.InvokeIntent) (intents []uip.TxIntentI, err error) {

	var srcInfo, dstInfo uip.Account
	var intent uip.TxIntentI
	srcInfo, err = ier.QueryAccount(invokeIntent.Src)
	if err != nil {
		return nil, errorn.NewGetAccountFailed(err).Desc(errorn.AtOpIntentField{Field: "src"})
	}
	dstInfo, err = ier.queryContractWCtx(invokeIntent.Dst, srcInfo.GetChainId())
	if err != nil {
		return nil, errorn.NewGetAccountFailed(err).Desc(errorn.AtOpIntentField{Field: "dst"})
	}

	if srcInfo.GetChainId() != dstInfo.GetChainId() {
		return nil, errorn.NewChainIDNotEqual(srcInfo.GetChainId(), dstInfo.GetChainId())
	}

	var meta lexer.ContractInvokeMeta
	meta.Code = invokeIntent.Code
	meta.FuncName = invokeIntent.FuncName
	meta.Meta, err = invokeIntent.Meta.RawBytes()
	if err != nil {
		return nil, err
	}

	meta.Params = make([]token_types.Param, len(invokeIntent.Params))
	for i := range invokeIntent.Params {
		meta.Params[i], err = invokeIntent.Params[i].Determine(ier)
		if err != nil {
			return nil, err
		}
	}
	b, err := lexer.MarshalContractInvokeMeta(&meta)
	if err != nil {
		return nil, err
	}

	intent = newIntent(&instruction.TransactionIntent{
		TransType: trans_type.ContractInvoke,
		Src:       srcInfo.GetAddress(),
		Dst:       dstInfo.GetAddress(),
		Meta:      b,
		Amt:       invokeIntent.Amount,
		ChainID:   srcInfo.GetChainId(),
	}, invokeIntent.GetName())
	proposals, err := ier.parseContractInvokeProof(&meta)
	if err != nil {
		return
	}
	intent.SetProposals(proposals)

	translator, err := ier.ChainGetter.GetTranslator(srcInfo.GetChainId())
	if err != nil {
		return nil, errorn.NewGetTranslatorError(err).Desc(errorn.AtChainID{srcInfo.GetChainId()})
	}

	intent, err = translator.ParseTransactionIntent(intent)
	if err != nil {
		return nil, errorn.NewParseTransactionIntentError(err)
	}

	intents = append(intents, intent)
	return
}

func (ier *Parser) parseContractInvokeProof(meta *lexer.ContractInvokeMeta) (proposals uip.MerkleProofProposalsImpl, err error) {
	//var b []byte
	//var txp transactionProofSourceDescription
	//txp.ChainID = intent.Src.ChainId
	//b, err = ier.marshal(&txp)
	//if err != nil {
	//	return
	//}
	//var mpt merkle_proof.Type
	//mpt, err = ier.accountBase.GetTransactionProofType(intent.Src.ChainId)
	//if err != nil {
	//	return nil, newGetTransactionProofType(err).Desc(AtChainID{intent.Src.ChainId})
	//}
	//
	//proposals = append(proposals, MerkleProofProposal{
	//	DescriptionType:   merkleproof_proposal_type.TransactionProof,
	//	MerkleProofType:   mpt,
	//	ValueType:         value_type.Bytes,
	//	SourceDescription: b,
	//})

	for _, param := range meta.Params {
		proposals, err = ier.addProposal(param, proposals)
		if err != nil {
			return
		}
	}
	return
}

//
//type StateVariable struct {
//	Type     value_type.Type `json:"type"`
//	Contract Account `json:"contract"`
//	Pos      []byte `json:"pos"`
//	Field    []byte `json:"field"`
//}
func (ier *Parser) addProposal(param token_types.Param, proposal uip.MerkleProofProposalsImpl) (uip.MerkleProofProposalsImpl, error) {

	switch param.GetGVMTok() {
	case token_type.Constant:
	case token_type.StateVariable:
		c := param.(token_types.StateVariableI).GetContract()
		if u, ok := c.(uip.Account); !ok {
			return nil, errorn.NewNoDeterminedAccount()
		} else {

			mpt, err := ier.AccountBase.GetTransactionProofType(u.GetChainId())
			if err != nil {
				return nil, errorn.NewGetTransactionProofType(err).Desc(errorn.AtChainID{ChainID: u.GetChainId()})
			}
			var b = bytes.NewBuffer(nil)
			lexer_types.EncodeVTok(b, param, &err)
			if err != nil {
				return nil, err
			}
			proposal = append(proposal, MerkleProofProposal{
				DescriptionType:   merkleproof_proposal_type.DataProof,
				MerkleProofType:   mpt,
				ValueType:         uip.TypeIDUnderlyingType(param.GetGVMType()),
				SourceDescription: b.Bytes(),
			})
		}
	}

	return proposal, nil
}
