package opintent

type DataType interface {
	GetValue() interface{}
	GetDesc() []byte
}

type help_info string
type chain_id uint64
type meta []byte
type Encoder interface {
}

type Decoder interface {
}

type ProcessorInterface interface {
	CheckAddress([]byte) bool
	// ParseData([]DataType, meta, chain_id) (bool, []byte, help_info)
	// GetEncoder() Encoder
	// GetDecoder() Decoder
}
