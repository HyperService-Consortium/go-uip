package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInitializer_initContractInvocationR(t *testing.T) {
	type args struct {
		info    *IntentImpl
		content document.Document
	}
	tests := []struct {
		name    string
		args    args
		wantSub interface{}
		wantErr bool
		errType string
	}{
		{"good", args{
			info: &IntentImpl{
				Name:   "op1",
				OpType: token.Invoke,
			},
			content: sugar.HandlerError(document.NewMapDocument(nil)).(document.Document),
		}, nil, true, errorn.ErrorTypeFieldNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ier := new(DocumentLexer)
			gotSub, err := ier.initContractInvocation(tt.args.info, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("initContractInvocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				pe := err.(*errorn.ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
				return
			}
			if !reflect.DeepEqual(gotSub, tt.wantSub) {
				t.Errorf("initContractInvocation() gotSub = %v, want %v", gotSub, tt.wantSub)
			}
		})
	}
}

//func Test_checkContractInvocationFields(t *testing.T) {
//	type args struct {
//		intent *token.InvokeIntent
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//		errType string
//	}{
//		{name: "good", args: args{intent: &token.InvokeIntent{
//			Src:      &token.RawAccountInfo{
//				ChainId: 0,
//				Name:    "",
//			},
//			Dst:      "11",
//			Code:     nil,
//			FuncName: "a",
//			Params:   []token.ParamImpl{},
//			Amount:   "0200",
//			Meta:     nil,
//		}}},
//		{name: "good with empty amount", args: args{intent: &token.InvokeIntent{
//			Src:      &token.RawAccountInfo{
//				ChainId: 0,
//				Name:    "",
//			},
//			Dst:      "11",
//			Code:     nil,
//			FuncName: "a",
//			Params:   []token.ParamImpl{},
//			Amount:   "",
//			Meta:     nil,
//		}}},
//		{"src absent", args{intent: &token.InvokeIntent{
//			Src:      nil,
//			Dst:      "11",
//			Code:     nil,
//			FuncName: "a",
//			Params:   []token.ParamImpl{},
//			Amount:   "0200",
//			Meta:     nil,
//		}}, true, errorn.ErrorTypeFieldNotFound},
//		{"dst invalid", args{intent: &token.InvokeIntent{
//			Src:      &token.RawAccountInfo{
//				ChainId: 0,
//				Name:    "",
//			},
//			Dst:      "",
//			Code:     nil,
//			FuncName: "a",
//			Params:   []token.ParamImpl{},
//			Amount:   "0200",
//			Meta:     nil,
//		}}, true, errorn.ErrorTypeInvalidField},
//		{"dst invalid", args{intent: &token.InvokeIntent{
//			Src:      &token.RawAccountInfo{
//				ChainId: 0,
//				Name:    "",
//			},
//			Dst:      "1",
//			Code:     nil,
//			FuncName: "a",
//			Params:   []token.ParamImpl{},
//			Amount:   "0200",
//			Meta:     nil,
//		}}, true, errorn.ErrorTypeInvalidField},
//		{"amount invalid", args{intent: &token.InvokeIntent{
//			Src:      &token.RawAccountInfo{
//				ChainId: 0,
//				Name:    "",
//			},
//			Dst:      "11",
//			Code:     nil,
//			FuncName: "a",
//			Params:   []token.ParamImpl{},
//			Amount:   "020",
//			Meta:     nil,
//		}}, true, errorn.ErrorTypeInvalidField},
//		{"fn invalid", args{intent: &token.InvokeIntent{
//			Src:      &token.RawAccountInfo{
//				ChainId: 0,
//				Name:    "",
//			},
//			Dst:      "11",
//			Code:     nil,
//			FuncName: "",
//			Params:   []token.ParamImpl{},
//			Amount:   "0020",
//			Meta:     nil,
//		}}, true, errorn.ErrorTypeFieldNotFound},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ier := new(DocumentLexer)
//			if _, err := ier.checkContractInvocation(tt.args.intent); (err != nil) != tt.wantErr {
//				t.Errorf("checkContractInvocationFields() error = %v, wantErr %v", err, tt.wantErr)
//			} else if err != nil {
//				pe := err.(*errorn.ParseError)
//				if pe.ErrType != tt.errType {
//					t.Errorf("err.t = %v, want = %v", pe, tt.errType)
//				}
//			}
//		})
//	}
//}

func TestDocumentLexer_initAccounts(t *testing.T) {
	type fields struct {
		BaseLexer BaseLexer
	}
	type args struct {
		nameKey string
		source  document.Document
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantAccounts []FullAccount
		wantErr      bool
	}{
		{"good", fields{BaseLexer: BaseLexer{}}, args{nameKey: "userName", source: sugar.HandlerError(
			document.NewMapDocument([]interface{}{document.MObj{
				"userName": "a1",
				"domain":   1,
				"address":  "0x3723",
			}})).(document.Document)},
			[]FullAccount{
				{Name: "a1", ChainID: 1, Address: []byte{0x37, 0x23}},
			}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DocumentLexer{
				BaseLexer: tt.fields.BaseLexer,
			}
			gotAccounts, err := l.initAccounts(tt.args.nameKey, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("initAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.wantAccounts, gotAccounts, ) {
				t.Errorf("initAccounts() gotAccounts = %v, want %v", gotAccounts, tt.wantAccounts)
			}
		})
	}
}

func TestBuildAccountMap(t *testing.T) {
	type args struct {
		accounts []FullAccount
	}
	tests := []struct {
		name    string
		args    args
		wantRes AccountMap
		wantErr bool
	}{
		{"good", args{[]FullAccount{
			{Name:"a1", ChainID: 1, Address: []byte{}},
		}}, AccountMap{
			"a1" : ChainMap{
				0: &FullAccount{Name:"a1", ChainID: 1, Address: []byte{}},
				1: &FullAccount{Name:"a1", ChainID: 1, Address: []byte{}},
			},
		}, false},
		{"good on diff chain", args{[]FullAccount{
			{Name:"a1", ChainID: 1, Address: []byte{}},
			{Name:"a1", ChainID: 2, Address: []byte{1}},
		}}, AccountMap{
			"a1" : ChainMap{
				1: &FullAccount{Name:"a1", ChainID: 1, Address: []byte{}},
				2: &FullAccount{Name:"a1", ChainID: 2, Address: []byte{1}},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := BuildAccountMap(tt.args.accounts)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildAccountMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.EqualValues(t, tt.wantRes, gotRes) {
				t.Errorf("BuildAccountMap() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}