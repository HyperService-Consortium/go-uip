package gvm


type Impl struct {
	GVM
}

func (i *Impl) Eval(p VTok) (Ref, error) {
	return EvalG(i.GVM, p)
}

func (i *Impl) Decode(x []byte, t RefType) (Ref, error) {
	return DecodeG(i.GVM, x, t)
}

func (i *Impl) Encode(v Ref) ([]byte, error) {
	return EncodeG(i.GVM, v)
}

func (i *Impl) ConsumePC(fn string, pc uint64) (uint64, error) {
	return ConsumePC(i.GVM, fn, pc)
}

func (i *Impl) InitPC(fn string, pc uint64) (uint64, error) {
	return InitPC(i.GVM, fn, pc)
}

func (i *Impl) NextPC(fn string, pc uint64) (uint64, error) {
	return NextPC(i.GVM, fn, pc)
}

type impl = Impl
type ImplX struct {
	impl
	g *Mem
}

func (i *ImplX) AddFunction(fn string, instructions []Instruction) error {
	if i.g.InstSet == nil {
		i.g.InstSet = make(map[string]Function)
	}

	i.g.InstSet[fn] = FunctionImpl(instructions)
	return nil
}

type VTok interface {
	Tok
	GetGVMType() RefType
}

type Function interface {
	LazyFetch(pc uint64) (LazyInstruction, error)
	Fetch(pc uint64) (Instruction, error)
	Len() int
}

type GVM interface {
	CreateRef(t RefType, v interface{}) Ref

	GetPackPrototype(t RefType) (PackPrototype, error)
	GetTypePrototype(t TokType) (TypePrototype, error)
	GetInstPrototype(t InstructionType) (InstPrototype, error)

	GetFunction(function string) (Function, error)

	Load(field string, t RefType) (Ref, error)
	Save(field string, k Ref) error
}

type TypePrototype interface {
	Eval(g GVM, p VTok) (Ref, error)
}

type Base struct{}

var IObjs = [...]InstPrototype {
	InstGoto: (*_Goto)(nil),
	InstConditionGoto: (*_ConditionGoto)(nil),
	InstSetState: (*_SetState)(nil),
	InstConditionSetState: (*_ConditionSetState)(nil),
}

func (b *Base) GetInstPrototype(t InstructionType) (InstPrototype, error) {
	return IObjs[t], nil
}

func (b *Base) GetPackPrototype(t RefType) (PackPrototype, error) {
	return PObjs[t], nil
}

func (b *Base) GetTypePrototype(t TokType) (TypePrototype, error) {
	return TObjs[t], nil
}

func (*Base) CreateRef(t RefType, v interface{}) Ref {
	return &ref{t, v}
}