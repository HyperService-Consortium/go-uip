package standard

import (
	"encoding/hex"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"testing"
)

func TestIsValidHexString(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name:"nil", args:args{h:""}, want: false},
		{name:"single", args:args{h:"1"}, want: false},
		{name:"length 3", args:args{h:"123"}, want: false},
		{name:"with prefix", args:args{h:"0x123"}, want: false},
		{name:"with prefix even length", args:args{h:"0xff"}, want: false},
		{name:"x", args:args{h:"x"}, want: false},
		{name:"ff", args:args{h:"ff"}, want: true},
		{name:"FF", args:args{h:"FF"}, want: true},
		{name:"Ff", args:args{h:"Ff"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidHexString(tt.args.h); got != tt.want {
				t.Errorf("IsValidHexString() = %v, want %v", got, tt.want)
			}

			if tt.want == true {
				sugar.HandlerError(hex.DecodeString(tt.args.h))
			}
		})
	}
}