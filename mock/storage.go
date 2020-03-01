package mock

import (
	"github.com/HyperService-Consortium/go-uip/storage"
)

type MerkMap map[string][]byte

func (m MerkMap) Update(key []byte, value []byte) error {
	m[string(key)] = value
	return nil
}

func (m MerkMap) Get(key []byte) ([]byte, error) {
	return m[string(key)], nil
}

func (m MerkMap) Delete(key []byte) error {
	delete(m, string(key))
	return nil
}

func (m MerkMap) Prove(key []byte) ([][]byte, error) {
	panic("implement me")
}

func (m MerkMap) MakeProof(key []byte) string {
	panic("implement me")
}

func (m MerkMap) MakeErrorProof(err error) string {
	panic("implement me")
}

func (m MerkMap) MakeErrorProofFromString(str string) string {
	panic("implement me")
}

type LocalStorage struct {
	maps map[string]storage.MerkMap
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{make(map[string]storage.MerkMap)}
}

func (l *LocalStorage) ArrangeSlot(newSlot string) storage.MerkMap {
	if m, ok := l.maps[newSlot]; ok {
		return m
	}
	m := make(MerkMap)
	l.maps[newSlot] = m
	return m
}

func (l LocalStorage) Commit() {

}
