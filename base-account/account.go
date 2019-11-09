package base_account

type Account struct {
	ChainId uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Account) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}
