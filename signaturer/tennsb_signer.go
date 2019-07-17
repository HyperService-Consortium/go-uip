package signaturer

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

func (ten *TendermintNSBSigner) Sign(b []byte) []byte {
	sig := ten.prikey.Sign(b)
	if sig != nil {
		return sig.Bytes()
	}
	return nil
}

func (ten *TendermintNSBSigner) GetPublicKey() []byte {
	return ten.pubKey.Bytes()
}
