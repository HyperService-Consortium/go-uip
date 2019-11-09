package signaturer

import (
	"fmt"
	"testing"

	types "github.com/HyperService-Consortium/go-uip/uiptypes"
)

var _ ECCPublicKey = new(Ed25519PublicKey)
var _ ECCPrivateKey = NewEd25519PrivateKeyFromBytes([]byte{0, 1})

var _ ECCSignature = new(Ed25519Signature)

var _ ECCSignaturer = new(Ed25519Signaturer)

var _ types.Signature = new(Signature)


func TestSignature(t *testing.T) {
	fmt.Println(FromBaseSignature(NewEd25519PrivateKeyFromBytes([]byte{
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
	}).Sign([]byte("orz"))))
}

func TestNewTendermintNSBSigner(t *testing.T) {
	var _s, err = NewTendermintNSBSigner([]byte{
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	var _ = _s.Sign([]byte{0, 1})
}