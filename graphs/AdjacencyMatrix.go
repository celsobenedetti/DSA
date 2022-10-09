package graphs

import (
	"fmt"
)

type AdjMatrix struct {
	matrix [][]int
}

func (g *AdjMatrix) CreateGraph(n int) {
	*g = AdjMatrix{}
	for i := 0; i < n; i++ {
		g.matrix = append(g.matrix, make([]int, 0))
		for j := 0; j < n; j++ {
			g.matrix[i] = append(g.matrix[i], EMPTY)
		}
	}
}

func (g *AdjMatrix) AddEdge(from, to, weight int) {
	g.matrix[from][to] = weight
}

func (g *AdjMatrix) GetEdge(from, to int) int {
	return g.matrix[from][to]
}

func (g *AdjMatrix) SetEdge(from, to, weight int) {
	g.matrix[from][to] = weight
}

func (g *AdjMatrix) RemoveEdge(from, to int) {
	g.matrix[from][to] = EMPTY
}

func (g *AdjMatrix) GetNumberOfVertices() int {
	return len(g.matrix)
}

func (g *AdjMatrix) GetNumberOfEdges() (n int) {
	for _, idx := range g.matrix {
		for _, edge := range idx {
			if edge != EMPTY {
				n++
			}
		}
	}
	return n
}

func (g *AdjMatrix) GetAdjecentNodesIndexes(from int) (idxs []int) {
	for i := range g.matrix[from] {
		if i != EMPTY {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func (g *AdjMatrix) GetAdjecentNodes(from int) (nodes []Node) {
	for i, w := range g.matrix[from] {
		if w != EMPTY {
			nodes = append(nodes, Node{To: i, Weight: w})
		}
	}
	return nodes
}

func (g *AdjMatrix) GetRepresentationType() Representation {
	return ADJ_MATRIX
}

func (g *AdjMatrix) PrintGraph() {
	for from, vtx := range g.matrix {
		for to, weight := range vtx {
			if weight != EMPTY {
				fmt.Printf("[%d,%d|%d|] ", from, to, weight)
			}
		}
		fmt.Println("")
	}
}
