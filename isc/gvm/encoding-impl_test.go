package gvm

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestGoodEDecode(t *testing.T) {
	type args struct {
		variable Ref
	}
	tests := []struct {
		name     string
		variable Ref
	}{
		{"bool", ref{
			t:  RefBool,
			v: true,
		}},
		{"uint8", ref{
			t:  RefUint8,
			v: uint8(1),
		}},
		{"uint16", ref{
			t:  RefUint16,
			v: uint16(1),
		}},
		{"uint32", ref{
			t:  RefUint32,
			v: uint32(1),
		}},
		{"uint64", ref{
			t:  RefUint64,
			v: uint64(1),
		}},
		{"uint128", ref{
			t:  RefUint128,
			v: big.NewInt(1),
		}},
		{"uint256", ref{
			t:  RefUint256,
			v: big.NewInt(1),
		}},
		{"int8", ref{
			t:  RefInt8,
			v: int8(1),
		}},
		{"int16", ref{
			t:  RefInt16,
			v: int16(1),
		}},
		{"int32", ref{
			t:  RefInt32,
			v: int32(1),
		}},
		{"int64", ref{
			t:  RefInt64,
			v: int64(1),
		}},
		{"int128", ref{
			t:  RefInt128,
			v: big.NewInt(1),
		}},
		{"int256", ref{
			t:  RefInt256,
			v: big.NewInt(1),
		}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := sugar.HandlerError(
				Decode(sugar.HandlerError(Encode(tt.variable)).([]byte), tt.variable.GetGVMType()))
			if !assert.EqualValues(t, tt.variable, v) {
				t.Errorf("got = %v, want = %v", v, tt.variable)
			}
		})
	}
}