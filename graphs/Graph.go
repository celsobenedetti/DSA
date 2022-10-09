package graphs

type Representation int

const (
	ADJ_LIST Representation = iota
	ADJ_MATRIX
)

var EMPTY = -1

type Node struct {
	To     int
	Weight int
}

type Graph interface {
	CreateGraph(nVertices int)

	AddEdge(from, to, weight int)
	SetEdge(from, to, weight int)
	GetEdge(from, to int) int
	RemoveEdge(from, to int)

	GetAdjecentNodesIndexes(from int) []int
	GetAdjecentNodes(from int) []Node

	GetNumberOfEdges() int
	GetNumberOfVertices() int
	GetRepresentationType() Representation
	PrintGraph()
}
