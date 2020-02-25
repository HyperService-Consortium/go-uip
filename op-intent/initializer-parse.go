package opintent

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
)

type OpIntents interface {
	GetContents() [][]byte
	GetDependencies() [][]byte
}

func (ier *Initializer) Parse(
	opIntents OpIntents,
) (txIntents TxIntents, err error) {
	l := &lexer.RawDocumentLexer{}
	rawIntents, err := l.InitContents(opIntents.GetContents())
	if err != nil {
		return nil, err
	}

	rawDependencies, err := l.InitDependencies(
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
	res, err := document.NewGJSONDocument(opIntents.GetContent())
	if err != nil {
		return nil, err
	}

	l := &lexer.DocumentLexer{}

	rawIntents, err := l.InitContents(res)
	if err != nil {
		return nil, err
	}

	rawDependencies, err := l.InitDependencies(res.Get(lexer.FieldOpIntentsDependencies))
	if err != nil {
		return nil, err
	}

	return ier.Parse_(rawIntents, rawDependencies)
}

func (ier *Initializer) Parse_(rawIntents *lexer.RootIntents, rawDeps *lexer.RawDependencies) (txIntents TxIntents, err error) {

	deps, err := ier.InitDependencies(
		rawDeps, rawIntents.NameMap)
	if err != nil {
		return nil, err
	}

	intents, err := ier.ParseIntents(rawIntents)
	if err != nil {
		return nil, err
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(intents, deps.dependencies); err != nil {
		return nil, errorn.NewSortError(err)
	}
	return intents, nil
}
