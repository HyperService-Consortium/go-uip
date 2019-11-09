package signaturer

import (
	signaturetype "github.com/HyperService-Consortium/go-uip/const/signature_type"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	ed25519 "golang.org/x/crypto/ed25519"
)

type Ed25519PublicKey struct {
	*BaseHexType
}

func NewEd25519PublicKeyFromBytes(b []byte) (ed *Ed25519PublicKey) {
	ed = new(Ed25519PublicKey)
	ed.BaseHexType = new(BaseHexType)
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
	ed = new(Ed25519PrivateKey)
	ed.BaseHexType = new(BaseHexType)
	*ed.BaseHexType = b
	if !ed.IsValid() {
		return nil
	}
	return
}

func (s *Ed25519PrivateKey) IsValid() bool {
	return len(*s.BaseHexType) == 64
}

func (s *Ed25519PrivateKey) ToPublic() ECCPublicKey {
	return NewEd25519PublicKeyFromBytes([]byte(ed25519.PrivateKey(*s.BaseHexType).Public().(ed25519.PublicKey)))
}

func (s *Ed25519PrivateKey) Sign(b []byte) ECCSignature {
	sig := newEd25519Sig()
	v := []byte(*s.BaseHexType)
	*sig.BaseHexType = ed25519.Sign(v, b)
	return sig
}

type Ed25519Signature struct {
	*BaseHexType
}

func newEd25519Sig() *Ed25519Signature {
	return &Ed25519Signature{
		BaseHexType: new(BaseHexType),
	}
}

func (s *Ed25519Signature) GetContent() []byte {
	return s.BaseHexType.Bytes()
}

func (s *Ed25519Signature) GetSignatureType() uiptypes.SignatureType {
	return signaturetype.Ed25519
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
