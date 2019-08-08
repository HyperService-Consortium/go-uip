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

func (bn *BlockChainInterfaceTest) WaitForTransact(chainID uint64, receipt []byte, opt *WaitOption) ([]byte, []byte, error) {
	return nil, nil, errors.New("must impl method WaitForTransact(cid, receipt, opt) (bid, info, err)")
}

func (bn *BlockChainInterfaceTest) GetTransactionProof(chainID uint64, blockID []byte, additional []byte) (MerkleProof, error) {
	return nil, errors.New("must impl method GetTransactionProof(cid, bid, additional) (info, err)")
}

// func (bn *BlockChainInterfaceTest) Route(intent *TransactionIntent, kvs map[string][]byte) ([]byte, error) {
// 	return Route(bn, intent, kvs)
// }

func (bn *BlockChainInterfaceTest) RouteRaw(uint64, []byte) ([]byte, error) {
	return nil, errors.New("must impl method RouteRaw(cid, rtx) (info, err)")
}

func (bn *BlockChainInterfaceTest) RouteRawTransaction(chainID, rawTransaction) (receipt, error) {
	return nil, errors.New("must impl method RouteRawTransaction(cid, rawTx) (receipt, error)")
}

func (bn *BlockChainInterfaceTest) Translate(*TransactionIntent, KVGetter) ([]byte, error) {
	return nil, errors.New("must impl method Translate(tx, kvs) (rtx, err)")
}

func (bn *BlockChainInterfaceTest) GetStorageAt(chainID, typeID, contract, pos, desc) (interface{}, error) {
	return nil, errors.New("must impl method GetStorageAt(chainID, typeID, contract, pos, desc) (interface{}, error)")
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

func (s *signature) Bytes() []byte {
	return []byte("666")
}
func (s *signature) String() string {
	return "666"
}
func (s *signature) FromBytes([]byte) bool {
	return true
}
func (s *signature) FromString(string) bool {
	return true
}

func (s *signature) Equal(HexType) bool {
	return true
}
func (s *signature) IsValid() bool {
	return true
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
