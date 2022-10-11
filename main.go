package main

import (
	"fmt"

	"github.com/celso-patiri/DSA/graphs"
)

func run(g graphs.Graph) {
	g.CreateGraph(4)

	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 8)
	g.AddEdge(2, 3, 4)

	fmt.Println(graphs.Dijkstra(g, 0, 3))

	g.PrintGraph()

}

func main() {
	var g1 graphs.Graph
	var g2 graphs.Graph
	g1 = &graphs.AdjMatrix{}
	g2 = &graphs.AdjList{}

	run(g1)
	run(g2)
}
