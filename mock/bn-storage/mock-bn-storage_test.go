package bn_storage

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"math/big"
	"testing"
)

func TestMockBNIStorage(t *testing.T) {
	b := &MockBNIStorage{}
	tester := bNIStorageTestSet{b}
	b.insertMockData(tester.MockingData())
	tester.RunTests(t)
}

func TestMockValue_IsValid(t *testing.T) {
	type fields struct {
		T value_type.Type
		V interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "int8VGood", fields: fields{
			T: value_type.Int8,
			V: int8(8),
		}, want: true},
		{name: "int8VBad", fields: fields{
			T: value_type.Int8,
			V: uint8(8),
		}},

		{name: "int16VGood", fields: fields{
			T: value_type.Int16,
			V: int16(8),
		}, want: true},
		{name: "int16VBad", fields: fields{
			T: value_type.Int16,
			V: uint16(16),
		}},

		{name: "int32VGood", fields: fields{
			T: value_type.Int32,
			V: int32(8),
		}, want: true},
		{name: "int32VBad", fields: fields{
			T: value_type.Int32,
			V: uint32(32),
		}},

		{name: "int64VGood", fields: fields{
			T: value_type.Int64,
			V: int64(8),
		}, want: true},
		{name: "int64VBad", fields: fields{
			T: value_type.Int64,
			V: uint64(64),
		}},

		{name: "int128VGood", fields: fields{
			T: value_type.Int128,
			V: big.NewInt(8),
		}, want: true},
		{name: "int128VBad", fields: fields{
			T: value_type.Int128,
			V: 8,
		}},

		{name: "int256VGood", fields: fields{
			T: value_type.Int256,
			V: big.NewInt(8),
		}, want: true},
		{name: "int256VBad", fields: fields{
			T: value_type.Int256,
			V: 8,
		}},

		{name: "uint8VBad", fields: fields{
			T: value_type.Uint8,
			V: int(8),
		}},
		{name: "uint8VGood", fields: fields{
			T: value_type.Uint8,
			V: uint8(8),
		}, want: true},

		{name: "uint16VBad", fields: fields{
			T: value_type.Uint16,
			V: int16(8),
		}},
		{name: "uint16VGood", fields: fields{
			T: value_type.Uint16,
			V: uint16(16),
		}, want: true},

		{name: "uint32VBad", fields: fields{
			T: value_type.Uint32,
			V: int32(8),
		}},
		{name: "uint32VGood", fields: fields{
			T: value_type.Uint32,
			V: uint32(32),
		}, want: true},

		{name: "uint64VBad", fields: fields{
			T: value_type.Uint64,
			V: int64(8),
		}},
		{name: "uint64VGood", fields: fields{
			T: value_type.Uint64,
			V: uint64(64),
		}, want: true},

		{name: "uint128VGood", fields: fields{
			T: value_type.Uint128,
			V: big.NewInt(8),
		}, want: true},
		{name: "uint128VBad", fields: fields{
			T: value_type.Uint128,
			V: 8,
		}},

		{name: "uint256VGood", fields: fields{
			T: value_type.Uint256,
			V: big.NewInt(8),
		}, want: true},
		{name: "uint256VBad", fields: fields{
			T: value_type.Uint256,
			V: 8,
		}},

		{name: "SliceInt8VGood", fields: fields{
			T: value_type.SliceInt8,
			V: []int8{8},
		}, want: true},
		{name: "SliceInt8VBad", fields: fields{
			T: value_type.SliceInt8,
			V: []uint8{8},
		}},

		{name: "SliceInt16VGood", fields: fields{
			T: value_type.SliceInt16,
			V: []int16{8},
		}, want: true},
		{name: "SliceInt16VBad", fields: fields{
			T: value_type.SliceInt16,
			V: []uint16{16},
		}},

		{name: "SliceInt32VGood", fields: fields{
			T: value_type.SliceInt32,
			V: []int32{8},
		}, want: true},
		{name: "SliceInt32VBad", fields: fields{
			T: value_type.SliceInt32,
			V: []uint32{32},
		}},

		{name: "SliceInt64VGood", fields: fields{
			T: value_type.SliceInt64,
			V: []int64{8},
		}, want: true},
		{name: "SliceInt64VBad", fields: fields{
			T: value_type.SliceInt64,
			V: []uint64{64},
		}},

		{name: "SliceInt128VGood", fields: fields{
			T: value_type.SliceInt128,
			V: []*big.Int{big.NewInt(8)},
		}, want: true},
		{name: "SliceInt128VBad", fields: fields{
			T: value_type.SliceInt128,
			V: []int{8},
		}},

		{name: "SliceInt256VGood", fields: fields{
			T: value_type.SliceInt256,
			V: []*big.Int{big.NewInt(8)},
		}, want: true},
		{name: "SliceInt256VBad", fields: fields{
			T: value_type.SliceInt256,
			V: []int{8},
		}},

		{name: "SliceUint8VBad", fields: fields{
			T: value_type.SliceUint8,
			V: []int{8},
		}},
		{name: "SliceUint8VGood", fields: fields{
			T: value_type.SliceUint8,
			V: []uint8{8},
		}, want: true},

		{name: "SliceUint16VBad", fields: fields{
			T: value_type.SliceUint16,
			V: []int16{8},
		}},
		{name: "SliceUint16VGood", fields: fields{
			T: value_type.SliceUint16,
			V: []uint16{16},
		}, want: true},

		{name: "SliceUint32VBad", fields: fields{
			T: value_type.SliceUint32,
			V: []int32{8},
		}},
		{name: "SliceUint32VGood", fields: fields{
			T: value_type.SliceUint32,
			V: []uint32{32},
		}, want: true},

		{name: "SliceUint64VBad", fields: fields{
			T: value_type.SliceUint64,
			V: []int64{8},
		}},
		{name: "SliceUint64VGood", fields: fields{
			T: value_type.SliceUint64,
			V: []uint64{64},
		}, want: true},

		{name: "SliceUint128VGood", fields: fields{
			T: value_type.SliceUint128,
			V: []*big.Int{big.NewInt(8)},
		}, want: true},
		{name: "SliceUint128VBad", fields: fields{
			T: value_type.SliceUint128,
			V: []int{8},
		}},

		{name: "SliceUint256VGood", fields: fields{
			T: value_type.SliceUint256,
			V: []*big.Int{big.NewInt(8)},
		}, want: true},
		{name: "SliceUint256VBad", fields: fields{
			T: value_type.SliceUint256,
			V: []int{8},
		}},
		{name: "String", fields: fields{
			T: value_type.String,
			V: "",
		}, want: true},
		{name: "StringBad", fields: fields{
			T: value_type.String,
			V: []byte(""),
		}},
		{name: "Bool", fields: fields{
			T: value_type.Bool,
			V: true,
		}, want: true},
		{name: "BoolBad", fields: fields{
			T: value_type.Bool,
			V: "",
		}},

		{name: "StringNil", fields: fields{
			T: value_type.String,
			V: nil,
		}},
		{name: "BoolNil", fields: fields{
			T: value_type.Bool,
			V: nil,
		}},
		{name: "BytesNil", fields: fields{
			T: value_type.Bytes,
			V: nil,
		}, want: true},
		{name: "SliceWithNil", fields: fields{
			T: value_type.SliceUint256,
			V: nil,
		}, want: true},
		{name: "BigIntWithNil", fields: fields{
			T: value_type.Uint256,
			V: nil,
		}},
		{name: "Unknown", fields: fields{
			T: value_type.Unknown,
			V: nil,
		}},
		{name: "Unknown", fields: fields{
			T: value_type.Unknown,
			V: "",
		}},
		{name: "Unknown", fields: fields{
			T: value_type.Unknown,
			V: []int8{},
		}},
		{name: "Unknown", fields: fields{
			T: value_type.Unknown,
			V: big.NewInt(1),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MockValue{
				T: tt.fields.T,
				V: tt.fields.V,
			}
			if got := m.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
