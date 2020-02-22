package opintent


type OpIntents interface {
	GetContents() [][]byte
	GetDependencies() [][]byte
}

func (ier *Initializer) Parse(
	opIntents OpIntents,
) (txIntents TxIntents, err error) {
	rawIntents, err := ier.InitContents(opIntents.GetContents())
	if err != nil {
		return nil, err
	}

	rawDependencies, err := ier.InitDependencies(
		opIntents.GetDependencies())
	if err != nil {
		return nil, err
	}
	return ier.Parse_(rawIntents, rawDependencies)
}

type OpIntentsPacket interface {
	GetContent() (content []byte)
}

func (ier *Initializer) ParseR(opIntents OpIntentsPacket) (txIntents TxIntents, err error) {
	res, err := NewGJSONResult(opIntents.GetContent())
	if err != nil {
		return nil, err
	}

	rawIntents, err := ier.InitContentsR(res)
	if err != nil {
		return nil, err
	}

	rawDependencies, err := ier.InitDependenciesR(res.Get(FieldOpIntentsDependencies))
	if err != nil {
		return nil, err
	}

	return ier.Parse_(rawIntents, rawDependencies)
}

func (ier *Initializer) Parse_(rawIntents *RawIntents, rawDeps *RawDependencies) (txIntents TxIntents, err error) {

	deps, err := ier.ParseDependencies(
		rawDeps, rawIntents.nameMap)
	if err != nil {
		return nil, err
	}

	intents, err := ier.ParseIntents(rawIntents)
	if err != nil {
		return nil, err
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(intents, deps.dependencies); err != nil {
		return nil, newSortError(err)
	}
	return intents, nil
}
