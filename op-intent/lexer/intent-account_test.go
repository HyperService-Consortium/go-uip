package lexer

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
	"github.com/HyperService-Consortium/go-uip/internal/lexer_types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"reflect"
	"testing"
)

func TestAccountUnmarshalResult(t *testing.T) {
	type args struct {
		i document.Document
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType string
		want    lexer_types.Account
	}{
		{name: "good identity string", args: args{i: sugar.HandlerError(
			document.NewMapDocument("a1")).(document.Document)}, wantErr: false, want: lexer_types.NameAccount{
			Name: "a1",
		}},
		{name: "good hex string", args: args{i: sugar.HandlerError(
			document.NewMapDocument("0xa1")).(document.Document)}, wantErr: false, want: lexer_types.RawAccount{
			Address: []byte{0xa1},
		}},
		{name: "good domain-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldOpIntentsUserName: "a1"})).(document.Document)}, wantErr: false, want: lexer_types.NameAccount{
			Name: "a1",
		}},
		// whom chain id must be implied
		{name: "good address only", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{"address": "0001"})).(document.Document)}, wantErr: false, want: lexer_types.RawAccount{
			Address: []byte{0, 1},
		}},
		{name: "good address with namespace", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{"address": "0001", FieldOpIntentsDomain: 1})).(document.Document)}, wantErr: false, want: lexer_types.NamespacedRawAccount{
			Address: []byte{0, 1},
			ChainID: 1,
		}},
		{name: "name-absent", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldOpIntentsDomain: 1})).(document.Document)}, wantErr: true, errType: errorn.ErrorTypeInvalidField},
		{name: "good", args: args{i: sugar.HandlerError(
			document.NewMapDocument(document.MObj{FieldOpIntentsUserName: "a1", FieldOpIntentsDomain: 1})).(document.Document)},
			wantErr: false, want: lexer_types.NamespacedNameAccount{
				ChainID: 1,
				Name:    "a1",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if a, err := AccountUnmarshalResult(FieldOpIntentsUserName, tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalDocument() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				pe := err.(*errorn.ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
			} else {
				if !reflect.DeepEqual(reflect.ValueOf(a).Elem().Interface(), tt.want) {
					t.Errorf("content = %v, want = %v", reflect.ValueOf(a).Elem().Interface(), tt.want)
				}
			}
		})
	}
}
