package uiptypes


type RawTransaction interface {
	Bytes() []byte
	Signed() bool
	Sign(Signer) (RawTransaction, error)
}
