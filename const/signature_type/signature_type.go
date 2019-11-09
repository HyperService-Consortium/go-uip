package signaturetype

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type Type = uiptypes.SignatureType

const (
	Unknown Type = 0 + iota
	Secp256k1
	Ed25519
)
