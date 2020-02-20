package opintent

import (
	"encoding/hex"
	"github.com/HyperService-Consortium/go-uip/const/merkleproof_proposal_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/tidwall/gjson"
	"strconv"
)

func (ier *Initializer) parseContractInvocation(rawIntent RawIntentI) (intents []uip.TxIntentI, err error) {
	invokeIntent := rawIntent.GetSub().(*BaseContractInvocationOpIntent)
	var srcInfo uip.Account
	var intent uip.TxIntentI
	srcInfo, err = ier.accountBase.Get(
		invokeIntent.Src.Name, invokeIntent.Src.ChainId)
	if err != nil {
		return nil, newGetAccountFailed(err).Desc(AtOpIntentField{"src"})
	}

	var meta uip.ContractInvokeMeta
	meta.Code = invokeIntent.Code
	meta.FuncName = invokeIntent.FuncName
	meta.Params = invokeIntent.Params
	meta.Meta = invokeIntent.Meta
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
		return nil, newGetTranslatorError(err).Desc(AtChainID{srcInfo.GetChainId()})
	}

	intent, err = translator.ParseTransactionIntent(intent)
	if err != nil {
		return nil, newParseTransactionIntentError(err)
	}

	return
}

func DecodeContractAddress(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

func DecodeContractPos(src string) ([]byte, error) {
	return hex.DecodeString(src)
}

func (ier *Initializer) parseContractInvokeProof(intent *BaseContractInvocationOpIntent) (proposals []MerkleProofProposal, err error) {
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

func (ier *Initializer) addProposal(param uip.RawParam, proposal []MerkleProofProposal) ([]MerkleProofProposal, error) {
	var intDesc value_type.Type
	if intDesc = value_type.FromString(param.Type); intDesc == value_type.Unknown {
		return nil, newValueTypeNotFound(param.Type)
	}

	result := gjson.ParseBytes(param.Value)
	if !result.Get("constant").Exists() {
		if result.Get("contract").Exists() &&
			result.Get("pos").Exists() &&
			result.Get("field").Exists() {
			ca, err := DecodeContractAddress(result.Get("contract").String())
			if err != nil {
				return nil, newDecodeContractAddressError(err)
			}
			pos, err := DecodeContractPos(result.Get("pos").String())
			if err != nil {
				return nil, newDecodeContractPosError(err)
			}
			var domain uint64
			if result.Get("domain").Exists() {
				domain, err = strconv.ParseUint(result.Get("domain").String(), 10, 64)
				if err != nil {
					return nil, newDecodeDomainError(err)
				}
			} else {
				domain, err = ier.contractBase.GetChainID(ca)
				if err != nil {
					return nil, newGetDomainError(err)
				}
			}

			mpt, err := ier.accountBase.GetTransactionProofType(domain)
			if err != nil {
				return nil, newGetTransactionProofType(err).Desc(AtChainID{domain})
			}

			_, _ = ca, pos

			proposal = append(proposal, MerkleProofProposal{
				DescriptionType:   merkleproof_proposal_type.DataProof,
				MerkleProofType:   mpt,
				ValueType:         intDesc,
				SourceDescription: param.Value,
			})
		} else {
			return nil, newNotEnoughParamInformation()
		}
	}
	return proposal, nil
}
