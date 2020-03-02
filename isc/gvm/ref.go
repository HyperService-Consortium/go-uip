package gvm

type RefType = TokType

const (
	Unknown RefType = iota
	RefBytes
	RefString
	RefUint8
	RefUint16 // 4
	RefUint32
	RefUint64
	RefUint128
	RefUint256
	RefInt8 // 9
	RefInt16
	RefInt32
	RefInt64
	RefInt128
	RefInt256 // 14

	// Slice
	RefBool // 15
	Length
)

type Ref interface {
	GetGVMType() RefType
	Unwrap() interface{}
}


type ref struct {
	t RefType
	v interface{}
}

func (r ref) GetGVMType() RefType {
	return r.t
}

func (r ref) Unwrap() interface{} {
	return r.v
}
