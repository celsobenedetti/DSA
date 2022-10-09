package graphs

import "github.com/celso-patiri/DSA/structures"

func BFS(g Graph, source, target int) (path []int) {
	nVertices := g.GetNumberOfVertices()

	if target >= nVertices {
		return []int{}
	}

	cameFrom := []int{}
	visited := make([]bool, nVertices)

	for i := 0; i < nVertices; i++ {
		cameFrom = append(cameFrom, EMPTY)
	}

	q := structures.Queue[int]{}

	visited[source] = true
	q.Enqueue(source)

	for q.Size() > 0 {
		curr, _ := q.Dequeue()
		if curr == target {
			break
		}

		adj := g.GetAdjecentNodes(curr)

		for _, edge := range adj {
			destination := edge.To

			if visited[destination] {
				continue
			}
			visited[destination] = true
			cameFrom[destination] = curr
			q.Enqueue(destination)
		}
	}

	if cameFrom[target] == EMPTY {
		return []int{}
	}

	//Path starts from target
	curr := target

	for cameFrom[curr] != EMPTY {
		path = append(path, curr)
		curr = cameFrom[curr]
	}

	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	//Prepend source to path
	return append([]int{source}, path...)
}
