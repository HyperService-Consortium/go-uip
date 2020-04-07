package instruction_type

type Type = uint64

const (
	Payment Type = iota
	ContractInvoke

	Goto
	ConditionGoto
	SetState
	ConditionSetState

	Length
	StandardStep = 1 << 16
	Standard     = StandardStep << 4
)

const (
	RawGoto Type = iota + StandardStep
	RawConditionGoto
	RawSetState
	RawConditionSetState

	RawLengthL = RawGoto
	RawLengthR = RawConditionSetState + 1
)

func IsStandard(t Type) bool {
	return t < Standard
}

func IsValidType(t Type) bool {
	return t < Length
}

func IsRaw(t Type) bool {
	return t < RawLengthR && RawLengthL <= t
}
