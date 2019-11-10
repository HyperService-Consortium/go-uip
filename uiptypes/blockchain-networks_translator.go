package uiptypes


type Translator interface {
	Translate(intent *TransactionIntent, storage Storage) (rawTransaction RawTransaction, err error)
}

type TranslatorGetter interface {
	GetTranslator(chainID ChainID) (intent Translator)
}
