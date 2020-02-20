package uip


type Checker interface {
	CheckAddress(Address) error
}

type CheckerGetter interface {
	GetChecker(ChainID) (checker Checker, err error)
}
