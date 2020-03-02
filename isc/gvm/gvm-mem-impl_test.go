package gvm_test

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/instruction_type"
	"github.com/HyperService-Consortium/go-uip/const/sign_type"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/isc/gvm"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func runMemoryGVM(callback func(g *gvm.ImplX), instructions []gvm.Instruction) {
	g := sugar.HandlerError(gvm.NewGVM()).(*gvm.ImplX)
	sugar.HandlerError0(g.AddFunction("main", instructions))
	var pc uint64
	var err error
	for pc, err = g.ConsumePC("main", pc); err == nil;  {
		pc, err = g.ConsumePC("main", pc)
		time.Sleep(time.Second)
	}

	_ = pc
	callback(g)
}

func setStateTestCase() []gvm.Instruction {
	return []gvm.Instruction{
		parser.GVMSetState{
			IType:           instruction_type.SetState,
			Type:            value_type.Bool,
			Target:          "a",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: true,
			},
		},
		parser.GVMSetState{
			IType:           instruction_type.SetState,
			Type:            value_type.Bool,
			Target:          "b",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: false,
			},
		},
		parser.GVMSetState{
			IType:           instruction_type.SetState,
			Type:            value_type.Bool,
			Target:          "c",
			RightExpression: lexer.BinaryExpression{
				Type:  value_type.Bool,
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
			IType:           instruction_type.SetState,
			Type:            value_type.Bool,
			Target:          "d",
			RightExpression: lexer.BinaryExpression{
				Type:  value_type.Bool,
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

func BenchmarkBase(b *testing.B) {
	for i := 0 ; i < b.N; i++ {
		runMemoryGVM(func(g *gvm.ImplX) {}, setStateTestCase())
	}
}

func BenchmarkPureBase(b *testing.B) {
	g := sugar.HandlerError(gvm.NewGVM()).(*gvm.ImplX)
	sugar.HandlerError0(g.AddFunction("main", setStateTestCase()))
	b.ResetTimer()
	for i := 0 ; i < b.N; i++ {
		var pc uint64
		var err error
		for pc, err = g.ConsumePC("main", pc); err == nil;  {
			pc, err = g.ConsumePC("main", pc)
			time.Sleep(time.Second)
		}
	}
}

func BenchmarkPureSetStatus(b *testing.B) {
	g := sugar.HandlerError(gvm.NewGVM()).(*gvm.ImplX)
	sugar.HandlerError0(g.AddFunction("main", []gvm.Instruction{
		parser.GVMSetState{
			IType:           instruction_type.SetState,
			Type:            value_type.Bool,
			Target:          "a",
			RightExpression: lexer.ConstantVariable{
				Type:  value_type.Bool,
				Const: true,
			},
		},
	}))
	b.ResetTimer()
	for i := 0 ; i < b.N; i++ {
		var pc uint64
		var err error
		for pc, err = g.ConsumePC("main", pc); err == nil;  {
			pc, err = g.ConsumePC("main", pc)
			time.Sleep(time.Second)
		}
	}
}

func TestBase(t *testing.T) {
	t.Run("set state", func(t *testing.T) {
		runMemoryGVM(func(g *gvm.ImplX) {
			fmt.Println(g.GVM.(*gvm.Mem).Context)
			assert.EqualValues(t, true, g.GVM.(*gvm.Mem).Context["a"].Unwrap())
			assert.EqualValues(t, false, g.GVM.(*gvm.Mem).Context["b"].Unwrap())
			assert.EqualValues(t, false, g.GVM.(*gvm.Mem).Context["c"].Unwrap())
			assert.EqualValues(t, true, g.GVM.(*gvm.Mem).Context["d"].Unwrap())
		}, setStateTestCase())
	})
	t.Run("branch condition", func(t *testing.T) {
		runMemoryGVM(func(g *gvm.ImplX) {

		}, nil)
	})
	t.Run("get local state", func(t *testing.T) {
		runMemoryGVM(func(g *gvm.ImplX) {

		}, nil)
	})
	t.Run("call function", func(t *testing.T) {
		runMemoryGVM(func(g *gvm.ImplX) {

		}, nil)
	})
}
