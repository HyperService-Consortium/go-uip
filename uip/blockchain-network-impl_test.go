package uip_test

import (
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type BlockChainInterfaceTest struct {
	uip.BlockChainInterface

	Signer uip.Signer
}

func (bn *BlockChainInterfaceTest) Deserialize(raw []byte) (rawTransaction uip.RawTransaction, err error) {
	panic("implement me")
}

func (bn *BlockChainInterfaceTest) MustWithSigner() bool {
	return true
}

func (bn *BlockChainInterfaceTest) RouteWithSigner(signer uip.Signer) (uip.Router, error) {
	var nbn = *bn
	nbn.Signer = signer
	return &nbn, nil
}

func (bn *BlockChainInterfaceTest) WaitForTransact(chainID uint64, receipt []byte, opts ...interface{}) ([]byte, []byte, error) {
	return nil, nil, errors.New("must impl method WaitForTransact(cid, receipt, opt) (bid, info, err)")
}

func (bn *BlockChainInterfaceTest) GetTransactionProof(chainID uint64, blockID []byte, additional []byte) (uip.MerkleProof, error) {
	return nil, errors.New("must impl method GetTransactionProof(cid, bid, additional) (info, err)")
}

// func (bn *BlockChainInterfaceTest) Route(intent *TransactionIntent, kvs map[string][]byte) ([]byte, error) {
// 	return Route(bn, intent, kvs)
// }

func (bn *BlockChainInterfaceTest) RouteRaw(uip.ChainID, uip.RawTransaction) (uip.TransactionReceipt, error) {
	return nil, errors.New("must impl method RouteRaw(cid, rtx) (info, err)")
}

func (bn *BlockChainInterfaceTest) Translate(uip.TransactionIntent, uip.Storage) (uip.RawTransaction, error) {
	return nil, errors.New("must impl method Translate(tx, kvs) (rtx, err)")
}

func (bn *BlockChainInterfaceTest) GetStorageAt(uip.ChainID, uip.TypeID, uip.ContractAddress, []byte, []byte) (uip.Variable, error) {
	return nil, errors.New("must impl method GetStorageAt(chainID, typeID, contract, pos, desc) (interface{}, error)")
}

var _ uip.BlockChainInterface = &BlockChainInterfaceTest{}

type signer struct {
}

func (s signer) GetPublicKey() []byte {
	return nil
}

type signature struct {
}

func (s *signature) GetSignatureType() uip.SignatureType {
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

func (s *signature) Equal(uip.HexType) bool {
	return true
}
func (s *signature) IsValid() bool {
	return true
}

func (s signer) Sign(op []byte, options ...interface{}) (uip.Signature, error) {
	return &signature{}, nil
}

func init() {
	fmt.Println((&BlockChainInterfaceTest{}).MustWithSigner())

	var a uip.BlockChainInterface
	var b = &BlockChainInterfaceTest{BlockChainInterface: a}
	// b.Signer = signer{}
	fmt.Println(&b.Signer)
	var ss uip.Signer = signer{}
	fmt.Printf("%p", &ss)
	fmt.Println((ss.Sign([]byte(""))))
}
