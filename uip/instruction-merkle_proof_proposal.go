package uip

type MerkleProofProposalsImpl []MerkleProofProposal

func (m MerkleProofProposalsImpl) Len() int {
	return len(m)
}

func (m MerkleProofProposalsImpl) Cap() int {
	return cap(m)
}

func (m MerkleProofProposalsImpl) Index(i int) MerkleProofProposal {
	return m[i]
}

func (m MerkleProofProposalsImpl) Slice(l, r int) MerkleProofProposals {
	return m[l:r]
}

func (m MerkleProofProposalsImpl) Append(appends ...MerkleProofProposal) (T MerkleProofProposals) {
	return append(m, appends...)
}
