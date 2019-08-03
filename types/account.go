package types

// an implementation of types.Account is uiprpc.Account from "github.com/Myriad-Dreamin/go-ves/grpc"
type Account interface {
	GetChainId() chainID
	GetAddress() address
}

type signatureType = uint32
type signatureContent = []byte
type signedContent = []byte

type HexType interface {
	Bytes() []byte
	String() string
	FromBytes([]byte) bool
	FromString(string) bool
	Equal(HexType) bool
}

type BaseSignature interface {
	GetSignatureType() signatureType
	GetContent() signedContent
}

type Signature interface {
	HexType
	BaseSignature
	IsValid() bool
}

type publicKey []byte

type Signer interface {
	GetPublicKey() []byte
	Sign(signatureContent) Signature
}
