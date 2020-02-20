package uip

type TransactionId = uint64
type ActionId = uint64

type Attestation interface {
	GetTid() TransactionId
	GetAid() ActionId
	GetContent() (content []byte)
	GetSignatures() []Signature
}

