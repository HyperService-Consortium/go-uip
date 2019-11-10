package uiptypes


type RawTransaction interface {
	// must have no logic error
	Serialize() ([]byte, error)
	Signed() bool
	Sign(Signer) (RawTransaction, error)
}
