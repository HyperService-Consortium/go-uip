package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"github.com/stretchr/testify/assert"
	"math/big"
	"reflect"
	"testing"
)

func TestTryConvertUnsigned(t *testing.T) {
	type args struct {
		k abstraction.Ref
	}
	tests := []struct {
		name    string
		args    args
		want    abstraction.Ref
		wantErr bool
	}{
		{"uint8", args{k: Uint8(1)}, Uint8(1), false},
		{"uint16", args{k: Uint16(1)}, Uint16(1), false},
		{"uint32", args{k: Uint32(1)}, Uint32(1), false},
		{"uint64", args{k: Uint64(1)}, Uint64(1), false},
		{"uint128", args{k: (*Uint128)(big.NewInt(1))}, (*Uint128)(big.NewInt(1)), false},
		{"uint256", args{k: (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(1)), false},
		{"int8", args{k: Int8(1)}, Uint8(1), false},
		{"int16", args{k: Int16(1)}, Uint16(1), false},
		{"int32", args{k: Int32(1)}, Uint32(1), false},
		{"int64", args{k: Int64(1)}, Uint64(1), false},
		{"int128", args{k: (*Int128)(big.NewInt(1))}, (*Uint128)(big.NewInt(1)), false},
		{"int256", args{k: (*Int256)(big.NewInt(1))}, (*Uint256)(big.NewInt(1)), false},
		{"negative-int8", args{k: Int8(-1)}, Uint8(1), true},
		{"negative-int16", args{k: Int16(-1)}, Uint16(1), true},
		{"negative-int32", args{k: Int32(-1)}, Uint32(1), true},
		{"negative-int64", args{k: Int64(-1)}, Uint64(1), true},
		{"negative-int128", args{k: (*Int128)(big.NewInt(-1))}, (*Uint128)(big.NewInt(1)), true},
		{"negative-int256", args{k: (*Int256)(big.NewInt(-1))}, (*Uint256)(big.NewInt(1)), true},
		{"not arithmetic", args{k: Bool(false)}, (*Uint256)(big.NewInt(1)), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryConvertUnsigned(tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryConvertUnsigned() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil && (!assert.EqualValues(t, tt.want, got) || reflect.TypeOf(tt.want) != reflect.TypeOf(got)) {
				t.Errorf("TryConvertUnsigned() got = (%v, %T), want (%v, %T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestExtendBit(t *testing.T) {
	type args struct {
		l        abstraction.Ref
		bitCount int
	}
	tests := []struct {
		name    string
		args    args
		want    abstraction.Ref
		wantErr bool
	}{
		{"uint8-to-uint256", args{bitCount: 256, l: Uint8(1)}, (*Uint256)(big.NewInt(1)), false},
		{"uint16-to-uint256", args{bitCount: 256, l: Uint16(1)}, (*Uint256)(big.NewInt(1)), false},
		{"uint32-to-uint256", args{bitCount: 256, l: Uint32(1)}, (*Uint256)(big.NewInt(1)), false},
		{"uint64-to-uint256", args{bitCount: 256, l: Uint64(1)}, (*Uint256)(big.NewInt(1)), false},
		{"uint8-to-uint64", args{bitCount: 64, l: Uint8(1)}, Uint64(1), false},
		{"uint16-to-uint64", args{bitCount: 64, l: Uint16(1)}, Uint64(1), false},
		{"uint32-to-uint64", args{bitCount: 64, l: Uint32(1)}, Uint64(1), false},
		{"uint64-to-uint64", args{bitCount: 64, l: Uint64(1)}, Uint64(1), false},
		{"uint128-to-uint256", args{bitCount: 256, l: (*Uint128)(big.NewInt(1))}, (*Uint256)(big.NewInt(1)), false},
		{"uint256-to-uint256", args{bitCount: 256, l: (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(1)), false},
		{"int8-to-int256", args{bitCount: 256, l: Int8(1)}, (*Int256)(big.NewInt(1)), false},
		{"int16-to-int256", args{bitCount: 256, l: Int16(1)}, (*Int256)(big.NewInt(1)), false},
		{"int32-to-int256", args{bitCount: 256, l: Int32(1)}, (*Int256)(big.NewInt(1)), false},
		{"int64-to-int256", args{bitCount: 256, l: Int64(1)}, (*Int256)(big.NewInt(1)), false},
		{"int8-to-int64", args{bitCount: 64, l: Int8(1)}, Int64(1), false},
		{"int16-to-int64", args{bitCount: 64, l: Int16(1)}, Int64(1), false},
		{"int32-to-int64", args{bitCount: 64, l: Int32(1)}, Int64(1), false},
		{"int64-to-int64", args{bitCount: 64, l: Int64(1)}, Int64(1), false},
		{"int128-to-int256", args{bitCount: 256, l: (*Int128)(big.NewInt(1))}, (*Int256)(big.NewInt(1)), false},
		{"int256-to-int256", args{bitCount: 256, l: (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(1)), false},
		{"not arithmetic", args{bitCount: 256, l: Bool(false)}, (*Uint256)(big.NewInt(1)), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtendBit(tt.args.l, tt.args.bitCount)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryConvertUnsigned() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil && (!assert.EqualValues(t, tt.want, got) || reflect.TypeOf(tt.want) != reflect.TypeOf(got)) {
				t.Errorf("TryConvertUnsigned() got = (%v, %T), want (%v, %T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestGetArithmeticTraits(t *testing.T) {
	type args struct {
		l abstraction.Ref
	}
	tests := []struct {
		name         string
		args         args
		wantIntegral bool
		wantUnsigned bool
		wantK        int
	}{
		{"uint8", args{Uint8(1)}, true, true, 8},
		{"uint16", args{Uint16(1)}, true, true, 16},
		{"uint32", args{Uint32(1)}, true, true, 32},
		{"uint64", args{Uint64(1)}, true, true, 64},
		{"int8", args{Int8(1)}, true, false, 8},
		{"int16", args{Int16(1)}, true, false, 16},
		{"int32", args{Int32(1)}, true, false, 32},
		{"int64", args{Int64(1)}, true, false, 64},
		{"int128", args{(*Int128)(big.NewInt(1))}, true, false, 128},
		{"int256", args{(*Int256)(big.NewInt(1))}, true, false, 256},
		{"uint128", args{(*Uint128)(big.NewInt(1))}, true, true, 128},
		{"uint256", args{(*Uint256)(big.NewInt(1))}, true, true, 256},
		{"not arithemetic", args{Bool(false)}, false, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIntegral, gotUnsigned, gotK := GetArithmeticalCharacteristic(tt.args.l)
			if gotIntegral != tt.wantIntegral {
				t.Errorf("GetArithmeticalCharacteristic() gotIntegral = %v, want %v", gotIntegral, tt.wantIntegral)
			}
			if gotUnsigned != tt.wantUnsigned {
				t.Errorf("GetArithmeticalCharacteristic() gotUnsigned = %v, want %v", gotUnsigned, tt.wantUnsigned)
			}
			if gotK != tt.wantK {
				t.Errorf("GetArithmeticalCharacteristic() gotK = %v, want %v", gotK, tt.wantK)
			}
		})
	}
}

func TestTryAlignArithmetic(t *testing.T) {
	type args struct {
		l abstraction.Ref
		r abstraction.Ref
	}
	tests := []struct {
		name    string
		args    args
		want    abstraction.Ref
		want1   abstraction.Ref
		wantErr bool
	}{
		{"align-uint8-uint8", args{Uint8(2), Uint8(1)}, Uint8(2), Uint8(1), false},
		{"align-uint8-uint64", args{Uint8(2), Uint64(1)}, Uint64(2), Uint64(1), false},
		{"align-uint64-uint8", args{Uint64(2), Uint8(1)}, Uint64(2), Uint64(1), false},
		{"align-uint8-uint256", args{Uint8(2), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-uint8", args{(*Uint256)(big.NewInt(2)), Uint8(1)}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint16-uint16", args{Uint16(2), Uint16(1)}, Uint16(2), Uint16(1), false},
		{"align-uint16-uint64", args{Uint16(2), Uint64(1)}, Uint64(2), Uint64(1), false},
		{"align-uint64-uint16", args{Uint64(2), Uint16(1)}, Uint64(2), Uint64(1), false},
		{"align-uint16-uint256", args{Uint16(2), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-uint16", args{(*Uint256)(big.NewInt(2)), Uint16(1)}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint32-uint32", args{Uint32(2), Uint32(1)}, Uint32(2), Uint32(1), false},
		{"align-uint32-uint64", args{Uint32(2), Uint64(1)}, Uint64(2), Uint64(1), false},
		{"align-uint64-uint32", args{Uint64(2), Uint32(1)}, Uint64(2), Uint64(1), false},
		{"align-uint32-uint256", args{Uint32(2), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-uint32", args{(*Uint256)(big.NewInt(2)), Uint32(1)}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint64-uint64", args{Uint64(2), Uint64(1)}, Uint64(2), Uint64(1), false},
		{"align-uint64-uint256", args{Uint64(2), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-uint64", args{(*Uint256)(big.NewInt(2)), Uint64(1)}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint128-uint128", args{(*Uint128)(big.NewInt(2)), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint128-uint256", args{(*Uint128)(big.NewInt(2)), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-uint128", args{(*Uint256)(big.NewInt(2)), (*Uint128)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-uint256", args{(*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint256-to-not-arithmetic-type", args{(*Uint256)(big.NewInt(2)), Bool(false)}, Uint8(2), Uint8(0), true},

		{"align-int8-int8", args{Int8(2), Int8(1)}, Int8(2), Int8(1), false},
		{"align-int8-int64", args{Int8(2), Int64(1)}, Int64(2), Int64(1), false},
		{"align-int64-int8", args{Int64(2), Int8(1)}, Int64(2), Int64(1), false},
		{"align-int8-int256", args{Int8(2), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-int8", args{(*Int256)(big.NewInt(2)), Int8(1)}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int16-int16", args{Int16(2), Int16(1)}, Int16(2), Int16(1), false},
		{"align-int16-int64", args{Int16(2), Int64(1)}, Int64(2), Int64(1), false},
		{"align-int64-int16", args{Int64(2), Int16(1)}, Int64(2), Int64(1), false},
		{"align-int16-int256", args{Int16(2), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-int16", args{(*Int256)(big.NewInt(2)), Int16(1)}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int32-int32", args{Int32(2), Int32(1)}, Int32(2), Int32(1), false},
		{"align-int32-int64", args{Int32(2), Int64(1)}, Int64(2), Int64(1), false},
		{"align-int64-int32", args{Int64(2), Int32(1)}, Int64(2), Int64(1), false},
		{"align-int32-int256", args{Int32(2), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-int32", args{(*Int256)(big.NewInt(2)), Int32(1)}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int64-int64", args{Int64(2), Int64(1)}, Int64(2), Int64(1), false},
		{"align-int64-int256", args{Int64(2), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-int64", args{(*Int256)(big.NewInt(2)), Int64(1)}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int128-int128", args{(*Int128)(big.NewInt(2)), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int128-int256", args{(*Int128)(big.NewInt(2)), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-int128", args{(*Int256)(big.NewInt(2)), (*Int128)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-int256", args{(*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int256-to-not-arithmetic-type", args{(*Int256)(big.NewInt(2)), Bool(false)}, Int8(2), Int8(0), true},

		{"align-int256-uint256", args{(*Int256)(big.NewInt(2)), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-int64-uint256", args{Int64(2), (*Uint256)(big.NewInt(1))}, (*Uint256)(big.NewInt(2)), (*Uint256)(big.NewInt(1)), false},
		{"align-uint64-int256", args{Uint64(2), (*Int256)(big.NewInt(1))}, (*Int256)(big.NewInt(2)), (*Int256)(big.NewInt(1)), false},
		{"align-int64-uint64", args{Int64(2), Uint64(1)}, Uint64(2), Uint64(1), false},
		{"align-int8-uint64", args{Int8(2), Uint64(1)}, Uint64(2), Uint64(1), false},
		{"align-uint8-int64", args{Uint8(2), Int64(1)}, Int64(2), Int64(1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := TryAlignArithmetic(tt.args.l, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryAlignArithmetic() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil && (!assert.EqualValues(t, tt.want, got) || reflect.TypeOf(tt.want) != reflect.TypeOf(got)) {
				t.Errorf("TryAlignArithmetic() got = (%v, %T), want (%v, %T)", got, got, tt.want, tt.want)
			} else if err == nil && (!assert.EqualValues(t, tt.want1, got1) || reflect.TypeOf(tt.want1) != reflect.TypeOf(got1)) {
				t.Errorf("TryAlignArithmetic() got1 = (%v, %T), want1 (%v, %T)", got1, got1, tt.want1, tt.want1)
			}
		})
	}
}

func TestUnsafeShrinkType(t *testing.T) {
	type args struct {
		value     interface{}
		paramType abstraction.RefType
	}
	tests := []struct {
		name    string
		args    args
		want    abstraction.Ref
		wantErr bool
	}{
		{"uint8", args{uint64(2), RefUint8}, Uint8(2), false},
		{"uint16", args{uint64(2), RefUint16}, Uint16(2), false},
		{"uint32", args{uint64(2), RefUint32}, Uint32(2), false},
		{"uint64", args{uint64(2), RefUint64}, Uint64(2), false},

		{"int8", args{int64(2), RefInt8}, Int8(2), false},
		{"int16", args{int64(2), RefInt16}, Int16(2), false},
		{"int32", args{int64(2), RefInt32}, Int32(2), false},
		{"int64", args{int64(2), RefInt64}, Int64(2), false},

		{"uint128", args{big.NewInt(2), RefUint128}, (*Uint128)(big.NewInt(2)), false},
		{"uint256", args{big.NewInt(2), RefUint256}, (*Uint256)(big.NewInt(2)), false},

		{"int128", args{big.NewInt(2), RefInt128}, (*Int128)(big.NewInt(2)), false},
		{"int256", args{big.NewInt(2), RefInt256}, (*Int256)(big.NewInt(2)), false},

		{"not-arithmetic-type", args{big.NewInt(2), RefBool}, (*Int256)(big.NewInt(2)), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnsafeShrinkType(tt.args.value, tt.args.paramType)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnsafeShrinkType() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err == nil && (!assert.EqualValues(t, tt.want, got) || reflect.TypeOf(tt.want) != reflect.TypeOf(got)) {
				t.Errorf("TryAlignArithmetic() got = (%v, %T), want (%v, %T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		l abstraction.Ref
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"uint8-true", args{Uint8(1)}, true, false},
		{"uint16-true", args{Uint16(1)}, true, false},
		{"uint32-true", args{Uint32(1)}, true, false},
		{"uint64-true", args{Uint64(1)}, true, false},
		{"uint8-false", args{Uint8(0)}, false, false},
		{"uint16-false", args{Uint16(0)}, false, false},
		{"uint32-false", args{Uint32(0)}, false, false},
		{"uint64-false", args{Uint64(0)}, false, false},

		{"uint128-true", args{(*Uint128)(big.NewInt(1))}, true, false},
		{"uint128-false", args{(*Uint128)(big.NewInt(0))}, false, false},
		{"uint256-true", args{(*Uint256)(big.NewInt(1))}, true, false},
		{"uint256-false", args{(*Uint256)(big.NewInt(0))}, false, false},

		{"int8-true", args{Int8(1)}, true, false},
		{"int16-true", args{Int16(1)}, true, false},
		{"int32-true", args{Int32(1)}, true, false},
		{"int64-true", args{Int64(1)}, true, false},
		{"int8-true", args{Int8(-1)}, true, false},
		{"int16-true", args{Int16(-1)}, true, false},
		{"int32-true", args{Int32(-1)}, true, false},
		{"int64-true", args{Int64(-1)}, true, false},
		{"int8-false", args{Int8(0)}, false, false},
		{"int16-false", args{Int16(0)}, false, false},
		{"int32-false", args{Int32(0)}, false, false},
		{"int64-false", args{Int64(0)}, false, false},

		{"int128-true", args{(*Int128)(big.NewInt(1))}, true, false},
		{"int128-true", args{(*Int128)(big.NewInt(-1))}, true, false},
		{"int128-false", args{(*Int128)(big.NewInt(0))}, false, false},
		{"int256-true", args{(*Int256)(big.NewInt(1))}, true, false},
		{"int256-true", args{(*Int256)(big.NewInt(-1))}, true, false},
		{"int256-false", args{(*Int256)(big.NewInt(0))}, false, false},

		{"bool-true", args{Bool(true)}, true, false},
		{"bool-false", args{Bool(false)}, false, false},

		{"error-type", args{String("")}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtendType(t *testing.T) {
	type args struct {
		l abstraction.Ref
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtendType(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtendType() = %v, want %v", got, tt.want)
			}
		})
	}
}
