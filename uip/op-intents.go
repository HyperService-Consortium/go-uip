package uip


type OpIntents interface {
	GetContents() (contents [][]byte)
	GetDependencies() (dependencies [][]byte)
}