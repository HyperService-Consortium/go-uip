package abstraction

type VTok interface {
	Tok
	GetGVMType() RefType
	Eval(g GVM) (Ref, error)
}

type Function interface {
	Fetch(pc uint64) (Instruction, error)
	Len() int
}

type GVM interface {
	CreateRef(t RefType, v interface{}) Ref
	DecodeRef(t RefType, b []byte) (Ref, error)

	//GetPackPrototype(t RefType) (libgvm.PackPrototype, error)
	GetFunction(function string) (Function, error)

	Load(field string, t RefType) (Ref, error)
	Save(field string, k Ref) error
}

type ExecCtx struct {
	GVM
	Function
	PC uint64
}
