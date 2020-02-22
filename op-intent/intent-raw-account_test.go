package opintent

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestRawAccountInfo_UnmarshalResult(t *testing.T) {
	type args struct {
		i ResultI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType string
		want RawAccountInfo
	}{
		{name: "domain-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldOpIntentsUserName: "a1"})).(ResultI)}, wantErr: true, errType: ErrorTypeFieldNotFound},
		{name: "name-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldOpIntentsDomain: 1})).(ResultI)}, wantErr: true, errType: ErrorTypeFieldNotFound},
		{name: "good", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldOpIntentsUserName: "a1", FieldOpIntentsDomain: 1})).(ResultI)},
			wantErr: false, want: RawAccountInfo{
			ChainId: 1,
			Name:    "a1",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(RawAccountInfo)
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
