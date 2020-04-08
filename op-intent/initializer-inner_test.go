package opintent

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"testing"
)

func TestInitializer_unmarshal(t *testing.T) {
	type args struct {
		b []byte
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errType string
	}{
		{name: "good", args: args{[]byte("0"), new(int)}},
		{name: "failed", args: args{[]byte("{"), nil}, wantErr: true, errType: errorn.ErrorTypeUnmarshalError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ier := &Initializer{}
			if err := ier.unmarshal(tt.args.b, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				pe := err.(*errorn.ParseError)
				if pe.ErrType != tt.errType {
					t.Errorf("err.t = %v, want = %v", pe.ErrType, tt.errType)
				}
				return
			}
		})
	}
}
