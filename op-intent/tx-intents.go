package opintent

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

//transactionIntents []*TransactionIntent, proposals []*MerkleProofProposal

type TxIntentImpl struct {
	Intent    *TransactionIntent
	proposals []MerkleProofProposal
}

func (t *TxIntentImpl) GetIntent() *uip.TransactionIntent {
	return t.Intent
}

func (t *TxIntentImpl) SetIntent(i *uip.TransactionIntent) {
	t.Intent = i
	return
}

func (t *TxIntentImpl) GetProposals() []uip.MerkleProofProposal {
	return t.proposals
}

func (t *TxIntentImpl) SetProposals(p []uip.MerkleProofProposal) {
	t.proposals = p
	return
}

type TxIntents interface {
	GetTxIntents() []uip.TxIntentI
	GetInformation() interface{}
}

type TxIntentsImpl []uip.TxIntentI

func (t TxIntentsImpl) Len() int {
	return len(t)
}

func (t TxIntentsImpl) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TxIntentsImpl) GetTxIntents() []uip.TxIntentI {
	return t
}

func (t TxIntentsImpl) GetInformation() interface{} {
	return nil
}
