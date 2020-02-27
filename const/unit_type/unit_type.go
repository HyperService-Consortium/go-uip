package UnitType

type Type = uint64

const (
	Default Type = iota
	Wei
	Iew
	Ether
)

var Mapping = make(map[string]Type)

func init() {
	Mapping[""] = Default
	Mapping["ether"] = Ether
	Mapping["wei"] = Wei
	Mapping["iew"] = Iew
}
