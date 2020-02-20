package opintent

import types "github.com/HyperService-Consortium/go-uip/uip"

type TransactionIntent = types.TransactionIntent
type MerkleProofProposal = types.MerkleProofProposal


type RawAccountInfo struct {
	ChainId uint64 `json:"domain"`
	Name    string `json:"user_name"`
}
