package opintent

import types "github.com/Myriad-Dreamin/go-uip/types"

type TransactionIntent = types.TransactionIntent
type MerkleProofProposal = types.MerkleProofProposal
type hexstring = string

// the LeftName intent is before RightName intent
type Dependency struct {
	Src uint32
	Dst uint32
}
