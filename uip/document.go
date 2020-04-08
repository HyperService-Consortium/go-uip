package uip

import "github.com/HyperService-Consortium/go-uip/internal/document"

type Document = document.Document
type Documents = document.Documents
type GJSONDocument = document.GJSONDocument
type MapDocument = document.MapDocument
type GJSONDocuments = document.GJSONDocuments
type MapDocuments = document.MapDocuments

func NewMapDocument(m interface{}) (MapDocument, error) {
	return document.NewMapDocument(m)
}

func NewGJSONDocument(b []byte) (GJSONDocument, error) {
	return document.NewGJSONDocument(b)
}

func NewMapDocuments(b []interface{}) (MapDocuments, error) {
	return document.NewMapDocuments(b)
}
