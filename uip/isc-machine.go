package uip

import "github.com/Myriad-Dreamin/gvm"

type ISCMachine interface {
	GetPC() uint64
	GetMuPC() uint64

	SetPC(pc uint64)
	SetMuPC(muPC uint64)

	GetExternalStorageAt(chainID ChainID, typeID TypeID,
		contractAddress ContractAddress, pos []byte, description []byte) (gvm.Ref, error)
}
