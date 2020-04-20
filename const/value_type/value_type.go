package value_type

type Type = uint16

//noinspection GoUnusedConst
const (
	Unknown Type = iota
	Bytes
	String
	Uint8
	Uint16 // 4
	Uint32
	Uint64
	Uint128
	Uint256
	Int8 // 9
	Int16
	Int32
	Int64
	Int128
	Int256 // 14
	Bool

	SliceUint8
	SliceUint16
	SliceUint32
	SliceUint64 // 19
	SliceUint128
	SliceUint256
	SliceInt8
	SliceInt16
	SliceInt32 // 24
	SliceInt64
	SliceInt128
	SliceInt256 // 27

	// Slice
	Length
)

func FromString(desc string) Type {
	switch desc {
	case "bytes":
		return Bytes
	case "string":
		return String
	case "uint8":
		return Uint8
	case "uint16":
		return Uint16
	case "uint32":
		return Uint32
	case "uint64":
		return Uint64
	case "uint128":
		return Uint128
	case "uint256":
		return Uint256
	case "int8":
		return Int8
	case "int16":
		return Int16
	case "int32":
		return Int32
	case "int64":
		return Int64
	case "int128":
		return Int128
	case "int256":
		return Int256
	case "[]uint8":
		return SliceUint8
	case "[]uint16":
		return SliceUint16
	case "[]uint32":
		return SliceUint32
	case "[]uint64":
		return SliceUint64
	case "[]uint128":
		return SliceUint128
	case "[]uint256":
		return SliceUint256
	case "[]int8":
		return SliceInt8
	case "[]int16":
		return SliceInt16
	case "[]int32":
		return SliceInt32
	case "[]int64":
		return SliceInt64
	case "[]int128":
		return SliceInt128
	case "[]int256":
		return SliceInt256
	case "bool":
		return Bool
	default:
		return Unknown
	}
}
