package storage

type Node struct {
	Id int64

	Name string
	Data map[string]string

	InEdges  []*Edge
	OutEdges []*Edge
}

type Edge struct {
	Id int64

	Name string
	Data map[string]string

	InNodes  []*Node
	OutNodes []*Node
}
