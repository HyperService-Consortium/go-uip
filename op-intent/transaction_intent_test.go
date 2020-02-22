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
	Contents     [][]byte `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	Dependencies [][]byte `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (m *_opIntents) Reset() { *m = _opIntents{} }

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

func TestGenerateInconsistentTransactionIntent(t *testing.T) {
	type obj map[string]interface{}
	var err error
	var opintent = obj{
		"name":    "Op1",
		"op_type": "ContractInvocation",
		"invoker": obj{
			"domain":    6,
			"user_name": "a1",
		},
		"contract_addr": "263fef3fe76fd4075ac16271d5115d01206d3674",
		"func":          "updateStake",
		"parameters": []obj{
			{
				"type": "uint256",
				"value": obj{
					"constant": "1050",
				},
			},
		},
		"amount": "00",
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

	opintent = obj{
		"name":    "Op2",
		"op_type": "Payment",
		"src": obj{
			"domain":    1,
			"user_name": "a2",
		},
		"dst": obj{
			"domain":    6,
			"user_name": "a2",
		},
		"meta": obj{
			"value-inconsistent": obj{
				"type": "uint256",
				"value": obj{
					"domain":   6,
					"contract": "263fef3fe76fd4075ac16271d5115d01206d3674",
					"pos":      "01",
					"field":    "strikePrice",
				},
			},
		},
		"amount": "02e0",
		"unit":   "wei",
	}

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
		intent, proposals := intent.GetIntent(), intent.GetProposals()
		fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
		fmt.Println(string(intent.Meta))
		fmt.Println(intent.ChainID, intent.TransType)
		for _, proposal := range proposals {
			fmt.Println("qwq..................")
			fmt.Println(proposal.ValueType, proposal.Tid, proposal.DescriptionType, proposal.MerkleProofType)
			fmt.Println(string(proposal.SourceDescription))
			fmt.Println("qwq..................")
		}
		fmt.Println("qwq.....................")
	}
}



func genIntentsR(t testing.TB) map[string]interface{} {
	t.Helper()
	type obj = map[string]interface{}
	var opintent = obj{
		"name":    "Op1",
		"op_type": "ContractInvocation",
		"invoker": obj{
			"domain":    6,
			"user_name": "a1",
		},
		"contract_addr": "263fef3fe76fd4075ac16271d5115d01206d3674",
		"func":          "updateStake",
		"parameters": []obj{
			{
				"type": "uint256",
				"value": obj{
					"constant": "1050",
				},
			},
		},
		"amount": "00",
		"unit":   "wei",
	}
	var dep = obj{
		"left":  "Op1",
		"right": "Op2",
		"dep":   "before",
	}

	var opintent2 = obj{
		"name":    "Op2",
		"op_type": "Payment",
		"src": obj{
			"domain":    1,
			"user_name": "a2",
		},
		"dst": obj{
			"domain":    6,
			"user_name": "a2",
		},
		"meta": obj{
			"value-inconsistent": obj{
				"type": "uint256",
				"value": obj{
					"domain":   6,
					"contract": "263fef3fe76fd4075ac16271d5115d01206d3674",
					"pos":      "01",
					"field":    "strikePrice",
				},
			},
		},
		"amount": "02e0",
		"unit":   "wei",
	}

	return obj{
		FieldOpIntents: []obj {
			opintent,
			opintent2,
		},
		FieldOpIntentsDependencies: []obj {
			dep,
		},
	}

}

type packet struct {
	content []byte
	ct string
}

func (p packet) GetContent() (content []byte) {
	return p.content
}

func (p packet) GetContentType() string {
	return p.ct
}

func TestGenerateInconsistentTransactionIntentR(t *testing.T) {
	opIntents := genIntentsR(t)
	if t.Failed() {
		return
	}

	var intents TxIntents

	ier, err := NewInitializer(uip.BlockChainGetterNilImpl{}, mAccountProvider{})
	if err != nil {
		t.Error(err)
		return
	}

	p := packet{
		content: sugar.HandlerError(json.Marshal(opIntents)).([]byte),
		ct:      "json",
	}

	intents, err = ier.ParseR(p)
	if err != nil {
		t.Error(err)
		pe := err.(*ParseError)
		fmt.Println(string(sugar.HandlerError(pe.Serialize()).([]byte)))
		return
	}

	for _, intent := range intents.GetTxIntents() {
		intent, proposals := intent.GetIntent(), intent.GetProposals()
		fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
		fmt.Println(string(intent.Meta))
		fmt.Println(intent.ChainID, intent.TransType)
		for _, proposal := range proposals {
			fmt.Println("qwq..................")
			fmt.Println(proposal.ValueType, proposal.Tid, proposal.DescriptionType, proposal.MerkleProofType)
			fmt.Println(string(proposal.SourceDescription))
			fmt.Println("qwq..................")
		}
		fmt.Println("qwq.....................")
	}
}

//GOROOT=/home/kamiyoru/work/go #gosetup
//GOPATH=/home/kamiyoru/go #gosetup
///home/kamiyoru/work/go/bin/go test -c -o /tmp/___BenchmarkGenerateInconsistentTransactionIntent_in_github_com_HyperService_Consortium_go_uip_op_intent github.com/HyperService-Consortium/go-uip/op-intent #gosetup
///tmp/___BenchmarkGenerateInconsistentTransactionIntent_in_github_com_HyperService_Consortium_go_uip_op_intent -test.v -test.bench ^BenchmarkGenerateInconsistentTransactionIntent$ -test.run ^$ -test.cpuprofile /home/kamiyoru/.GoLand2019.2/system/profiler/go-uip.893f4101/op-intent'11 #gosetup
//goos: linux
//goarch: amd64
//pkg: github.com/HyperService-Consortium/go-uip/op-intent
//BenchmarkGenerateInconsistentTransactionIntent-4   	   31215	     39775 ns/op	    3856 B/op	      84 allocs/op
//BenchmarkGenerateInconsistentTransactionIntent-4   	   39921	     26840 ns/op	    5803 B/op	      87 allocs/op
// PASS
//
//Process finished with exit code 0
func BenchmarkGenerateInconsistentTransactionIntentR(b *testing.B) {
	opIntents := genIntentsR(b)
	if b.Failed() {
		return
	}

	ier, err := NewInitializer(uip.BlockChainGetterNilImpl{}, mAccountProvider{})
	if err != nil {
		b.Error(err)
		return
	}

	p := packet{
		content: sugar.HandlerError(json.Marshal(opIntents)).([]byte),
		ct:      "json",
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err = ier.ParseR(p)
		if err != nil {
			b.Error(err)
			pe := err.(*ParseError)
			fmt.Println(string(sugar.HandlerError(pe.Serialize()).([]byte)))
			return
		}
	}
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
