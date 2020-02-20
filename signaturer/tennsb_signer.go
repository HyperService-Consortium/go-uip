package signaturer

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type TendermintNSBSigner struct {
	prikey *Ed25519PrivateKey
	pubKey *Ed25519PublicKey
}

var ErrConvert = errors.New("could not covert it to signer's private key")

func NewTendermintNSBSigner(pri []byte) (ten *TendermintNSBSigner, err error) {
	ten = new(TendermintNSBSigner)
	ten.prikey = NewEd25519PrivateKeyFromBytes(pri)
	if ten.prikey == nil {
		return nil, ErrConvert
	}
	ten.pubKey = ten.prikey.ToPublic().(*Ed25519PublicKey)
	return
}

func (ten *TendermintNSBSigner) Sign(b []byte, ctxVars ...interface{}) (uip.Signature, error) {
	return ten.prikey.Sign(b), nil
}

func (ten *TendermintNSBSigner) GetPublicKey() []byte {
	return ten.pubKey.Bytes()
}
