package opintent

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestRawParam_UnmarshalResult(t *testing.T) {
	type args struct {
		i ResultI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType string
		want RawParam
	}{
		{name: "type-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldOpIntentsValue: mObj{
				"constant": 2000,
			}})).(ResultI)}, wantErr: true, errType: ErrorTypeFieldNotFound},
		{name: "value-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldOpIntentsType: "t"})).(ResultI)}, wantErr: true, errType: ErrorTypeFieldNotFound},
		{name: "good", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldOpIntentsType: "t", FieldOpIntentsValue: mObj{
				"constant": 2000,
			}})).(ResultI)},
			wantErr: false, want: RawParam{
			Type: "t",
			Value: sugar.HandlerError(NewMapResult(mObj{
				"constant": 2000,
			})).(ResultI),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(RawParam)
			if err := r.UnmarshalResult(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalResult() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				pe := err.(*ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
			} else {
				if !reflect.DeepEqual(*r, tt.want) {
					t.Errorf("content = %v, want = %v", *r, tt.want)
				}
			}
		})
	}
}

func Test_initParamsR(t *testing.T) {
	type args struct {
		i ResultI
	}
	tests := []struct {
		name       string
		args       args
		wantParams []RawParam
		wantErr    bool
		errType string
	}{
		{name: "good", args: args{i: sugar.HandlerError(
			NewMapResult([]interface{}{mObj{FieldOpIntentsType: "t", FieldOpIntentsValue: mObj{
				"constant": 2000,
			}}})).(ResultI)},
			wantErr: false, wantParams: []RawParam{{
			Type: "t",
			Value: sugar.HandlerError(NewMapResult(mObj{
				"constant": 2000,
			})).(ResultI),
		}}},
		{name: "good-empty", args: args{i: sugar.HandlerError(
			NewMapResult([]interface{}{})).(ResultI)},
			wantErr: false, wantParams: []RawParam{}},
		{name: "invalid type", args: args{i: sugar.HandlerError(
			NewMapResult("")).(ResultI)},
			wantErr: true, errType: ErrorTypeInvalidField},
		{name: "invalid parameter", args: args{i: sugar.HandlerError(
			NewMapResult([]interface{}{mObj{FieldOpIntentsType: "t"}})).(ResultI)},
			wantErr: true, errType: ErrorTypeFieldNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParams, err := initParamsR(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("initParamsR() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				pe := err.(*ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
				return
			}
			if !reflect.DeepEqual(gotParams, tt.wantParams) {
				t.Errorf("initParamsR() gotParams = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}