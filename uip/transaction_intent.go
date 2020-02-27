package uip

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
)

type hexstring = string

type NamedIntent struct {
	Name string `json:"-"`
}

func (n NamedIntent) GetName() string {
	return n.Name
}

type TransactionIntent struct {

	TransType uint64          `json:"trans_type"`
	Src       []byte          `json:"src"`
	Dst       []byte          `json:"dst"`
	Meta      json.RawMessage `json:"meta"`
	Amt       hexstring       `json:"amt"`
	ChainID   uint64          `json:"chain_id"`
}

func (tx *TransactionIntent) GetType() instruction_type.Type {
	return tx.TransType
}

//
//Goto
//ConditionGoto
//SetState
//ConditionSetState

type Goto struct {
	Index int `json:"goto"`
}

func (tx *Goto) GetType() instruction_type.Type {
	return instruction_type.Goto
}

type ConditionGoto struct {
	Index     int             `json:"goto"`
	Condition json.RawMessage `json:"condition"`
}

func (tx *ConditionGoto) GetType() instruction_type.Type {
	return instruction_type.ConditionGoto
}

type SetState struct {
	Target          []byte          `json:"target"`
	RightExpression json.RawMessage `json:"expression"`
}

func (tx *SetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}

type ConditionSetState struct {
	Target          []byte          `json:"target"`
	RightExpression json.RawMessage `json:"expression"`
	Condition       json.RawMessage `json:"condition"`
}

func (tx *ConditionSetState) GetType() instruction_type.Type {
	return instruction_type.ConditionSetState
}

func (tx *TransactionIntent) Bytes() []byte {
	b, err := json.Marshal(tx)
	if err != nil {
		panic(err)
	}
	return b
}

type MerkleProofProposal struct {
	Tid uint64
	// 0 = transaction proof, 1 = data proof
	DescriptionType MerkleProofProposalType
	// 0 = simpleMerkleTreeUsingSha256, 1 = simpleMerkleTreeUsingSha512, 2 = merklePatriciaTrieUsingKeccak256
	MerkleProofType   MerkleProofType
	ValueType         TypeID
	SourceDescription []byte
}

type Instruction interface {
	GetType() instruction_type.Type
}

type TxIntentI interface {
	GetName() string

	GetInstruction() Instruction
	SetInstruction(Instruction)

	GetProposals() []MerkleProofProposal
	SetProposals([]MerkleProofProposal)
}
