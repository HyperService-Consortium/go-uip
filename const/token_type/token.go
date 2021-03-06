package token_type

type Type = uint64

//noinspection ALL
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

	FuncParam
	ReturnParam
	LocalVariable

	NamespacedNameAccount
	NameAccount
	NamespacedRawAccount
	RawAccount
	FullAccount

	Length
)

//noinspection GoUnusedExportedFunction
func IsEvaluable(t Type) bool {
	return t <= LocalVariable && t >= Constant
}
