package uip

type BlockChainNilImpl struct {}

func (b BlockChainNilImpl) RouteRaw(chainID ChainID, rawTransaction RawTransaction) (
	transactionReceipt TransactionReceipt, err error) {
	panic("implement me")
}

func (b BlockChainNilImpl) WaitForTransact(chainID ChainID, transactionReceipt TransactionReceipt,
	rOptions ...interface{}) (blockID BlockID, color []byte, err error) {
	panic("implement me")
}

func (b BlockChainNilImpl) MustWithSigner() (yesOrNot bool) {
	panic("implement me")
}

func (b BlockChainNilImpl) RouteWithSigner(signer Signer) (Router, error) {
	panic("implement me")
}

func (b BlockChainNilImpl) Translate(intent TransactionIntent, storage Storage) (rawTransaction RawTransaction, err error) {
	panic("implement me")
}

func (b BlockChainNilImpl) Deserialize(raw []byte) (rawTransaction RawTransaction, err error) {
	panic("implement me")
}

func (b BlockChainNilImpl) ParseTransactionIntent(intent TxIntentI) (TxIntentI, error) {
	return intent, nil
}

func (b BlockChainNilImpl) CheckAddress(Address) error {
	panic("implement me")
}

func (b BlockChainNilImpl) GetTransactionProof(chainID ChainID, blockID BlockID, color []byte) (MerkleProof, error) {
	panic("implement me")
}

func (b BlockChainNilImpl) GetStorageAt(chainID ChainID, typeID TypeID, contractAddress ContractAddress, pos []byte, description []byte) (Variable, error) {
	panic("implement me")
}

type BlockChainGetterNilImpl struct {}

func (b BlockChainGetterNilImpl) GetChecker(ChainID) (checker Checker, err error) {
	panic("implement me")
}

func (b BlockChainGetterNilImpl) GetTranslator(chainID ChainID) (intent Translator, err error) {
	return BlockChainNilImpl{}, nil
}

func (b BlockChainGetterNilImpl) GetRouter(chainID ChainID) (router Router, err error) {
	panic("implement me")
}

func (b BlockChainGetterNilImpl) GetBlockStorage(chainID ChainID) (storage Storage, err error) {
	panic("implement me")
}

func (b BlockChainGetterNilImpl) GetBlockChainInterface(ChainID) BlockChainInterface {
	panic("implement me")
}
