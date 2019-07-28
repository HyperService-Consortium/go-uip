package opintent

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
)

type _opIntents struct {
	Contents             [][]byte `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	Dependencies         [][]byte `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *_opIntents) Reset()         { *m = _opIntents{} }
func (m *_opIntents) String() string { return proto.CompactTextString(m) }
func (*_opIntents) ProtoMessage()    {}

func (m *_opIntents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo__opIntents.Unmarshal(m, b)
}
func (m *_opIntents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo__opIntents.Marshal(b, m, deterministic)
}
func (m *_opIntents) XXX_Merge(src proto.Message) {
	xxx_messageInfo__opIntents.Merge(m, src)
}
func (m *_opIntents) XXX_Size() int {
	return xxx_messageInfo__opIntents.Size(m)
}
func (m *_opIntents) XXX_DiscardUnknown() {
	xxx_messageInfo__opIntents.DiscardUnknown(m)
}

var xxx_messageInfo__opIntents proto.InternalMessageInfo

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
		"amount": "2e0",
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
		t.Error(fmt.Errorf("Marshal failed: %v", err))
		return
	}
	opintent["name"] = "Op2"
	opintent["amount"] = "3e0"
	var c []byte
	c, err = json.Marshal(opintent)
	if err != nil {
		t.Error(fmt.Errorf("Marshal failed: %v", err))
		return
	}
	var d []byte
	d, err = json.Marshal(dep)
	if err != nil {
		t.Error(fmt.Errorf("Marshal failed: %v", err))
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

	var intents []*TransactionIntent

	intents, _, err = NewOpIntentInitializer().InitOpIntent(opIntents)

	if err != nil {
		t.Error(err)
		return
	}

	for _, intent := range intents {
		fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
		fmt.Println()
	}
}

func TestGenerateTransactionIntent(t *testing.T) {
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
		"amount": "2e0",
		"unit":   "wei",
	}
	var dep = obj{
		"left":  "Op1",
		"right": "Op2",
		"dep":   "after",
	}

	var b []byte
	b, err = json.Marshal(opintent)
	if err != nil {
		t.Error(fmt.Errorf("Marshal failed: %v", err))
		return
	}
	opintent["name"] = "Op2"
	opintent["amount"] = "3e0"
	var c []byte
	c, err = json.Marshal(opintent)
	if err != nil {
		t.Error(fmt.Errorf("Marshal failed: %v", err))
		return
	}
	var d []byte
	d, err = json.Marshal(dep)
	if err != nil {
		t.Error(fmt.Errorf("Marshal failed: %v", err))
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

	var intents []*TransactionIntent

	intents, _, err = NewOpIntentInitializer().InitOpIntent(opIntents)
	if err != nil {
		t.Error(err)
		return
	}

	for _, intent := range intents {
		fmt.Println(hex.EncodeToString(intent.Src), hex.EncodeToString(intent.Dst), intent.Amt)
	}
}
