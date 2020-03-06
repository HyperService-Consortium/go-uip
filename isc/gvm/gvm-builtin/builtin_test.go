package gvm_builtin

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		l interface{}
		r interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "int", args: args{1, 2}, want: 3},
		{name: "int64", args: args{int64(1), int64(2)}, want: int64(3)},
		{name: "uint64", args: args{uint64(1), uint64(2)}, want: uint64(3)},
		{name: "bigint", args: args{big.NewInt(1), big.NewInt(2)}, want: big.NewInt(3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.l, tt.args.r); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuo(t *testing.T) {
	type args struct {
		l interface{}
		r interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "int", args: args{1, 2}, want: 0},
		{name: "int64", args: args{int64(1), int64(2)}, want: int64(0)},
		{name: "uint64", args: args{uint64(1), uint64(2)}, want: uint64(0)},
		{name: "bigint", args: args{big.NewInt(1), big.NewInt(2)}, want: big.NewInt(0)},
		{name: "bigint", args: args{big.NewInt(2), big.NewInt(2)}, want: big.NewInt(1)},
		{name: "bigint", args: args{big.NewInt(2), big.NewInt(-2)}, want: big.NewInt(-1)},
		{name: "bigint", args: args{big.NewInt(3), big.NewInt(-2)}, want: big.NewInt(-1)},
		{name: "bigint", args: args{big.NewInt(-3), big.NewInt(-2)}, want: big.NewInt(1)},
		{name: "bigint", args: args{big.NewInt(-3), big.NewInt(2)}, want: big.NewInt(-1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quo(tt.args.l, tt.args.r); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
