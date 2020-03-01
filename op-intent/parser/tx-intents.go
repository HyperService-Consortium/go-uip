package parser

import "github.com/HyperService-Consortium/go-uip/uip"

//transactionIntents []*TransactionIntent, proposals []*MerkleProofProposal

type TxIntentImpl struct {
	NamedIntent
	Instruction uip.Instruction
	proposals   uip.MerkleProofProposals
}

func newIntent(instruction uip.Instruction, name string) (impl *TxIntentImpl) {
	impl = new(TxIntentImpl)
	impl.Instruction = instruction
	impl.Name = name
	return impl
}

func (t *TxIntentImpl) GetInstruction() uip.Instruction {
	return t.Instruction
}

func (t *TxIntentImpl) SetInstruction(i uip.Instruction) {
	t.Instruction = i
	return
}

func (t *TxIntentImpl) GetProposals() uip.MerkleProofProposals {
	return t.proposals
}

func (t *TxIntentImpl) SetProposals(p uip.MerkleProofProposals) {
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
