package merkleproof_proposal_type

import "github.com/HyperService-Consortium/go-uip/uiptypes"

type Type = uiptypes.MerkleProofProposalType

const (
	TransactionProof Type = iota
	DataProof
)
