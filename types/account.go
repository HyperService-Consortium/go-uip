package types

type chain_id = uint64
type address = []byte

// an implementation of types.Account is uiprpc.Account from "github.com/Myriad-Dreamin/go-ves/grpc"
type Account interface {
	GetChainId() chain_id
	GetAddress() address
}

type Signature interface {
	GetSignatureType() uint32
	GetContent() []byte
}

type TenSigner interface {
	GetPublicKey() []byte
	Sign([]byte) []byte
}
