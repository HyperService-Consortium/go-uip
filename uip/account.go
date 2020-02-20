package uip

// an implementation of types.Account is uiprpc.Account from "github.com/HyperService-Consortium/go-ves/grpc"
type Account interface {
	GetChainId() ChainID
	GetAddress() Address
}
