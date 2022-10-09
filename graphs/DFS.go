package graphs

func walk(g Graph, curr, target int, seen []bool, path *[]int) (found bool) {
	if seen[curr] {
		return false
	}
	//pre
	seen[curr] = true
	*path = append(*path, curr)

	if curr == target {
		return true
	}

	//recurse
	adjecent := g.GetAdjecentNodes(curr)
	for _, edge := range adjecent {
		if walk(g, edge.To, target, seen, path) {
			return true
		}
	}

	//post
	//pop current from path if not found
	*path = (*path)[:len(*path)-1]

	return false
}

func DFS(g Graph, source, target int) (path []int) {
	size := g.GetNumberOfVertices()

	if target >= size {
		return
	}

	walk(g, source, target, make([]bool, size), &path)

	return path
}
