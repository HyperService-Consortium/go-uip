package types

import (
	"errors"
	"fmt"
)

type BlockChainInterfaceTest struct {
	BaseBlockChainInterface
}

func (bn *BlockChainInterfaceTest) MustWithSigner() bool {
	return true
}

func (bn *BlockChainInterfaceTest) RouteWithSigner(signer Signer) Router {
	var nbn = *bn
	nbn.Signer = signer
	return &nbn
}

// func (bn *BlockChainInterfaceTest) Route(intent *TransactionIntent, kvs map[string][]byte) ([]byte, error) {
// 	return Route(bn, intent, kvs)
// }

func (bn *BlockChainInterfaceTest) RouteRaw(uint64, []byte) ([]byte, error) {
	return nil, errors.New("must impl method RouteRaw(cid, rtx) (info, err)")
}

func (bn *BlockChainInterfaceTest) Translate(*TransactionIntent, map[string][]byte) ([]byte, error) {
	return nil, errors.New("must impl method Translate(tx, kvs) (rtx, err)")
}

var _ BlockChainInterface = &BlockChainInterfaceTest{}

type signer struct {
}

func (s signer) GetPublicKey() []byte {
	return nil
}

type signature struct {
}

func (s *signature) GetSignatureType() uint32 {
	return 0
}
func (s *signature) GetContent() []byte {
	return []byte("666")
}

func (s signer) Sign([]byte) Signature {
	return &signature{}
}

func init() {
	fmt.Println((&BlockChainInterfaceTest{}).MustWithSigner())

	var a BaseBlockChainInterface
	var b = &BlockChainInterfaceTest{BaseBlockChainInterface: a}
	// b.Signer = signer{}
	fmt.Println(&b.Signer)
	var ss Signer = signer{}
	fmt.Printf("%p", &ss)
	fmt.Println(&(b.RouteWithSigner(ss).(*BlockChainInterfaceTest).Signer))
	fmt.Println((b.RouteWithSigner(ss).(*BlockChainInterfaceTest).Signer.Sign([]byte(""))))
}