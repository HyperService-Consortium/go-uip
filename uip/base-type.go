package uip

type HexType interface {
	Bytes() []byte
	String() string
	FromBytes([]byte) error
	FromString(string) error
	Equal(HexType) bool
}

type BaseSlice interface {
	Len() int
	Cap() int
}
