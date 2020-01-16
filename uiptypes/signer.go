package uiptypes

type Signer interface {
	GetPublicKey() PublicKey
	Sign(content SignatureContent, ctxVars ...interface{}) Signature
}

// in millisecond
type SignerOptionTimeout int
// in millisecond
type SignerOptionDuration int

type SignerOptionChainInfo = ChainInfo
