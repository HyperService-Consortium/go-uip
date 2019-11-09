package uiptypes

type SignatureType uint32
type SignatureUnderlyingType = uint32
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
