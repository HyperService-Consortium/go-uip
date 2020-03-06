package libgvm

import (
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"strconv"
)

var ExplainGVMTypeV = func(r abstraction.RefType) string {
	switch r {
	case RefUnknown:
		return "ref-unknown"
	case RefString:
		return "ref-string"
	case RefBytes:
		return "ref-bytes"
	case RefBool:
		return "ref-bool"
	case RefInt8:
		return "ref-int8"
	case RefInt16:
		return "ref-int16"
	case RefInt32:
		return "ref-int32"
	case RefInt64:
		return "ref-int64"
	case RefInt128:
		return "ref-int128"
	case RefInt256:
		return "ref-int256"
	case RefUint8:
		return "ref-uint8"
	case RefUint16:
		return "ref-uint16"
	case RefUint32:
		return "ref-uint32"
	case RefUint64:
		return "ref-uint64"
	case RefUint128:
		return "ref-uint128"
	case RefUint256:
		return "ref-uint256"
	default:
		return "ref-type(" + strconv.FormatUint(r, 10) + ")"
	}
}

func ExplainGVMType(r abstraction.RefType) string {
	return ExplainGVMTypeV(r)
}
