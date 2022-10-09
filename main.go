package main

import (
	"fmt"

	"github.com/celso-patiri/DSA/graphs"
)

func runBfs(g graphs.Graph) {
	g.CreateGraph(4)

	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(2, 3, 4)
	// g[1][3] = 3

	fmt.Println(graphs.BFS(g, 0, 3))

	g.PrintGraph()

}

func main() {
	var g graphs.Graph
	g = &graphs.AdjMatrix{}

	runBfs(g)
}
