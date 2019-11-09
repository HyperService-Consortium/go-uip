package uiptypes

type TransactionId = uint64
type ActionId = uint64

type Attestation interface {
	GetTid() TransactionId
	GetAid() ActionId
	GetContent() content
	GetSignatures() []Signature
}

