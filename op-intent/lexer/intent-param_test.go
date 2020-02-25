package lexer

import (
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//
//{
//"Type": "uint",
//"Value": {
//"contract": "c2",
//"field": "totalVotes",
//"pos": "1"
//}
//}
func TestParamUnmarshalResult(t *testing.T) {
	type args struct {
		i document.Document
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType string
		want    Param
	}{
		{name: "type-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldOpIntentsValue: document.MObj{
				"constant": 2000,
			}})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeFieldNotFound},
		{name: "value-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldOpIntentsType: "uint256"})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeFieldNotFound},
		{name: "type error", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint25",
				FieldOpIntentsValue: document.MObj{
					"constant": 2000,
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeValueTypeNotFound},
		{name: "good", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"constant": 2000,
				},
			})).(document.Document)}, want: &ConstantVariable{
			Type:  value_type.Uint256,
			Const: 2000,
		}, wantErr: false},
		{name: "good state", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"field":    "total",
					"pos":      "01",
				},
			})).(document.Document)}, want: &StateVariable{
			Type:  value_type.Uint256,
			Contract:  &NameAccount{Name: "c1"},
			Pos:   []byte{1},
			Field: []byte("total"),
		}, wantErr: false},
		{name: "state without pos", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"field":    "total",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "state with bad pos", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"field":    "total",
					"pos":      "1",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeDecodeContractPosError},
		{name: "state without field", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"pos":      "01",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "state without contract name", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"field": "total",
					"pos":   "01",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "good namespaced state", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"domain":   1,
					"field":    "total",
					"pos":      "01",
				},
			})).(document.Document)}, want: &StateVariable{
			Type: value_type.Uint256,
			Contract: &NamespacedAccount{
				ChainID: uint64(1),
				Name:    "c1",
			},
			Pos:   []byte{1},
			Field: []byte("total"),
		}, wantErr: false},
		{name: "namespaced state without contract name", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"domain": 1,
					"field":  "total",
					"pos":    "01",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "namespaced state without pos", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"domain":   1,
					"field":    "total",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "namespaced state without field", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{
				FieldOpIntentsType: "uint256",
				FieldOpIntentsValue: document.MObj{
					"contract": "c1",
					"domain":   1,
					"pos":      "01",
				},
			})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if p, err := ParamUnmarshalResult(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalDocument() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				pe := err.(*errorn.ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
			} else {
				if !assert.EqualValues(t, tt.want, p) {
					t.Errorf("content = %v, want = %v", reflect.ValueOf(p).Elem().Interface(), tt.want)
				}
			}
		})
	}
}

//func Test_initParamsR(t *testing.T) {
//	type args struct {
//		i document.Document
//	}
//	tests := []struct {
//		name       string
//		args       args
//		wantParams []ParamImpl
//		wantErr    bool
//		errType string
//	}{
//		{name: "good", args: args{i: sugar.HandlerError(
//			document.NewMapDocument([]interface{}{document.MObj{FieldOpIntentsType: "t", FieldOpIntentsValue: document.MObj{
//				"constant": 2000,
//			}}})).(document.Document)},
//			wantErr: false, wantParams: []ParamImpl{{
//			Type: "t",
//			Value: sugar.HandlerError(document.NewMapDocument(document.MObj{
//				"constant": 2000,
//			})).(document.Document),
//		}}},
//		{name: "good-empty", args: args{i: sugar.HandlerError(
//			document.NewMapDocument([]interface{}{})).(document.Document)},
//			wantErr: false, wantParams: []ParamImpl{}},
//		{name: "invalid type", args: args{i: sugar.HandlerError(
//			document.NewMapDocument("")).(document.Document)},
//			wantErr: true, errType: errorn.ErrorTypeInvalidField},
//		{name: "invalid parameter", args: args{i: sugar.HandlerError(
//			document.NewMapDocument([]interface{}{document.MObj{FieldOpIntentsType: "t"}})).(document.Document)},
//			wantErr: true, errType: errorn.ErrorTypeFieldNotFound},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			gotParams, err := initParamsR(tt.args.i)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("initParamsR() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			} else if err != nil {
//				pe := err.(*errorn.ParseError)
//				if pe.ErrType != tt.errType {
//					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
//				}
//				return
//			}
//			if !reflect.DeepEqual(gotParams, tt.wantParams) {
//				t.Errorf("initParamsR() gotParams = %v, want %v", gotParams, tt.wantParams)
//			}
//		})
//	}
//}
