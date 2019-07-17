package signaturer

import (
	types "github.com/Myriad-Dreamin/go-uip/types"
)

var _ ECCPublicKey = new(Ed25519PublicKey)
var _ ECCPrivateKey = NewEd25519PrivateKeyFromBytes([]byte{0, 1})
var _s types.Signer = NewTendermintNSBSigner([]byte{
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
})
var _ ECCSignature = new(Ed25519Signature)
var _ = _s.Sign([]byte{0, 1})
var _ ECCSignaturer = new(Ed25519Signaturer)
