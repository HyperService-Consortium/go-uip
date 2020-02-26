package opintent

import (
	"github.com/HyperService-Consortium/go-uip/op-intent/document"
	"github.com/HyperService-Consortium/go-uip/op-intent/errorn"
	"github.com/HyperService-Consortium/go-uip/op-intent/lexer"
	"github.com/HyperService-Consortium/go-uip/op-intent/parser"
)

type OpIntents interface {
	GetContents() [][]byte
	GetDependencies() [][]byte
}

func (ier *Initializer) Parse(
	opIntents OpIntents,
) (txIntents parser.TxIntents, err error) {
	l, r := &lexer.RawDocumentLexer{}, &parser.LexerResult{
		ContractMapping: make(lexer.AccountMap),
		AccountMapping:  make(lexer.AccountMap),
	}
	r.RootIntents, err = l.InitContents(opIntents.GetContents())
	if err != nil {
		return nil, err
	}

	r.RawDependencies, err = l.InitDependencies(
		opIntents.GetDependencies())
	if err != nil {
		return nil, err
	}
	return ier.Parse_(r)
}

type OpIntentsPacket interface {
	GetContent() (content []byte)
}

func (ier *Initializer) ParseR(opIntents OpIntentsPacket) (txIntents parser.TxIntents, err error) {
	res, err := document.NewGJSONDocument(opIntents.GetContent())
	if err != nil {
		return nil, err
	}

	l, r := &lexer.DocumentLexer{}, &parser.LexerResult{}

	r.RootIntents, err = l.InitContents(res)
	if err != nil {
		return nil, err
	}

	r.RawDependencies, err = l.InitDependencies(res.Get(lexer.FieldOpIntentsDependencies))
	if err != nil {
		return nil, err
	}

	r.ContractMapping, err = l.InitContracts(
		res.Get("contracts"))
	if err != nil {
		return nil, err
	}

	r.AccountMapping, err = l.InitAccounts(
		res.Get("accounts"))
	if err != nil {
		return nil, err
	}

	return ier.Parse_(r)
}

func (ier *Initializer) Parse_(l *parser.LexerResult) (txIntents parser.TxIntents, err error) {
	p := &parser.Parser{Program: l, AccountBase:ier.accountBase, ContractBase:ier.contractBase, ChainGetter:ier.chainGetter}


	deps, err := p.InitDependencies(
		l.RawDependencies, l.RootIntents.NameMap)
	if err != nil {
		return nil, err
	}

	intents, err := p.ParseIntents(l.RootIntents)
	if err != nil {
		return nil, err
	}

	// WARNING: ier.TopologicalSort assume that the size of total intents is <= 2 * len(rtx)
	if err = ier.TopologicalSort(intents, deps.Dependencies); err != nil {
		return nil, errorn.NewSortError(err)
	}
	return intents, nil
}
