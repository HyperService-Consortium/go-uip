package uip

import (
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/Myriad-Dreamin/gvm"
	"io"
)

type Serializable interface {
	Marshal(w io.Writer, err *error)
}

type Instruction interface {
	Serializable
	gvm.Instruction
	GetType() instruction_type.Type

	Unmarshal(r io.Reader, i *Instruction, err *error)
}

type VTok interface {
	Serializable
	gvm.VTok

	Unmarshal(r io.Reader, i *VTok, err *error)
}

//todo
type TransactionIntent interface {
	GetTxType() trans_type.Type
	GetChainID() ChainIDUnderlyingType
}

type MerkleProofProposal interface {
	GetMerkleProofProposalType() MerkleProofProposalType
	GetMerkleProofType() MerkleProofType
	GetValueType() TypeID
	GetSourceDescription() []byte
}

type BaseSlice interface {
	Len() int
	Cap() int
}

type MerkleProofProposals interface {
	BaseSlice
	Index(i int) MerkleProofProposal
	Slice(l, r int) MerkleProofProposals
	Append(appends ...MerkleProofProposal) (T MerkleProofProposals)
}

type MerkleProofProposalsImpl []MerkleProofProposal

func (m MerkleProofProposalsImpl) Len() int {
	return len(m)
}

func (m MerkleProofProposalsImpl) Cap() int {
	return cap(m)
}

func (m MerkleProofProposalsImpl) Index(i int) MerkleProofProposal {
	return m[i]
}

func (m MerkleProofProposalsImpl) Slice(l, r int) MerkleProofProposals {
	return m[l:r]
}

func (m MerkleProofProposalsImpl) Append(appends ...MerkleProofProposal) (T MerkleProofProposals) {
	return append(m, appends...)
}

type TxIntentI interface {
	GetName() string

	GetInstruction() Instruction
	SetInstruction(Instruction)

	GetProposals() MerkleProofProposals
	SetProposals(MerkleProofProposals)
}
