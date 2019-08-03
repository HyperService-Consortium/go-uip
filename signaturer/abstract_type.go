package signaturer

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/Myriad-Dreamin/go-uip/types"
)

type ECCSignature interface {
	types.Signature
}

type ECCPublicKey interface {
	types.HexType
	IsValid() bool
}

type ECCPrivateKey interface {
	types.HexType
	ToPublic() ECCPublicKey
	Sign([]byte) ECCSignature
}

type ECCSignaturer interface {
	Verify([]byte, []byte, []byte) bool
	Sign([]byte, []byte) ECCSignature
}

type BaseHexType []byte

func NewBaseHexTypeFromBytes(b []byte) (bh *BaseHexType) {
	bh = new(BaseHexType)
	*bh = b
	return
}

func NewBaseHexTypeFromPureString(b string) (bh *BaseHexType) {
	bh = new(BaseHexType)
	*bh = []byte(b)
	return
}

func NewBaseHexTypeFromString(b string) (bh *BaseHexType) {
	bod, err := hex.DecodeString(b)
	if err != nil {
		return nil
	}
	bh = new(BaseHexType)
	*bh = bod
	return
}

func (h *BaseHexType) Bytes() []byte {
	fmt.Println(h)
	return []byte(*h)
}

func (h *BaseHexType) String() string {
	return hex.EncodeToString(*h)
}

func (h *BaseHexType) PureString() string {
	return string(*h)
}

func (h *BaseHexType) FromBytes(b []byte) bool {
	if h == nil {
		h = new(BaseHexType)
	}
	*h = b
	return true
}

func (h *BaseHexType) FromPureString(b string) bool {
	if h == nil {
		h = new(BaseHexType)
	}
	*h = []byte(b)
	return true
}

func (h *BaseHexType) FromString(b string) bool {
	var err error
	if h == nil {
		h = new(BaseHexType)
	}
	*h, err = hex.DecodeString(b)
	return err != nil
}
func (h *BaseHexType) Equal(rh types.HexType) bool {
	return bytes.Equal(*h, rh.Bytes())
}
