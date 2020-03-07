package abstraction

type TokType = uint64

type Tok interface {
	GetGVMTok() TokType
}

type VTok interface {
	Tok
	GetGVMType() RefType
	Eval(g *ExecCtx) (Ref, error)
}

//
