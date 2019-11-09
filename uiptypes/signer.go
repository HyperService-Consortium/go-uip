package uiptypes

type Signer interface {
	GetPublicKey() PublicKey
	Sign(SignatureContent) Signature
}
