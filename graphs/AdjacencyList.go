package graphs

import "fmt"

type edges []Node

func (e edges) findEdge(to int) (idx int, node *Node) {
	for i, node := range e {
		if node.To == to {
			return i, &e[i]
		}
	}
	return EMPTY, nil
}

func (e *edges) removeEdge(idx int) {
	prev := *e
	*e = append(prev[:idx], prev[idx+1:]...)
}

type AdjList struct {
	vertices []edges
}

func (g *AdjList) CreateGraph(nVertices int) {
	*g = AdjList{}
	for i := 0; i < nVertices; i++ {
		g.vertices = append(g.vertices, edges{})
	}
}

func (g *AdjList) AddEdge(from, to, weight int) {
	g.vertices[from] = append(g.vertices[from], Node{To: to, Weight: weight})
}

func (g *AdjList) SetEdge(from, to, weight int) {
	_, node := g.vertices[from].findEdge(to)
	if node == nil {
		return
	}
	node.Weight = weight
}

func (g *AdjList) RemoveEdge(from, to int) {
	i, _ := g.vertices[from].findEdge(to)
	if i == EMPTY {
		return
	}
	g.vertices[from].removeEdge(i)
}

func (g *AdjList) GetEdge(from, to int) int {
	i, _ := g.vertices[from].findEdge(to)
	return i
}

func (g *AdjList) GetAdjecentNodesIndexes(from int) (idxs []int) {
	for _, e := range g.vertices[from] {
		idxs = append(idxs, e.To)
	}
	return idxs
}

func (g *AdjList) GetAdjecentNodes(from int) []Node {
	return g.vertices[from]
}

func (g *AdjList) GetNumberOfEdges() (nEdges int) {
	for _, v := range g.vertices {
		nEdges += len(v)
	}
	return nEdges
}

func (g *AdjList) GetNumberOfVertices() int {
	return len(g.vertices)
}

func (g *AdjList) GetRepresentationType() Representation {
	return ADJ_LIST
}

func (g *AdjList) PrintGraph() {
	for from, v := range g.vertices {
		for _, edge := range v {
			fmt.Printf("[%d,%d|%d|] ", from, edge.To, edge.Weight)
		}
		fmt.Println("")
	}
}
