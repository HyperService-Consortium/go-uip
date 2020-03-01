package isc

import "github.com/HyperService-Consortium/go-uip/storage"

const (
	state = "State"
)


type Storage struct {
	storage *storage.VM
}

func (s *Storage) getISCState() uint8 {
	return s.storage.GetUint8(state)
}

func (s *Storage) setISCState(u uint8)  {
	s.storage.SetUint8(state, u)
}

func (s *Storage) getFrozenInfoCount() uint64 {
	return s.storage.GetUint64("FrozenInfoCount")
}

func (s *Storage) setFrozenInfoCount(u uint64)  {
	s.storage.SetUint64("FrozenInfoCount", u)
}

func (s *Storage) getUserAckCount() uint64 {
	return s.storage.GetUint64("UserAckCount")
}

func (s *Storage) setUserAckCount(u uint64)  {
	s.storage.SetUint64("UserAckCount", u)
}

func (s *Storage) getPC() uint64 {
	return s.storage.GetUint64("_PC")
}

func (s *Storage) setPC(u uint64)  {
	s.storage.SetUint64("_PC", u)
}

//

//

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