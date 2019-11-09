package uiptypes


type Checker interface {
	CheckAddress(Address) bool
}

type CheckerGetter interface {
	GetChecker(ChainID) Checker
}
