package opintent

import types "github.com/HyperService-Consortium/go-uip/uiptypes"

type TransactionIntent = types.TransactionIntent
type MerkleProofProposal = types.MerkleProofProposal
type hexstring = string

// the LeftName intent is before RightName intent
type Dependency struct {
	Src uint32
	Dst uint32
}
