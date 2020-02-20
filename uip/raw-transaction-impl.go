package uip

type RawTransactionImpl []byte

func (t RawTransactionImpl) Serialize() ([]byte, error) {
	return t, nil
}

func (t RawTransactionImpl) Bytes() ([]byte, error) {
	return t, nil
}

func (t RawTransactionImpl) Signed() bool {
	return true
}

func (t RawTransactionImpl) Sign(signer Signer, ctxVars ...interface{}) (RawTransaction, error) {
	return t, nil
}
