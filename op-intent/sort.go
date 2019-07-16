package opintent

// the LeftName intent is before RightName intent
type Dependency struct {
	LeftName  uint64
	RightName uint64
}

type Grapher struct {
	Edge    []Dependency
	Counter []uint64
}

func (ier *OpIntentInitializer) TopologicalSort(
	transactionIntents [][]*TransactionIntent,
	rawDependencies [][]byte,
) error {
	// var dependencyGrapher Grapher
	return nil
}
