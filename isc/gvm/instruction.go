package gvm

type InstructionType = TokType

const (
	InstGoto = iota
	InstConditionGoto
	InstSetState
	InstConditionSetState
)

type Instruction interface {
	GetGVMIType() InstructionType
}

type LazyInstruction interface {
	DeserializeGVMI() (Instruction, error)
	Instruction
}

type InstPrototype interface {

	Exec(g *ExecCtx, i Instruction) error
}