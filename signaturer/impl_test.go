package signature

var _ ECCPublicKey = new(Ed25519PublicKey)
var _ ECCPrivateKey = new(Ed25519PrivateKey)
var _ ECCSignature = new(Ed25519Signature)
var _ ECCSignaturer = new(Ed25519Signaturer)
