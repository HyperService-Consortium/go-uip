package uiptypes

type Signer interface {
	GetPublicKey() PublicKey
	Sign(content SignatureContent, ctxVars ...interface{}) (Signature, error)
}

// in millisecond
type SignerOptionTimeout int
// in millisecond
type SignerOptionDuration int

type SignerOptionChainInfo = ChainInfo
