package TxState

type Type = uint64

const (
	Unknown Type = 0 + iota
	Initing // 未确认
	Inited // 确认
	Instantiating // 实例化
	Open // 请求创建Transaction
	Opened // 已创建Transaction，请求关闭
	Closed // 已关闭
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
