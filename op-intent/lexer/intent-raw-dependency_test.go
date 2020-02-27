package lexer

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestRawDependency_UnmarshalResult(t *testing.T) {
	type args struct {
		i document.Document
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    RawDependency
		errType string
	}{
		{name: "left-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldKeyRight: "op2", FieldDependencyDep: "before"})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeFieldNotFound},
		{name: "right-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldKeyLeft: "op1", FieldDependencyDep: "before"})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeFieldNotFound},
		{name: "dep-wrong", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldKeyLeft: "op1", FieldKeyRight: "op2",
				FieldDependencyDep: "<wrong-parsed>"})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "good-dep-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldKeyLeft: "op1", FieldKeyRight: "op2"})).(document.Document)},
			wantErr: false, want: RawDependency{
				Src: "op1", Dst: "op2",
		}},
		{name: "good-before", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldKeyLeft: "op1", FieldKeyRight: "op2", FieldDependencyDep: "before"})).(document.Document)},
			wantErr: false, want: RawDependency{
			Src: "op1", Dst: "op2",
		}},
		{name: "good-after", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldKeyLeft: "op1", FieldKeyRight: "op2", FieldDependencyDep: "after"})).(document.Document)},
			wantErr: false, want: RawDependency{
			Src: "op2", Dst: "op1",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := new(RawDependency)
			if err := r.UnmarshalResult(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalDocument() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				pe := err.(*errorn.ParseError)
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