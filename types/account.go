package types

type ChainId = uint64
type Address = []byte

// an implementation of types.Account is uiprpc.Account from "github.com/Myriad-Dreamin/go-ves/grpc"
type Account interface {
	GetChainId() ChainId
	GetAddress() Address
}

type Signature interface {
	GetSignatureType() uint32
	GetContent() []byte
}

type Signer interface {
	GetPublicKey() []byte
	Sign([]byte) []byte
}
