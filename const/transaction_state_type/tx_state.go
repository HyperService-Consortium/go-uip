package TxState

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

type Type = uint64

const (
	Unknown       Type = 0 + iota
	Initing            // Not Confirmed, but uploaded by VES
	Inited             // Confirmed by VES
	Instantiating      // src: Instantiating
	Open               // dst: Instantiated, you can open
	Opened             // src: The transaction is opened, confirm it
	Closed             // dst: Okay, closed
)

//noinspection GoUnusedExportedFunction
func Description(t Type) string {
	switch t {
	case Unknown:
		return "Unknown"
	case Initing:
		return "Initing"
	case Inited:
		return "Inited"
	case Instantiating:
		return "Instantiating"
	case Open:
		return "Open"
	case Opened:
		return "Opened"
	case Closed:
		return "Closed"
	default:
		return fmt.Sprintf("TxState(%v)", t)
	}
}
