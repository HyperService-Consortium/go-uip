package parser

import (
	"encoding/json"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/isc/gvm"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/libgvm"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
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
	Index uint64                `json:"goto"`
}

func (tx *Goto) GetGotoIndexGVMI() uint64 {
	return tx.Index
}

func NewGoto(index uint64) *Goto {
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
	Index     uint64                `json:"goto"`
	Condition json.RawMessage       `json:"condition"`
}

func NewConditionGoto(index uint64, condition json.RawMessage) *ConditionGoto {
	return &ConditionGoto{
		Type:      instruction_type.ConditionGoto,
		Index:     index,
		Condition: condition,
	}
}

func (tx *ConditionGoto) Convert() (g *GVMConditionGoto, err error) {
	g = &GVMConditionGoto{
		Type:  tx.Type,
		Index: tx.Index,
	}
	g.Condition, err = lexer.ParamUnmarshalJSON(tx.Condition)
	return
}

type GVMConditionGoto struct {
	Type      instruction_type.Type `json:"itype"`
	Index     uint64                `json:"goto"`
	Condition gvm.VTok              `json:"condition"`
}

func (inst *GVMConditionGoto) Exec(g *gvm.ExecCtx) error {
	v, err := inst.GetGotoConditionGVMI().Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != libgvm.RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}
	if v.Unwrap().(bool) {
		g.PC = uint64(inst.GetGotoIndexGVMI())
		return nil
	}
	g.PC++
	return nil
}

func (inst *GVMConditionGoto) GetGotoIndexGVMI() uint64 {
	return inst.Index
}

func (inst *GVMConditionGoto) GetGotoConditionGVMI() gvm.VTok {
	return inst.Condition
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
		IType:           instruction_type.SetState,
		Type:            t,
		Target:          target,
		RightExpression: rhs,
	}
}

type GVMSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          string                `json:"target"`
	RightExpression gvm.VTok              `json:"expression"`
}

func (G GVMSetState) Exec(g *gvm.ExecCtx) error {
	k, err := G.GetRightHandStatementGVMI().Eval(g)
	if err != nil {
		return err
	}
	err = g.Save(G.GetRefNameGVMI(), k)
	if err != nil {
		return err
	}

	g.PC++
	return nil
}

func (G GVMSetState) GetRefNameGVMI() string {
	return G.Target
}

func (G GVMSetState) GetRightHandStatementGVMI() gvm.VTok {
	return G.RightExpression
}

func (tx *SetState) Convert() (g *GVMSetState, err error) {
	g = &GVMSetState{
		IType:  tx.IType,
		Type:   tx.Type,
		Target: string(tx.Target),
	}
	g.RightExpression, err = lexer.ParamUnmarshalJSON(tx.RightExpression)
	return
}

func (tx *SetState) GetType() instruction_type.Type {
	return instruction_type.SetState
}

type ConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          []byte                `json:"target"`
	RightExpression json.RawMessage       `json:"expression"`
	Condition       json.RawMessage       `json:"condition"`
}

type GVMConditionSetState struct {
	IType           instruction_type.Type `json:"itype"`
	Type            value_type.Type       `json:"value_type"`
	Target          string                `json:"target"`
	RightExpression gvm.VTok              `json:"expression"`
	Condition       gvm.VTok              `json:"condition"`
}

func (tx *ConditionSetState) Convert() (g *GVMConditionSetState, err error) {
	g = &GVMConditionSetState{
		IType:  tx.IType,
		Type:   tx.Type,
		Target: string(tx.Target),
	}
	g.RightExpression, err = lexer.ParamUnmarshalJSON(tx.RightExpression)
	if err != nil {
		return nil, err
	}
	g.Condition, err = lexer.ParamUnmarshalJSON(tx.Condition)
	return
}

func (inst *GVMConditionSetState) Exec(g *gvm.ExecCtx) error {
	v, err := inst.GetGotoConditionGVMI().Eval(g)
	if err != nil {
		return err
	}
	if v.GetGVMType() != libgvm.RefBool {
		return fmt.Errorf("type error: not bool value, is %v", v.GetGVMType())
	}

	if v.Unwrap().(bool) {
		k, err := inst.GetRightHandStatementGVMI().Eval(g)
		if err != nil {
			return err
		}
		err = g.Save(inst.GetRefNameGVMI(), k)
		if err != nil {
			return err
		}
	}

	g.PC++
	return nil
}

func (G GVMConditionSetState) GetGotoConditionGVMI() gvm.VTok {
	return G.Condition
}

func (G GVMConditionSetState) GetRefNameGVMI() string {
	return G.Target
}

func (G GVMConditionSetState) GetRightHandStatementGVMI() gvm.VTok {
	return G.RightExpression
}

func NewConditionSetState(t value_type.Type, target []byte, rhs, cond json.RawMessage) *ConditionSetState {
	return &ConditionSetState{
		IType:           instruction_type.ConditionSetState,
		Type:            t,
		Target:          target,
		RightExpression: rhs,
		Condition:       cond,
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
