package storage

type MerkMap interface {
	Update(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error

	Prove(key []byte) ([][]byte, error)
	MakeProof(key []byte) string
	MakeErrorProof(err error) string
	MakeErrorProofFromString(str string) string
	//Revert() (err error)
	//CommitN() (root []byte, err error)
}

type LocalStorage interface {
	ArrangeSlot(newSlot string) MerkMap
	Commit() error
}
