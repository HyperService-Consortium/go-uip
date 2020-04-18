package serial

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"math/big"
	"strings"
	"testing"
)

func TestWriteBigInt(t *testing.T) {
	type args struct {
		i   *big.Int
		err *error
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{"normal", args{big.NewInt(1), new(error)},
			"\x00\x00\x00\x00\x00\x00\x00\x01" + string(big.NewInt(1).Bytes())},
		{"nil", args{nil, new(error)},
			"\x00\x00\x00\x00\x00\x00\x00\x00" + string(big.NewInt(1).SetBytes(nil).Bytes())},
		{"nil", args{big.NewInt(0).SetBytes(nil), new(error)},
			"\x00\x00\x00\x00\x00\x00\x00\x00" + string(big.NewInt(1).SetBytes(nil).Bytes())},
		{"nil", args{big.NewInt(1).SetBytes(nil), new(error)},
			"\x00\x00\x00\x00\x00\x00\x00\x00" + string(big.NewInt(1).SetBytes(nil).Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			WriteBigInt(w, tt.args.i, tt.args.err)
			if *tt.args.err != nil {
				t.Errorf("write error: %v", *tt.args.err)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("WriteBigInt() = %v, want %v", []byte(gotW), []byte(tt.wantW))
			}
		})
	}
}

func TestReadBigInt(t *testing.T) {
	type args struct {
		r   io.Reader
		i   *big.Int
		err *error
	}
	tests := []struct {
		name  string
		args  args
		wantI *big.Int
	}{
		{"normal",
			args{strings.NewReader("\x00\x00\x00\x00\x00\x00\x00\x01" + string(big.NewInt(1).Bytes())),
				big.NewInt(0), new(error)}, big.NewInt(1)},
		{"nil", args{strings.NewReader("\x00\x00\x00\x00\x00\x00\x00\x00" + string(big.NewInt(1).SetBytes(nil).Bytes())),
			big.NewInt(0), new(error)}, big.NewInt(0).SetBytes([]byte{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadBigInt(tt.args.r, tt.args.i, tt.args.err)
			if *tt.args.err != nil {
				t.Errorf("read error: %v", *tt.args.err)
			}
			assert.EqualValues(t, tt.wantI, tt.args.i)
		})
	}
}
