package uiptypes

type SignatureTypeUnderlyingType = uint32
type SignatureType SignatureTypeUnderlyingType
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
