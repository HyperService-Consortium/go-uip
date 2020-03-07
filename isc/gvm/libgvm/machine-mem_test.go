package libgvm_test

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/isc/gvm"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/libgvm"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type doInst struct {
	g func(g *abstraction.ExecCtx) error
}

func (d doInst) Exec(g *abstraction.ExecCtx) error {
	return d.g(g)
}

func runMemoryGVM(callback func(g *libgvm.GVMeX), instructions []abstraction.Instruction) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", instructions))
	sugar.HandlerError0(g.AddFunction("setA", funcSetA()))
	sugar.HandlerError0(g.AddFunction("fib", funcFib()))
	var err error
	for err = g.Run("main"); err == nil; {
		err = g.Run("main")
		time.Sleep(time.Second)
	}
	callback(g)
}

type BinaryExpression struct {
	Type  abstraction.RefType `json:"type"`
	Sign  sign_type.Type      `json:"sign"`
	Left  gvm.VTok            `json:"left"`
	Right gvm.VTok            `json:"right"`
}

func (b BinaryExpression) GetGVMTok() abstraction.TokType {
	return libgvm.TokBinaryExpression
}

func (b BinaryExpression) GetGVMType() abstraction.RefType {
	return b.Type
}

func (b BinaryExpression) Eval(g *abstraction.ExecCtx) (abstraction.Ref, error) {
	l, err := b.Left.Eval(g)
	if err != nil {
		return nil, err
	}
	r, err := b.Right.Eval(g)
	if err != nil {
		return nil, err
	}
	switch b.Sign {
	case sign_type.EQ:
		return libgvm.EQ(l, r)
	case sign_type.LE:
		return libgvm.LE(l, r)
	case sign_type.LT:
		return libgvm.LT(l, r)
	case sign_type.GE:
		return libgvm.GE(l, r)
	case sign_type.GT:
		return libgvm.GT(l, r)
	case sign_type.LAnd:
		return libgvm.LAnd(l, r)
	case sign_type.LOr:
		return libgvm.LOr(l, r)
	case sign_type.ADD:
		return libgvm.Add(l, r)
	case sign_type.SUB:
		return libgvm.Sub(l, r)
	case sign_type.MUL:
		return libgvm.Mul(l, r)
	case sign_type.QUO:
		return libgvm.Quo(l, r)
	case sign_type.REM:
		return libgvm.Rem(l, r)
	default:
		return nil, fmt.Errorf("unknown sign_type: %v", b.Sign)
	}
}

func setStateTestCase() []abstraction.Instruction {
	return []abstraction.Instruction{
		parser.GVMSetState{
			IType:  instruction_type.SetState,
			Type:   value_type.Bool,
			Target: "a",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: true,
			},
		},
		parser.GVMSetState{
			IType:  instruction_type.SetState,
			Type:   value_type.Bool,
			Target: "b",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: false,
			},
		},
		parser.GVMSetState{
			IType:  instruction_type.SetState,
			Type:   value_type.Bool,
			Target: "c",
			RightExpression: lexer.BinaryExpression{
				Type: value_type.Bool,
				Sign: sign_type.LAnd,
				Left: lexer.ConstantVariable{
					Type:  value_type.Bool,
					Const: false,
				},
				Right: lexer.ConstantVariable{
					Type:  value_type.Bool,
					Const: true,
				},
			},
		},
		parser.GVMSetState{
			IType:  instruction_type.SetState,
			Type:   value_type.Bool,
			Target: "d",
			RightExpression: lexer.BinaryExpression{
				Type: value_type.Bool,
				Sign: sign_type.LOr,
				Left: lexer.ConstantVariable{
					Type:  value_type.Bool,
					Const: false,
				},
				Right: lexer.ConstantVariable{
					Type:  value_type.Bool,
					Const: true,
				},
			},
		},
	}
}

func funcSetA() []abstraction.Instruction {
	return []abstraction.Instruction{
		parser.GVMSetState{
			IType:  instruction_type.SetState,
			Type:   value_type.Bool,
			Target: "a",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: true,
			},
		},
	}
}

func callSetBoolFuncTestCase() []abstraction.Instruction {
	return []abstraction.Instruction{
		libgvm.CallFunc{FN: "setA"},
	}
}

func BenchmarkBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runMemoryGVM(func(g *libgvm.GVMeX) {}, setStateTestCase())
	}
}

func BenchmarkPureBase(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", setStateTestCase()))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

func BenchmarkPureSetStatus(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		parser.GVMSetState{
			IType:  instruction_type.SetState,
			Type:   value_type.Bool,
			Target: "a",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: true,
			},
		},
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

type donothing struct {
}

func (d donothing) Exec(g *abstraction.ExecCtx) error {
	g.PC++
	return nil
}

func BenchmarkPureDoNothing(b *testing.B) {
	g := sugar.HandlerError(libgvm.NewGVM()).(*libgvm.GVMeX)
	sugar.HandlerError0(g.AddFunction("main", []abstraction.Instruction{
		donothing{},
	}))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.Run("main")
	}
}

func TestBase(t *testing.T) {
	t.Run("set state", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {
			fmt.Println(g.Machine.(*libgvm.Mem).Context)
			assert.EqualValues(t, true, g.Machine.(*libgvm.Mem).Context["a"].Unwrap())
			assert.EqualValues(t, false, g.Machine.(*libgvm.Mem).Context["b"].Unwrap())
			assert.EqualValues(t, false, g.Machine.(*libgvm.Mem).Context["c"].Unwrap())
			assert.EqualValues(t, true, g.Machine.(*libgvm.Mem).Context["d"].Unwrap())
		}, setStateTestCase())
	})
	t.Run("branch condition", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {

		}, nil)
	})
	t.Run("get local state", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {

		}, nil)
	})
	t.Run("call function", func(t *testing.T) {
		runMemoryGVM(func(g *libgvm.GVMeX) {
			assert.EqualValues(t, true, g.Machine.(*libgvm.Mem).Context["a"].Unwrap())
		}, callSetBoolFuncTestCase())
	})
}

func funcFib() []gvm.Instruction {
	// func fib(n int64) (r int64)
	return []gvm.Instruction{
		// q := 0
		libgvm.SetLocal{Target: "q", RightExpression: libgvm.Int64(0)},
		// if n > 0 { q = fib(n - 1); }
		libgvm.ConditionCallFunc{
			CallFunc: libgvm.CallFunc{
				FN: "fib", Left: []string{"q"}, Right: []gvm.VTok{BinaryExpression{
					Type: libgvm.RefInt64, Sign: sign_type.SUB, Left: libgvm.FuncParam{T: libgvm.RefInt64, K: 0}, Right: libgvm.Int64(1),
				}},
			},
			Condition: BinaryExpression{
				Type: libgvm.RefBool, Sign: sign_type.GT, Left: libgvm.FuncParam{T: libgvm.RefInt64, K: 0}, Right: libgvm.Int64(0),
			},
		},
		// r = n + q; return r
		libgvm.SetFuncReturn{Target: 0, RightExpression: BinaryExpression{
			Type: libgvm.RefInt64, Sign: sign_type.ADD, Left: libgvm.LocalVariable{Name: "q"}, Right: libgvm.FuncParam{T: libgvm.RefInt64, K: 0},
		}},
	}
}

func TestFibonacci(t *testing.T) {
	runMemoryGVM(func(g *libgvm.GVMeX) {
		//fmt.Println(g.Machine.(*libgvm.Mem).Context)
	}, []abstraction.Instruction{
		libgvm.CallFunc{FN: "fib", Left: []string{"res"}, Right: []abstraction.VTok{libgvm.Int64(3)}},
		doInst{g: func(g *abstraction.ExecCtx) error {
			fmt.Println("fib(3) =", g.This["res"])
			g.PC++
			return nil
		}},
	})
}

func BenchmarkFibnacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		runMemoryGVM(func(g *libgvm.GVMeX) {
		}, []abstraction.Instruction{
			libgvm.CallFunc{FN: "fib", Left: []string{"res"}, Right: []abstraction.VTok{
				libgvm.Int64(1)}},
		})
	}
}
