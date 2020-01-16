package base_raw_transaction

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type Transaction []byte

func (t Transaction) Serialize() ([]byte, error) {
	return t, nil
}

func (t Transaction) Bytes() ([]byte, error) {
	return t, nil
}


func (t Transaction) Signed() bool {
	return true
}

func (t Transaction) Sign(uiptypes.Signer) (uiptypes.RawTransaction, error) {
	return t, nil
}
