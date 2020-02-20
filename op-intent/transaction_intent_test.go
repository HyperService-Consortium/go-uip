package opintent

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	merkle_proof "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

type _opIntents struct {
	Contents             [][]byte `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	Dependencies         [][]byte `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (m *_opIntents) Reset()         { *m = _opIntents{} }

func (m *_opIntents) GetContents() [][]byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *_opIntents) GetDependencies() [][]byte {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func TestGenerateTransactionIntentRev(t *testing.T) {
	type obj map[string]interface{}
	var err error
	var opintent = obj{
		"name":    "Op1",
		"op_type": "Payment",
		"src": obj{
			"domain":    2,
			"user_name": "a1",
		},
		"dst": obj{
			"domain":    1,
			"user_name": "a2",
		},
		"amount": "02e0",
		"unit":   "wei",
	}
	var dep = obj{
		"left":  "Op1",
		"right": "Op2",
		"dep":   "before",
	}

	var b []byte
	b, err = json.Marshal(opintent)
	if err != nil {
		t.Error(fmt.Errorf("marshal failed: %v", err))
		return
	}
	opintent["name"] = "Op2"
	opintent["amount"] = "03e0"
	var c []byte
	c, err = json.Marshal(opintent)
	if err != nil {
		t.Error(fmt.Errorf("marshal failed: %v", err))
		return
	}
	var d []byte
	d, err = json.Marshal(dep)
	if err != nil {
		t.Error(fmt.Errorf("marshal failed: %v", err))
		return
	}

	var opIntents = &_opIntents{
		Dependencies: [][]byte{
			d,
		},
		Contents: [][]byte{
			b,
			c,
		},
	}

	var intents TxIntents

	ier, err := NewInitializer(uip.BlockChainGetterNilImpl{}, mAccountProvider{})
	if err != nil {
		t.Error(err)
		return
	}

	intents, err = ier.Parse(opIntents)
	if err != nil {
		t.Error(err)
		pe := err.(*ParseError)
		fmt.Println(string(sugar.HandlerError(pe.Serialize()).([]byte)))
		return
	}

	for _, intent := range intents.GetTxIntents() {
		intent := intent.GetIntent()
		fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
		fmt.Println()
	}
}

func TestGenerateTransactionIntent(t *testing.T) {
	//type obj map[string]interface{}
	//var err error
	//var opintent = obj{
	//	"name":    "Op1",
	//	"op_type": "Payment",
	//	"src": obj{
	//		"domain":    2,
	//		"user_name": "a1",
	//	},
	//	"dst": obj{
	//		"domain":    1,
	//		"user_name": "a2",
	//	},
	//	"amount": "2e0",
	//	"unit":   "wei",
	//}
	//var dep = obj{
	//	"left":  "Op1",
	//	"right": "Op2",
	//	"dep":   "after",
	//}
	//
	//var b []byte
	//b, err = json.marshal(opintent)
	//if err != nil {
	//	t.Error(fmt.Errorf("marshal failed: %v", err))
	//	return
	//}
	//opintent["name"] = "Op2"
	//opintent["amount"] = "3e0"
	//var c []byte
	//c, err = json.marshal(opintent)
	//if err != nil {
	//	t.Error(fmt.Errorf("marshal failed: %v", err))
	//	return
	//}
	//var d []byte
	//d, err = json.marshal(dep)
	//if err != nil {
	//	t.Error(fmt.Errorf("marshal failed: %v", err))
	//	return
	//}
	//
	//var opIntents = &_opIntents{
	//	Dependencies: [][]byte{
	//		d,
	//	},
	//	Contents: [][]byte{
	//		b,
	//		c,
	//	},
	//}
	//
	//var intents []*TransactionIntent
	//
	//intents, _, err = NewInitializer().InitOpIntent(opIntents)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//for _, intent := range intents {
	//	fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
	//}
}

/*

	Src      *RawAccountInfo `json:"invoker"`       // key
	Dst      []byte          `json:"contract_addr"` // key
	Code     []byte          `json:"contract_code"` // key
	FuncName string          `json:"func"`
	Params   []RawParams     `json:"parameters"`
	Meta     []byte          `json:"meta"`
*/

type mAccountProvider struct {
}

func (a mAccountProvider) AccountBase() uip.AccountBase {
	return a
}

func (mAccountProvider) Get(name string, chainId uint64) (uip.Account, error) {
	return &uip.AccountImpl{
		ChainId: chainId,
		Address: []byte("121313212313133123333333333333333313"),
	}, nil
}

func (mAccountProvider) GetRelay(domain uint64) (uip.Account, error) {
	return &uip.AccountImpl{
		ChainId: domain,
		Address: []byte("99999"),
	}, nil
}

func (mAccountProvider) GetTransactionProofType(chainId uint64) (uip.MerkleProofType, error) {
	return merkle_proof.MerklePatriciaTrieUsingKeccak256, nil
}

func TestGenerateInvokeTransactionIntent(t *testing.T) {
	//type obj map[string]interface{}
	//var err error
	//var opintent = obj{
	//	"name":    "Op1",
	//	"op_type": "ContractInvocation",
	//	"invoker": obj{
	//		"domain":    2,
	//		"user_name": "a1",
	//	},
	//	"contract_addr": "00e1eaa022cc40d4808bfe62b8997540c914d81e",
	//	"func":          "updateStake",
	//	"parameters": []obj{
	//		{
	//			"type": "uint256",
	//			"value": obj{
	//				"constant": "1000",
	//			},
	//		},
	//	},
	//	"amount": "0",
	//	"unit":   "wei",
	//}
	//
	//var b []byte
	//b, err = json.marshal(opintent)
	//if err != nil {
	//	t.Error(fmt.Errorf("marshal failed: %v", err))
	//	return
	//}
	//
	//var opIntents = &_opIntents{
	//	Contents: [][]byte{
	//		b,
	//	},
	//}
	//
	//var intents []*TransactionIntent
	//
	//intents, _, err = NewInitializer(mAccountProvider{}).InitOpIntent(opIntents)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//for _, intent := range intents {
	//	fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
	//	fmt.Println(intent.Meta)
	//	fmt.Println(string(intent.Meta))
	//	fmt.Println(intent.Amt)
	//}
}

func TestGeneratePaymentMetaTransactionIntent(t *testing.T) {
	//type obj map[string]interface{}
	//var err error
	//var opintent = obj{
	//	"name":    "Op1",
	//	"op_type": "Payment",
	//	"src": obj{
	//		"domain":    2,
	//		"user_name": "a1",
	//	},
	//	"dst": obj{
	//		"domain":    1,
	//		"user_name": "a2",
	//	},
	//	"meta": obj{
	//		"value-inconsistent": obj{
	//			"type": "uint256",
	//			"value": obj{
	//				"domain":   2,
	//				"contract": "00e1eaa022cc40d4808bfe62b8997540c914d81e",
	//				"pos":      "01",
	//				"field":    "strikePrice",
	//			},
	//		},
	//	},
	//	"amount": "2e0",
	//	"unit":   "wei",
	//}
	//
	//var b []byte
	//b, err = json.marshal(opintent)
	//if err != nil {
	//	t.Error(fmt.Errorf("marshal failed: %v", err))
	//	return
	//}
	//
	//var opIntents = &_opIntents{
	//	Contents: [][]byte{
	//		b,
	//	},
	//}
	//
	//var intents []*TransactionIntent
	//
	//intents, _, err = NewInitializer(mAccountProvider{}).InitOpIntent(opIntents)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//for _, intent := range intents {
	//	fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
	//	fmt.Println(intent.Meta)
	//	fmt.Println(string(intent.Meta))
	//	fmt.Println(intent.Amt)
	//}
}
