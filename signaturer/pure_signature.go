package signaturer

import (
	signaturetype "github.com/Myriad-Dreamin/go-uip/const/signature_type"
	types "github.com/Myriad-Dreamin/go-uip/types"
)

type PureSignature struct {
	BaseHexType
	SignatureType uint32
	isValid       func([]byte) bool
}

func (s *PureSignature) GetSignatureType() uint32 {
	return s.SignatureType
}

func (s *PureSignature) GetContent() []byte {
	return s.BaseHexType
}

func (s *PureSignature) IsValid() bool {
	return s.isValid(s.BaseHexType)
}

func FromBaseSignature(b types.BaseSignature) types.Signature {
	return &PureSignature{
		BaseHexType:   b.GetContent(),
		SignatureType: b.GetSignatureType(),
		isValid:       validator(b.GetSignatureType()),
	}
}

func FromRaw(b []byte, t uint32) types.Signature {
	return &PureSignature{
		BaseHexType:   b,
		SignatureType: t,
		isValid:       validator(t),
	}
}

var ed25519IsValid = func(b []byte) bool {
	return len(b) == 64
}
var secp256k1IsValid = func(b []byte) bool {
	return len(b) == 65
}
var unknownIsValid = func(b []byte) bool {
	return false
}

func validator(signatureType uint32) func([]byte) bool {
	switch signatureType {
	case signaturetype.Ed25519:
		return ed25519IsValid
	case signaturetype.Secp256k1:
		return secp256k1IsValid
	}
	return unknownIsValid
}
