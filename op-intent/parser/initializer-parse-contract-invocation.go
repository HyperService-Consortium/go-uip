package parser

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/merkleproof_proposal_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/HyperService-Consortium/go-uip/uip"
)




func (ier * Parser) parseContractInvocation(rawIntent lexer.Intent) (intents []uip.TxIntentI, err error) {
	invokeIntent := rawIntent.(*lexer.InvokeIntent)
	var srcInfo uip.Account
	var intent uip.TxIntentI
	srcInfo, err = ier.queryAccount(invokeIntent.Src)
	if err != nil {
		return nil, errorn.NewGetAccountFailed(err).Desc(errorn.AtOpIntentField{Field: "src"})
	}

	for i := range invokeIntent.Params {
		invokeIntent.Params[i], err = invokeIntent.Params[i].Determine(ier.queryContract)
		if err != nil {
			return nil, err
		}
	}

	var meta lexer.ContractInvokeMeta
	meta.Code = invokeIntent.Code
	meta.FuncName = invokeIntent.FuncName
	meta.Params = invokeIntent.Params
	meta.Meta, err = invokeIntent.Meta.RawBytes()
	if err != nil {
		return nil, err
	}
	b, err := ier.marshal(meta)
	if err != nil {
		return nil, err
	}

	var dstAddr []byte
	dstAddr, err = ier.decodeHex(invokeIntent.Dst)
	if err != nil {
		return
	}

	intent = new(TxIntentImpl)
	tx := &TransactionIntent{
		TransType: trans_type.ContractInvoke,
		Src:       srcInfo.GetAddress(),
		Dst:       dstAddr,
		Meta:      b,
		Amt:       invokeIntent.Amount,
		ChainID:   srcInfo.GetChainId(),
	}
	intent.SetIntent(tx)
	proposals, err := ier.parseContractInvokeProof(invokeIntent)
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

func DecodeContractAddress(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

func DecodeContractPos(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

func (ier * Parser) parseContractInvokeProof(intent *lexer.InvokeIntent) (proposals []MerkleProofProposal, err error) {
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

	for _, param := range intent.Params {
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
func (ier * Parser) addProposal(param lexer.Param, proposal []MerkleProofProposal) ([]MerkleProofProposal, error) {


	switch param.GetType() {
	case token.Constant:
	case token.StateVariable:
		c := param.(token.StateVariableI).GetContract()
		if u, ok := c.(uip.Account); !ok {
			return nil, errorn.NewNoDeterminedAccount()
		} else {

			mpt, err := ier.AccountBase.GetTransactionProofType(u.GetChainId())
			if err != nil {
				return nil, errorn.NewGetTransactionProofType(err).Desc(errorn.AtChainID{ChainID: u.GetChainId()})
			}
			v, err := ier.marshal(param)
			if err != nil {
				return nil, err
			}
			proposal = append(proposal, MerkleProofProposal{
				DescriptionType:   merkleproof_proposal_type.DataProof,
				MerkleProofType:   mpt,
				ValueType:         param.GetParamType(),
				SourceDescription: v,
			})
		}
	}


	return proposal, nil
}
