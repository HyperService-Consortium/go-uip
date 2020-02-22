package opintent

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/trans_type"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestInitializer_initContractInvocation(t *testing.T) {
	type args struct {
		info    *RawIntent
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		wantSub interface{}
		wantErr bool
		errType string
	}{
		{"good", args{
			info: &RawIntent{
				BaseOpIntent: BaseOpIntent{
					Name:         "op1",
					OpTypeString: "ContractInvocation",
				},
				Sub:          nil,
				OpType:       trans_type.ContractInvoke,
			},
			content: sugar.HandlerError(json.Marshal(&BaseContractInvocationOpIntent{
				Src:      nil,
				Dst:      "",
				Code:     nil,
				FuncName: "",
				Params:   nil,
				Amount:   "",
				Meta:     nil,
			})).([]byte),
		}, nil, true, ErrorTypeUnmarshalError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ier := new(Initializer)
			gotSub, err := ier.initContractInvocation(tt.args.info, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("initContractInvocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				pe := err.(*ParseError)
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

func TestInitializer_initContractInvocationR(t *testing.T) {
	type args struct {
		info    *RawIntent
		content ResultI
	}
	tests := []struct {
		name    string
		args    args
		wantSub interface{}
		wantErr bool
		errType string
	}{
		{"good", args{
			info: &RawIntent{
				BaseOpIntent: BaseOpIntent{
					Name:         "op1",
					OpTypeString: "ContractInvocation",
				},
				Sub:          nil,
				OpType:       trans_type.ContractInvoke,
			},
			content: sugar.HandlerError(NewMapResult(nil)).(ResultI),
		}, nil, true, ErrorTypeFieldNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ier := new(Initializer)
			gotSub, err := ier.initContractInvocationR(tt.args.info, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("initContractInvocationR() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				pe := err.(*ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
				return
			}
			if !reflect.DeepEqual(gotSub, tt.wantSub) {
				t.Errorf("initContractInvocationR() gotSub = %v, want %v", gotSub, tt.wantSub)
			}
		})
	}
}

func Test_checkContractInvocationFields(t *testing.T) {
	type args struct {
		intent *BaseContractInvocationOpIntent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType string
	}{
		{name: "good", args: args{intent: &BaseContractInvocationOpIntent{
			Src:      &RawAccountInfo{
				ChainId: 0,
				Name:    "",
			},
			Dst:      "11",
			Code:     nil,
			FuncName: "a",
			Params:   []RawParam{},
			Amount:   "0200",
			Meta:     nil,
		}}},
		{name: "good with empty amount", args: args{intent: &BaseContractInvocationOpIntent{
			Src:      &RawAccountInfo{
				ChainId: 0,
				Name:    "",
			},
			Dst:      "11",
			Code:     nil,
			FuncName: "a",
			Params:   []RawParam{},
			Amount:   "",
			Meta:     nil,
		}}},
		{"src absent", args{intent: &BaseContractInvocationOpIntent{
			Src:      nil,
			Dst:      "11",
			Code:     nil,
			FuncName: "a",
			Params:   []RawParam{},
			Amount:   "0200",
			Meta:     nil,
		}}, true, ErrorTypeFieldNotFound},
		{"dst invalid", args{intent: &BaseContractInvocationOpIntent{
			Src:      &RawAccountInfo{
				ChainId: 0,
				Name:    "",
			},
			Dst:      "",
			Code:     nil,
			FuncName: "a",
			Params:   []RawParam{},
			Amount:   "0200",
			Meta:     nil,
		}}, true, ErrorTypeInvalidField},
		{"dst invalid", args{intent: &BaseContractInvocationOpIntent{
			Src:      &RawAccountInfo{
				ChainId: 0,
				Name:    "",
			},
			Dst:      "1",
			Code:     nil,
			FuncName: "a",
			Params:   []RawParam{},
			Amount:   "0200",
			Meta:     nil,
		}}, true, ErrorTypeInvalidField},
		{"amount invalid", args{intent: &BaseContractInvocationOpIntent{
			Src:      &RawAccountInfo{
				ChainId: 0,
				Name:    "",
			},
			Dst:      "11",
			Code:     nil,
			FuncName: "a",
			Params:   []RawParam{},
			Amount:   "020",
			Meta:     nil,
		}}, true, ErrorTypeInvalidField},
		{"fn invalid", args{intent: &BaseContractInvocationOpIntent{
			Src:      &RawAccountInfo{
				ChainId: 0,
				Name:    "",
			},
			Dst:      "11",
			Code:     nil,
			FuncName: "",
			Params:   []RawParam{},
			Amount:   "0020",
			Meta:     nil,
		}}, true, ErrorTypeFieldNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkContractInvocationFields(tt.args.intent); (err != nil) != tt.wantErr {
				t.Errorf("checkContractInvocationFields() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				pe := err.(*ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe, tt.errType)
				}
			}
		})
	}
}
