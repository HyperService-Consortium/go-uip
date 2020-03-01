package uip

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
)

type BlockChainInterfaceTest struct {
	BlockChainInterface

	Signer Signer
}

func (bn *BlockChainInterfaceTest) Deserialize(raw []byte) (rawTransaction RawTransaction, err error) {
	panic("implement me")
}

func (bn *BlockChainInterfaceTest) MustWithSigner() bool {
	return true
}

func (bn *BlockChainInterfaceTest) RouteWithSigner(signer Signer) (Router, error) {
	var nbn = *bn
	nbn.Signer = signer
	return &nbn, nil
}

func (bn *BlockChainInterfaceTest) WaitForTransact(chainID uint64, receipt []byte, opts ...interface{}) ([]byte, []byte, error) {
	return nil, nil, errors.New("must impl method WaitForTransact(cid, receipt, opt) (bid, info, err)")
}

func (bn *BlockChainInterfaceTest) GetTransactionProof(chainID uint64, blockID []byte, additional []byte) (MerkleProof, error) {
	return nil, errors.New("must impl method GetTransactionProof(cid, bid, additional) (info, err)")
}

// func (bn *BlockChainInterfaceTest) Route(intent *TransactionIntent, kvs map[string][]byte) ([]byte, error) {
// 	return Route(bn, intent, kvs)
// }

func (bn *BlockChainInterfaceTest) RouteRaw(ChainID, RawTransaction) (TransactionReceipt, error) {
	return nil, errors.New("must impl method RouteRaw(cid, rtx) (info, err)")
}


func (bn *BlockChainInterfaceTest) Translate(*parser.TransactionIntent, Storage) (RawTransaction, error) {
	return nil, errors.New("must impl method Translate(tx, kvs) (rtx, err)")
}

func (bn *BlockChainInterfaceTest) GetStorageAt(ChainID, TypeID, ContractAddress, []byte, []byte) (Variable, error) {
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

func (s *signature) GetSignatureType() SignatureType {
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
func (s *signature) FromBytes([]byte) error {
	return nil
}
func (s *signature) FromString(string) error {
	return nil
}

func (s *signature) Equal(HexType) bool {
	return true
}
func (s *signature) IsValid() bool {
	return true
}

func (s signer) Sign(op []byte, options ...interface{}) (Signature, error) {
	return &signature{}, nil
}

func init() {
	fmt.Println((&BlockChainInterfaceTest{}).MustWithSigner())

	var a BlockChainInterface
	var b = &BlockChainInterfaceTest{BlockChainInterface: a}
	// b.Signer = signer{}
	fmt.Println(&b.Signer)
	var ss Signer = signer{}
	fmt.Printf("%p", &ss)
	fmt.Println((ss.Sign([]byte(""))))
}
