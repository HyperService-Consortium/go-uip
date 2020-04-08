package isc

import "github.com/HyperService-Consortium/go-uip/storage"

const (
	state = "State"
)

type Storage struct {
	storage *storage.VM
}

func (s *Storage) GetISCState() uint8 {
	return s.storage.GetUint8(state)
}

func (s *Storage) SetISCState(u uint8) {
	s.storage.SetUint8(state, u)
}

func (s *Storage) GetFrozenInfoCount() uint64 {
	return s.storage.GetUint64("FrozenInfoCount")
}

func (s *Storage) SetFrozenInfoCount(u uint64) {
	s.storage.SetUint64("FrozenInfoCount", u)
}

func (s *Storage) GetUserAckCount() uint64 {
	return s.storage.GetUint64("UserAckCount")
}

func (s *Storage) SetUserAckCount(u uint64) {
	s.storage.SetUint64("UserAckCount", u)
}

func (s *Storage) GetPC() uint64 {
	return s.storage.GetUint64("_PC")
}

func (s *Storage) SetPC(u uint64) {
	s.storage.SetUint64("_PC", u)
}

func (s *Storage) GetMuPC() uint64 {
	return s.storage.GetUint64("_muPC")
}

func (s *Storage) SetMuPC(u uint64) {
	s.storage.SetUint64("_muPC", u)
}

func (s *Storage) Owners() *storage.BytesArray {
	return s.storage.NewBytesArray("Owners")
}

func (s *Storage) UserAcknowledged() *storage.MappingBytesToBytes {
	return s.storage.NewMappingBytesToBytes("UserAcknowledged")
}

func (s *Storage) Instructions() *storage.BytesArray {
	return s.storage.NewBytesArray("Instructions")
}

//transactions

func (s *Storage) MustFunds() *storage.MappingBytesToUint64 {
	return s.storage.NewMappingBytesToUint64("MustFunds")
}

func (s *Storage) IsOwner() *storage.MappingBytesToBool {
	return s.storage.NewMappingBytesToBool("IsOwner")
}

func (s *Storage) AidMap() *storage.MappingUint64ToUint64 {
	return s.storage.NewMappingUint64ToUint64("AidMap")
}
