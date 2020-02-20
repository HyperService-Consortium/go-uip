package standard

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"math/big"
	"reflect"
)

var bytesType = reflect.TypeOf(new([]byte)).Elem()
var bigIntType = reflect.TypeOf(new(big.Int))

func IsValidStorageValue(v interface{}, t value_type.Type) bool {
	runtimeT := reflect.TypeOf(v)
	switch t {
	case value_type.Bytes,
		value_type.SliceInt8, value_type.SliceInt16, value_type.SliceInt32,
		value_type.SliceInt64, value_type.SliceInt128, value_type.SliceInt256,
		value_type.SliceUint8, value_type.SliceUint16, value_type.SliceUint32,
		value_type.SliceUint64, value_type.SliceUint128, value_type.SliceUint256:
		if v == nil {
			return true
		}
		switch t {
		case value_type.Bytes:
			if runtimeT != bytesType {
				return false
			}
		case value_type.SliceUint8:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint8 {
				return false
			}
		case value_type.SliceUint16:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint16 {
				return false
			}
		case value_type.SliceUint32:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint32 {
				return false
			}
		case value_type.SliceUint64:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint64 {
				return false
			}
		case value_type.SliceUint128:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		case value_type.SliceUint256:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		case value_type.SliceInt8:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int8 {
				return false
			}
		case value_type.SliceInt16:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int16 {
				return false
			}
		case value_type.SliceInt32:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int32 {
				return false
			}
		case value_type.SliceInt64:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int64 {
				return false
			}
		case value_type.SliceInt128:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		case value_type.SliceInt256:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		default:
			return false
		}
	case value_type.String, value_type.Bool, value_type.Uint128, value_type.Uint256, value_type.Int128, value_type.Int256,
		value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64,
		value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64:
		if v == nil {
			return false
		}
		switch t {
		case value_type.String:
			if runtimeT.Kind() != reflect.String {
				return false
			}
		case value_type.Uint128, value_type.Uint256, value_type.Int128, value_type.Int256:
			if runtimeT != bigIntType {
				return false
			}
		case value_type.Uint8:
			if runtimeT.Kind() != reflect.Uint8 {
				return false
			}
		case value_type.Uint16:
			if runtimeT.Kind() != reflect.Uint16 {
				return false
			}
		case value_type.Uint32:
			if runtimeT.Kind() != reflect.Uint32 {
				return false
			}
		case value_type.Uint64:
			if runtimeT.Kind() != reflect.Uint64 {
				return false
			}
		case value_type.Int8:
			if runtimeT.Kind() != reflect.Int8 {
				return false
			}
		case value_type.Int16:
			if runtimeT.Kind() != reflect.Int16 {
				return false
			}
		case value_type.Int32:
			if runtimeT.Kind() != reflect.Int32 {
				return false
			}
		case value_type.Int64:
			if runtimeT.Kind() != reflect.Int64 {
				return false
			}
		case value_type.Bool:
			if runtimeT.Kind() != reflect.Bool {
				return false
			}
		default:
			return false
		}
	default:
		return false
	}
	return true
}

