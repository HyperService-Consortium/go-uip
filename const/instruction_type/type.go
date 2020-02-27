package instruction_type


type Type = uint64

const (
	Payment Type = iota
	ContractInvoke

	Goto
	ConditionGoto
	SetState
	ConditionSetState
)

const (
	RawGoto Type = iota + 100000
	RawConditionGoto
	RawSetState
	RawConditionSetState
)

