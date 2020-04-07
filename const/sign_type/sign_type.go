package sign_type

type Type = uint32

//noinspection GoUnusedConst
const (
	Unknown Type = iota

	EQ
	NEQ
	LE
	LT
	GE
	GT
	LAnd
	LOr
	LNot

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	//ADD_ASSIGN // +=
	//SUB_ASSIGN // -=
	//MUL_ASSIGN // *=
	//QUO_ASSIGN // /=
	//REM_ASSIGN // %=

	//AND_ASSIGN     // &=
	//OR_ASSIGN      // |=
	//XOR_ASSIGN     // ^=
	//SHL_ASSIGN     // <<=
	//SHR_ASSIGN     // >>=
	//AND_NOT_ASSIGN // &^=

	LogicL = EQ
	LogicR = LOr + 1
)

func IsLogic(s Type) bool {
	return LogicL <= s && s < LogicR
}
