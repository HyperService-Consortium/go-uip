package libgvm

import (
	"fmt"
	"github.com/HyperService-Consortium/go-uip/isc/gvm/internal/abstraction"
	"testing"
)

type ref struct {
}

func (r ref) GetGVMType() abstraction.RefType {
	return 2777777777
}

func (r ref) Unwrap() interface{} {
	return r
}

func (r ref) Encode() ([]byte, error) {
	return nil, nil
}

func TestConvertError(t *testing.T) {
	type args struct {
		l abstraction.Ref
		r abstraction.Ref
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{"easy", args{Uint8(1), Uint16(1)}, fmt.Errorf("cant convert ref-uint8 and ref-uint16 to the same type")},
		{"u-explain", args{Uint8(1), ref{}}, fmt.Errorf("cant convert ref-uint8 and ref-type(2777777777) to the same type")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertError(tt.args.l, tt.args.r); err == nil || err.Error() != tt.err.Error() {
				t.Errorf("ConvertError() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
