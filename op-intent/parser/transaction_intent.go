package parser

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

//
//Goto
//ConditionGoto
//SetState
//ConditionSetState

type MerkleProofProposal struct {
	// 0 = transaction proof, 1 = data proof
	DescriptionType uip.MerkleProofProposalType
	// 0 = simpleMerkleTreeUsingSha256, 1 = simpleMerkleTreeUsingSha512, 2 = merklePatriciaTrieUsingKeccak256
	MerkleProofType   uip.MerkleProofType
	ValueType         uip.TypeID
	SourceDescription []byte
}

func (m MerkleProofProposal) GetMerkleProofProposalType() uip.MerkleProofProposalType {
	return m.DescriptionType
}

func (m MerkleProofProposal) GetMerkleProofType() uip.MerkleProofType {
	return m.MerkleProofType
}

func (m MerkleProofProposal) GetValueType() uip.TypeID {
	return m.ValueType
}

func (m MerkleProofProposal) GetSourceDescription() []byte {
	return m.SourceDescription
}
