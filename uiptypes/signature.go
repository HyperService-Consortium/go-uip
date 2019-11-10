package uiptypes

type SignatureUnderlyingType = uint32
type SignatureType SignatureUnderlyingType
type SignatureContent = []byte
type SignedContent = []byte

type BaseSignature interface {
	GetSignatureType() SignatureType
	GetContent() SignedContent
}

type Signature interface {
	HexType
	BaseSignature
	IsValid() bool
}
