package TxState

type Type = uint64

const (
	Unknown Type = 0 + iota
	Initing
	Inited
	Instantiating
	Instantiated
	Open
	Opened
	Closed
)

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
	case Instantiated:
		return "Instantiated"
	case Open:
		return "Open"
	case Opened:
		return "Opened"
	case Closed:
		return "Closed"
	default:
		return "no such status"
	}
}
