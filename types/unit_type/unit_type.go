package UnitType

type Type = uint64

const (
	Default Type = iota
	Wei
	Iew
)

var Mapping = make(map[string]Type)

func init() {
	Mapping[""] = Default
	Mapping["wei"] = Wei
	Mapping["iew"] = Iew
}
