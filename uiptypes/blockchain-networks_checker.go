package uiptypes


type Checker interface {
	CheckAddress(Address) error
}

type CheckerGetter interface {
	GetChecker(ChainID) Checker
}
