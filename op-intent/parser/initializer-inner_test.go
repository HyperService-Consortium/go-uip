package parser

import (
	"errors"
	"github.com/HyperService-Consortium/go-uip/errorn"
	"reflect"
	"testing"
)

func TestInitializer_decodeHex(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr bool
		errType string
	}{
		{name: "good", args: args{"00"}, wantB: []byte{0}},
		{name: "failed", args: args{"0"}, wantB: nil, wantErr: true, errType: errorn.ErrorTypeUnmarshalError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ier := &Parser{}
			gotB, err := ier.decodeHex(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				pe := err.(*errorn.ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("decodeHex() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

type errorT struct {
}

func (p errorT) MarshalJSON() ([]byte, error) {
	return nil, errors.New("marshal error")
}

func TestInitializer_marshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantB   []byte
		wantErr bool
		errType string
	}{
		{name: "good", args: args{"00"}, wantB: []byte(`"00"`)},
		{name: "failed", args: args{errorT{}}, wantB: nil, wantErr: true, errType: errorn.ErrorTypeMarshalError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ier := &Parser{}
			gotB, err := ier.marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if err != nil {
				pe := err.(*errorn.ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("marshal() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
