package merkleproof_proposal_type

import "github.com/HyperService-Consortium/go-uip/uip"

type Type = uip.MerkleProofProposalType

const (
	TransactionProof Type = iota
	DataProof
)
