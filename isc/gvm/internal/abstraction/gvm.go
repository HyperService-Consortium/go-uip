package abstraction

type Function interface {
	Fetch(pc uint64) (Instruction, error)
	Len() int
}

type Machine interface {
	CreateRef(t RefType, v interface{}) Ref
	DecodeRef(t RefType, b []byte) (Ref, error)

	GetFunction(function string) (Function, error)

	Load(field string, t RefType) (Ref, error)
	Save(field string, k Ref) error
	Delete(field string) error
}

type Locals = map[string]Ref

type ExecCtx struct {
	Machine
	Function
	Depth        uint64
	PC           uint64
	FN           string
	Parent, This Locals
}

func NewExecCtx() *ExecCtx {
	return &ExecCtx{
		Parent: make(Locals),
		This:   make(Locals),
	}
}
