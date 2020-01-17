package signaturer

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"math/big"

	signaturetype "github.com/HyperService-Consortium/go-uip/const/signature_type"
	secp256k1 "github.com/HyperService-Consortium/go-uip/signaturer/go-ethereum-secp256k1"
	ed25519 "golang.org/x/crypto/ed25519"
)

var (
	secp256k1N, _  = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	secp256k1halfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

type Secp256k1PublicKey struct {
	*BaseHexType
}

// S256 returns an instance of the secp256k1 curve.
func S256() elliptic.Curve {
	return secp256k1.S256()
}

// toECDSA creates a private key with the given D value. The strict parameter
// controls whether the key's length should be enforced at the curve size or
// it can also accept legacy encodings (0 prefixes).
func toECDSA(d []byte, strict bool) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = S256()
	if strict && 8*len(d) != priv.Params().BitSize {
		return nil, fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	if priv.D.Cmp(secp256k1N) >= 0 {
		return nil, fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return nil, fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return nil, errors.New("invalid private key")
	}
	return priv, nil
}

func NewSecp256k1PublicKeyFromBytes(b []byte) (se *Secp256k1PublicKey) {
	se = new(Secp256k1PublicKey)
	*se.BaseHexType = b
	return
}

func (s *Secp256k1PublicKey) IsValid() bool {
	return len(*s.BaseHexType) == 32
}

type Secp256k1PrivateKey struct {
	*BaseHexType
}

func (s *Secp256k1PrivateKey) IsValid() bool {
	return len(*s.BaseHexType) == 64
}

func (s *Secp256k1PrivateKey) ToPublic() ECCPublicKey {
	return NewSecp256k1PublicKeyFromBytes(ed25519.PrivateKey(*s.BaseHexType).Public().(ed25519.PublicKey))
}

func (s *Secp256k1PrivateKey) Sign(b []byte) (ECCSignature, error) {
	sig := new(Secp256k1Signature)
	if err := sig.FromBytes(ed25519.Sign([]byte(*s.BaseHexType), b)); err != nil {
		return nil, err
	} else {
		return sig, nil
	}
}

type Secp256k1Signature struct {
	*BaseHexType
}

func (s *Secp256k1Signature) IsValid() bool {
	return len(*s.BaseHexType) == 65
}

func (s *Secp256k1Signature) GetContent() []byte {
	return s.BaseHexType.Bytes()
}

func (s *Secp256k1Signature) GetSignatureType() uiptypes.SignatureType {
	return signaturetype.Secp256k1
}

type Secp256k1Signaturer struct {
}

func (ed *Secp256k1Signaturer) Sign(privateKey, msg []byte) (ECCSignature, error) {
	var edpri = new(Secp256k1PrivateKey)
	*edpri.BaseHexType = privateKey
	if !edpri.IsValid() {
		return nil, errors.New("invalid private key")
	}
	return edpri.Sign(msg)
}

func (ed *Secp256k1Signaturer) Verify(pb, msg, eccsig []byte) bool {
	return ed25519.Verify(pb, msg, eccsig)
}
