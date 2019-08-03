package signaturetype

type Type = uint32

const (
	Unknown Type = 0 + iota
	Secp256k1
	Ed25519
)
