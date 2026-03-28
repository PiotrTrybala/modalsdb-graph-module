package storage

type Node struct {
	Id int64

	Name string
	Data map[string]string

	In  []*Edge
	Out []*Edge
}

type Edge struct {
	Id int64

	Name string
	Data map[string]string

	Source      *Node
	Destination *Node
}
