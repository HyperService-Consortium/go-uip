package uiptypes


type content = []byte
type dependencies = []byte

type OpIntents interface {
	GetContents() []content
	GetDependencies() []dependencies
}