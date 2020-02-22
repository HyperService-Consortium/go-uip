package opintent

type BaseOpIntent struct {
	Name         string `json:"name"`
	OpTypeString string `json:"op_type"`
}

type ContractInvokeMeta struct {
	Code     []byte     `json:"contract_code"` // key
	FuncName string     `json:"func"`
	Params   []RawParam `json:"parameters"`
	Meta     []byte     `json:"meta"`
}

type OpIntents interface {
	GetContents() [][]byte
	GetDependencies() [][]byte
}

func (ier *Initializer) Parse(
	opIntents OpIntents,
) (txIntents TxIntents, err error) {
	initializedIntents, err := ier.InitContents(opIntents.GetContents())
	if err != nil {
		return nil, err
	}

	initializedDependencies, err := ier.InitDependencies(
		opIntents.GetDependencies())
	if err != nil {
		return nil, err
	}

	parsedDependencies, err := ier.ParseDependencies(
		initializedDependencies, initializedIntents.nameMap)
	if err != nil {
		return nil, err
	}

	intents, err := ier.ParseIntents(initializedIntents)
	if err != nil {
		return nil, err
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(intents, parsedDependencies.dependencies); err != nil {
		return nil, newSortError(err)
	}

	return intents, nil
}

type OpIntentsPacket interface {
	GetContent() (content []byte)
}

func handlerRawIntents(opIntents OpIntentsPacket) (ResultI, error) {
	return NewGJSONResult(opIntents.GetContent())
}

func (ier *Initializer) ParseR(opIntents OpIntentsPacket) (txIntents TxIntents, err error) {
	res, err := handlerRawIntents(opIntents)
	if err != nil {
		return nil, err
	}

	initializedIntents, err := ier.InitContentsR(res)
	if err != nil {
		return nil, err
	}

	initializedDependencies, err := ier.InitDependenciesR(res)
	if err != nil {
		return nil, err
	}

	parsedDependencies, err := ier.ParseDependencies(
		initializedDependencies, initializedIntents.nameMap)
	if err != nil {
		return nil, err
	}

	intents, err := ier.ParseIntents(initializedIntents)
	if err != nil {
		return nil, err
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(intents, parsedDependencies.dependencies); err != nil {
		return nil, newSortError(err)
	}

	return intents, nil
}
