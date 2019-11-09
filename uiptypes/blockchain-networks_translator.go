package uiptypes


type Translator interface {
	Translate(*TransactionIntent, KVGetter) (RawTransaction, error)
}

type TranslatorGetter interface {
	GetTranslator(ChainID) Translator
}
