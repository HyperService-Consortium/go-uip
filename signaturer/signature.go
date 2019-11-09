package signaturer

import (
	signaturetype "github.com/HyperService-Consortium/go-uip/const/signature_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
)

type Signature struct {
	BaseHexType
	SignatureType uiptypes.SignatureType
	isValid       func([]byte) bool
}

func (s *Signature) GetSignatureType() uiptypes.SignatureType {
	return s.SignatureType
}

func (s *Signature) GetContent() []byte {
	return s.BaseHexType
}

func (s *Signature) IsValid() bool {
	return s.isValid(s.BaseHexType)
}

func FromBaseSignature(b uiptypes.BaseSignature) uiptypes.Signature {
	return &Signature{
		BaseHexType:   b.GetContent(),
		SignatureType: b.GetSignatureType(),
		isValid:       validator(b.GetSignatureType()),
	}
}

func FromRaw(b []byte, t uint32) uiptypes.Signature {
	return &Signature{
		BaseHexType:   b,
		SignatureType: uiptypes.SignatureType(t),
		isValid:       validator(uiptypes.SignatureType(t)),
	}
}

func ed25519IsValid(b []byte) bool {
	return len(b) == 64
}
func secp256k1IsValid(b []byte) bool {
	return len(b) == 65
}
func unknownIsValid(b []byte) bool {
	return false
}

func validator(signatureType uiptypes.SignatureType) func([]byte) bool {
	switch signatureType {
	case signaturetype.Ed25519:
		return ed25519IsValid
	case signaturetype.Secp256k1:
		return secp256k1IsValid
	}
	return unknownIsValid
}
