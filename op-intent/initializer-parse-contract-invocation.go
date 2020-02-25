package opintent

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/merkleproof_proposal_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/uip"
	"strconv"
)


func (ier *Initializer) parseContractInvocation(rawIntent lexer.Intent) (intents []uip.TxIntentI, err error) {
	invokeIntent := rawIntent.(*lexer.InvokeIntent)
	var srcInfo uip.Account
	var intent uip.TxIntentI
	srcInfo, err = ier.accountBase.Get(
		invokeIntent.Src.Name, invokeIntent.Src.ChainId)
	if err != nil {
		return nil, errorn.NewGetAccountFailed(err).Desc(errorn.AtOpIntentField{"src"})
	}

	var meta lexer.ContractInvokeMeta
	meta.Code = invokeIntent.Code
	meta.FuncName = invokeIntent.FuncName
	meta.Params = invokeIntent.Params
	meta.Meta, err = invokeIntent.Meta.RawBytes()
	if err != nil {
		return nil, err
	}
	b, err := ier.marshal(&meta)
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

	translator, err := ier.chainGetter.GetTranslator(srcInfo.GetChainId())
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

func (ier *Initializer) parseContractInvokeProof(intent *lexer.InvokeIntent) (proposals []MerkleProofProposal, err error) {
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

func (ier *Initializer) addProposal(param lexer.ParamImpl, proposal []MerkleProofProposal) ([]MerkleProofProposal, error) {
	var intDesc value_type.Type
	if intDesc = value_type.FromString(param.Type); intDesc == value_type.Unknown {
		return nil, errorn.NewValueTypeNotFound(param.Type)
	}

	result := param.Value
	if !result.Get("constant").Exists() {
		if result.Get("contract").Exists() &&
			result.Get("pos").Exists() &&
			result.Get("field").Exists() {
			ca, err := DecodeContractAddress(result.Get("contract").String())
			if err != nil {
				return nil, errorn.NewDecodeAddressError(err)
			}
			pos, err := DecodeContractPos(result.Get("pos").String())
			if err != nil {
				return nil, errorn.NewDecodeContractPosError(err)
			}
			var domain uint64
			if result.Get("domain").Exists() {
				domain, err = strconv.ParseUint(result.Get("domain").String(), 10, 64)
				if err != nil {
					return nil, errorn.NewDecodeDomainError(err)
				}
			} else {
				domain, err = ier.contractBase.GetChainID(ca)
				if err != nil {
					return nil, errorn.NewGetDomainError(err)
				}
			}

			mpt, err := ier.accountBase.GetTransactionProofType(domain)
			if err != nil {
				return nil, errorn.NewGetTransactionProofType(err).Desc(errorn.AtChainID{domain})
			}

			_, _ = ca, pos

			v, err := param.Value.RawBytes()
			if err != nil {
				return nil, err
			}
			proposal = append(proposal, MerkleProofProposal{
				DescriptionType:   merkleproof_proposal_type.DataProof,
				MerkleProofType:   mpt,
				ValueType:         intDesc,
				SourceDescription: v,
			})
		} else {
			return nil, errorn.NewNotEnoughParamInformation()
		}
	}
	return proposal, nil
}
