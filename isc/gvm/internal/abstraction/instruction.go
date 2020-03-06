package abstraction

type Instruction interface {
	Exec(g *ExecCtx) error
}
