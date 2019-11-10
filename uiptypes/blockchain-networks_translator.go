package uiptypes


type Translator interface {
	Translate(intent *TransactionIntent, kvGetter KVGetter) (rawTransaction RawTransaction, err error)
}

type TranslatorGetter interface {
	GetTranslator(chainID ChainID) (intent Translator)
}
