package opintent

import (
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestRawDependency_UnmarshalResult(t *testing.T) {
	type args struct {
		i ResultI
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want RawDependency
		errType string
	}{
		{name: "left-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldDependencyRight: "op2", FieldDependencyDep: "before"})).(ResultI)}, wantErr: true, errType: ErrorTypeFieldNotFound},
		{name: "right-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldDependencyLeft: "op1", FieldDependencyDep: "before"})).(ResultI)}, wantErr: true, errType: ErrorTypeFieldNotFound},
		{name: "dep-wrong", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldDependencyLeft: "op1", FieldDependencyRight: "op2",
				FieldDependencyDep: "<wrong-parsed>"})).(ResultI)}, wantErr: true, errType: ErrorTypeInvalidField},
		{name: "good-dep-absent", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldDependencyLeft: "op1", FieldDependencyRight: "op2"})).(ResultI)},
			wantErr: false, want: RawDependency{
				Src: "op1", Dst: "op2",
		}},
		{name: "good-before", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldDependencyLeft: "op1", FieldDependencyRight: "op2", FieldDependencyDep: "before"})).(ResultI)},
			wantErr: false, want: RawDependency{
			Src: "op1", Dst: "op2",
		}},
		{name: "good-after", args: args{i: sugar.HandlerError(
			NewMapResult(mObj{FieldDependencyLeft: "op1", FieldDependencyRight: "op2", FieldDependencyDep: "after"})).(ResultI)},
			wantErr: false, want: RawDependency{
			Src: "op2", Dst: "op1",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(RawDependency)
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