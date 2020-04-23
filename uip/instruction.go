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

type InstructionType = instruction_type.Type

type Instruction interface {
	Serializable
	gvm.Instruction
	GetType() InstructionType

	Unmarshal(r io.Reader, i *Instruction, err *error)
}

type VTok interface {
	Serializable
	gvm.VTok

	Unmarshal(r io.Reader, i *VTok, err *error)
}

type TransactionIntent interface {
	GetTxType() trans_type.Type
	GetChainID() ChainIDUnderlyingType
	GetSrc() []byte
	GetDst() []byte
	GetMeta() []byte
	GetAmt() string
}

type MerkleProofProposal interface {
	GetMerkleProofProposalType() MerkleProofProposalType
	GetMerkleProofType() MerkleProofType
	GetValueType() TypeID
	GetSourceDescription() []byte
}

type MerkleProofProposals interface {
	BaseSlice
	Index(i int) MerkleProofProposal
	Slice(l, r int) MerkleProofProposals
	Append(appends ...MerkleProofProposal) (T MerkleProofProposals)
}

type TxIntentI interface {
	GetName() string

	GetInstruction() Instruction
	SetInstruction(Instruction)

	GetProposals() MerkleProofProposals
	SetProposals(MerkleProofProposals)
}
