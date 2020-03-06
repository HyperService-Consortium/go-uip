package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
)

const (
	RefUnknown abstraction.RefType = iota
	RefBytes
	RefString
	RefUint8
	RefUint16 // 4
	RefUint32
	RefUint64
	RefUint128
	RefUint256
	RefInt8 // 9
	RefInt16
	RefInt32
	RefInt64
	RefInt128
	RefInt256 // 14

	// Slice
	RefBool // 15
	Length
)
