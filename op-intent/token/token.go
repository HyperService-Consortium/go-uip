package token

type Type uint64

const (
	Pay Type = iota
	Invoke
	If
	Loop

	Constant
	StateVariable

	NamespacedNameAccount
	NameAccount
	NamespacedRawAccount
	RawAccount
)

