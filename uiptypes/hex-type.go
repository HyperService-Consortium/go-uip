package uiptypes

type HexType interface {
	Bytes() []byte
	String() string
	FromBytes([]byte) bool
	FromString(string) bool
	Equal(HexType) bool
}
