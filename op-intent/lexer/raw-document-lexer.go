package lexer

import (
	"github.com/HyperService-Consortium/go-uip/errorn"
	"github.com/HyperService-Consortium/go-uip/internal/document"
)

type RawDocumentLexer struct {
	BaseLexer
}

func (l *RawDocumentLexer) InitContents(contents [][]byte) (r *RootIntents, err error) {
	r = &RootIntents{}
	r.Infos = make([]Intent, len(contents))
	r.NameMap = make(map[string]int)

	for idx, content := range contents {
		r.Infos[idx], err = l.InitContent(content)
		if err != nil {
			return nil, err.(*errorn.ParseError).Desc(errorn.AtOpIntentsPos{Pos: idx})
		}
		r.NameMap[r.Infos[idx].GetName()] = idx
	}
	return
}

func (l *RawDocumentLexer) InitContent(content []byte) (i Intent, err error) {
	c, err := document.NewGJSONDocument(content)
	if err != nil {
		return nil, err
	}
	return (&DocumentLexer{}).InitContent(c)
}

func (l *RawDocumentLexer) InitDependencies(rawDeps [][]byte) (deps *RawDependencies, err error) {
	deps = &RawDependencies{
		dependencies: make([]RawDependency, len(rawDeps)),
	}
	var res document.Document
	for idx, dependency := range rawDeps {
		res, err = document.NewGJSONDocument(dependency)
		if err != nil {
			return
		}
		err = deps.dependencies[idx].UnmarshalResult(res)
		if err != nil {
			return
		}
	}
	return
}
