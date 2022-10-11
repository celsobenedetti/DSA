package graphs

import (
	"math"
)

type SeenArray []bool

func (s SeenArray) hasUnvisited(dists []int) bool {
	for i, seen := range s {
		if !seen && dists[i] < math.MaxInt {
			return true
		}
	}
	return false
}

func (s SeenArray) getLowestUnvisited(dists []int) int {
	idx := EMPTY
	min := math.MaxInt
	for i, seen := range s {
		if !seen && dists[i] < min {
			min = dists[i]
			idx = i
		}
	}
	return idx
}

func Dijkstra(g Graph, source, target int) (path []int) {
	nVertices := g.GetNumberOfVertices()

	if target > nVertices {
		return []int{}
	}

	prev := make([]int, nVertices)
	seen := make(SeenArray, nVertices)
	dists := make([]int, nVertices)

	for i := 0; i < nVertices; i++ {
		prev[i] = EMPTY
		dists[i] = math.MaxInt
	}

	dists[source] = 0

	for seen.hasUnvisited(dists) {
		curr := seen.getLowestUnvisited(dists)
		seen[curr] = true

		for _, edge := range g.GetAdjecentNodes(curr) {
			if seen[edge.To] {
				continue
			}

			dist := edge.Weight + dists[curr]

			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	curr := target

	for prev[curr] != EMPTY {
		path = append(path, curr)
		curr = prev[curr]
	}

	//Append source to path
	path = append(path, source)

	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
