package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/token"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
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
				Name:         "op1",
				OpType:       token.Invoke,
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
