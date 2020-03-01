package parser

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type hexstring = string

type NamedIntent struct {
	Name string `json:"-"`
}

func (n NamedIntent) GetName() string {
	return n.Name
}

type TransactionIntent struct {
	TransType trans_type.Type `json:"itype"`
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
	Type  instruction_type.Type `json:"itype"`
	Index int                   `json:"goto"`
}

func NewGoto(index int) *Goto {
	return &Goto{
		Type:  instruction_type.Goto,
		Index: index,
	}
}

func (tx *Goto) GetType() instruction_type.Type {
	return instruction_type.Goto
}

type ConditionGoto struct {
	Type      instruction_type.Type `json:"itype"`
	Index     int                   `json:"goto"`
	Condition json.RawMessage       `json:"condition"`
}

func NewConditionGoto(index int, condition json.RawMessage) *ConditionGoto {
	return &ConditionGoto{
		Type:      instruction_type.ConditionGoto,
		Index:     index,
		Condition: condition,
	}
}

func (tx *ConditionGoto) GetType() instruction_type.Type {
	return instruction_type.ConditionGoto
}

type SetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          []byte                `json:"target"`
	RightExpression json.RawMessage       `json:"expression"`
}

func NewSetState(t value_type.Type, target []byte, rhs json.RawMessage) *SetState {
	return &SetState{
		IType:     instruction_type.SetState,
		Type:     t,
		Target:     target,
		RightExpression: rhs,
	}
}

func (tx *SetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}

type ConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          []byte                `json:"target"`
	RightExpression json.RawMessage `json:"expression"`
	Condition       json.RawMessage `json:"condition"`
}

func NewConditionSetState(t value_type.Type, target []byte, rhs, cond json.RawMessage) *ConditionSetState {
	return &ConditionSetState{
		IType:     instruction_type.ConditionSetState,
		Type:     t,
		Target:     target,
		RightExpression: rhs,
		Condition: cond,
	}
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
