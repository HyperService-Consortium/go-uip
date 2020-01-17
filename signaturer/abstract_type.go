package signaturer

import (
	"bytes"
	"encoding/hex"
	
	"github.com/HyperService-Consortium/go-uip/uiptypes"
)

type ECCSignature = uiptypes.Signature

type ECCPublicKey interface {
	uiptypes.HexType
	IsValid() bool
}

type ECCPrivateKey interface {
	uiptypes.HexType
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
	return *h
}

func (h *BaseHexType) String() string {
	return hex.EncodeToString(*h)
}

func (h *BaseHexType) PureString() string {
	return string(*h)
}

func (h *BaseHexType) FromBytes(b []byte) error {
	if h == nil {
		h = new(BaseHexType)
	}
	*h = b
	return nil
}

func (h *BaseHexType) FromPureString(b string) error {
	if h == nil {
		h = new(BaseHexType)
	}
	*h = []byte(b)
	return nil
}

func (h *BaseHexType) FromString(b string) error {
	var err error
	if h == nil {
		h = new(BaseHexType)
	}
	*h, err = hex.DecodeString(b)
	return err
}

func (h *BaseHexType) Equal(rh uiptypes.HexType) bool {
	return bytes.Equal(*h, rh.Bytes())
}
