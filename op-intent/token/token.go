package token

type Type uint64

const (
	Pay Type = iota
	Invoke
	If
	Loop

	Constant
	StateVariable
	LocalStateVariable
	BinaryExpression
	UnaryExpression

	NamespacedNameAccount
	NameAccount
	NamespacedRawAccount
	RawAccount

	Length
)

func IsEvaluable(t Type) bool {
	return t <= UnaryExpression && t >= Constant
}
