package signaturer

import "github.com/Myriad-Dreamin/go-uip/types"

type TendermintNSBSigner struct {
	prikey *Ed25519PrivateKey
	pubKey *Ed25519PublicKey
}

func NewTendermintNSBSigner(pri []byte) (ten *TendermintNSBSigner) {
	ten = new(TendermintNSBSigner)
	ten.prikey = NewEd25519PrivateKeyFromBytes(pri)
	if ten.prikey == nil {
		return nil
	}
	ten.pubKey = ten.prikey.ToPublic().(*Ed25519PublicKey)
	return
}

func (ten *TendermintNSBSigner) Sign(b []byte) types.Signature {
	return ten.prikey.Sign(b)
}

func (ten *TendermintNSBSigner) GetPublicKey() []byte {
	return ten.pubKey.Bytes()
}
