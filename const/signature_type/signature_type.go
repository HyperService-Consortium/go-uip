package signaturetype

import "github.com/HyperService-Consortium/go-uip/uip"

type Type = uip.SignatureType

const (
	Unknown Type = 0 + iota
	Secp256k1
	Ed25519
)
