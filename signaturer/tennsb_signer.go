package signaturer

type TendermintNSBSigner struct {
	prikey *Ed25519PrivateKey
	pubKey *Ed25519PublicKey
}

func NewTendermintNSBSigner(pri []byte) (ten *TendermintNSBSigner) {
	ten = new(TendermintNSBSigner)
	ten.prikey = NewEd25519PrivateKeyFromBytes(pri)
	ten.pubKey = ten.prikey.ToPublic().(*Ed25519PublicKey)
	return
}

func (ten *TendermintNSBSigner) Sign(b []byte) []byte {
	return ten.prikey.Sign(b).Bytes()
}

func (ten *TendermintNSBSigner) GetPublicKey() []byte {
	return ten.pubKey.Bytes()
}
