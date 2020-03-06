package abstraction

type RefType = TokType

type Ref interface {
	GetGVMType() RefType
	Unwrap() interface{}
	Encode() ([]byte, error)
}
