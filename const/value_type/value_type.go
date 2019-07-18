package value_type

type Type = uint16

const (
	Bytes Type = iota
	String
	Uint8
	Uint16
	Uint32
	Uint64
	Uint128
	Uint256
	Int8
	Int16
	Int32
	Int64
	Int128
	Int256
	SliceUint8
	SliceUint16
	SliceUint32
	SliceUint64
	SliceUint128
	SliceUint256
	SliceInt8
	SliceInt16
	SliceInt32
	SliceInt64
	SliceInt128
	SliceInt256

	// Slice
	Bool
)
