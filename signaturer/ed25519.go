package signaturer

import (
	ed25519 "golang.org/x/crypto/ed25519"
)

type Ed25519PublicKey struct {
	*BaseHexType
}

func NewEd25519PublicKeyFromBytes(b []byte) (ed *Ed25519PublicKey) {
	*ed.BaseHexType = b
	return
}

func (s *Ed25519PublicKey) IsValid() bool {
	return len(*s.BaseHexType) == 32
}

type Ed25519PrivateKey struct {
	*BaseHexType
}

func NewEd25519PrivateKeyFromBytes(b []byte) (ed *Ed25519PrivateKey) {
	*ed.BaseHexType = b
	return
}

func (s *Ed25519PrivateKey) IsValid() bool {
	return len(*s.BaseHexType) == 64
}

func (s *Ed25519PrivateKey) ToPublic() ECCPublicKey {
	return NewEd25519PublicKeyFromBytes([]byte(ed25519.PrivateKey(*s.BaseHexType).Public().(ed25519.PublicKey)))
}

func (s *Ed25519PrivateKey) Sign(b []byte) ECCSignature {
	if sig := new(Ed25519Signature); !sig.FromBytes(ed25519.Sign([]byte(*s.BaseHexType), b)) {
		return nil
	} else {
		return sig
	}
}

type Ed25519Signature struct {
	*BaseHexType
}

func (s *Ed25519Signature) IsValid() bool {
	return len(*s.BaseHexType) == 64
}

type Ed25519Signaturer struct {
}

func (ed *Ed25519Signaturer) Sign(privateKey, msg []byte) ECCSignature {
	var edpri = new(Ed25519PrivateKey)
	*edpri.BaseHexType = privateKey
	if !edpri.IsValid() {
		return nil
	}
	return edpri.Sign(msg)
}

func (ed *Ed25519Signaturer) Verify(pb, msg, eccsig []byte) bool {
	return ed25519.Verify(pb, msg, eccsig)
}
