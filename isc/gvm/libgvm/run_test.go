package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

type gotoImpl struct{}

func (g gotoImpl) Exec(c *abstraction.ExecCtx) error {
	c.PC = uint64(g.GetGotoIndexGVMI())
	return nil
}

func (g gotoImpl) GetGotoIndexGVMI() uint64 {
	return 2
}

func Test__Goto_Exec(t *testing.T) {
	type args struct {
		g *abstraction.ExecCtx
		i abstraction.Instruction
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"easy", args{&abstraction.ExecCtx{
			Machine:  nil,
			Function: nil,
			PC:       1,
		}, gotoImpl{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.i.Exec(tt.args.g); (err != nil) != tt.wantErr {
				t.Fatalf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				return
			}

			if uint64(tt.args.i.(gotoImpl).GetGotoIndexGVMI()) != tt.args.g.PC {
				t.Errorf("Exec() get pc = %v, want: %v", tt.args.g.PC, tt.args.i.(gotoImpl).GetGotoIndexGVMI())
			}
		})
	}
}

type donothing struct {
}

func (d donothing) Exec(g *abstraction.ExecCtx) error {
	g.PC++
	return nil
}

func BenchmarkContinue(b *testing.B) {
	g := sugar.HandlerError(NewGVM()).(*GVMeX)
	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		donothing{},
	}))
	var c = &abstraction.ExecCtx{Machine: g, Depth: 0, This: make(abstraction.Locals)}
	sugar.HandlerError0(pushFrame(c, "main"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c0 := *c
		_ = _Continue(&c0)
	}
}
