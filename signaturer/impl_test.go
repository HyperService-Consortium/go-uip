package signaturer

var _ ECCPublicKey = new(Ed25519PublicKey)
var _ ECCPrivateKey = NewEd25519PrivateKeyFromBytes([]byte{0, 1})
var _ ECCSignature = new(Ed25519Signature)
var _ ECCSignaturer = new(Ed25519Signaturer)
