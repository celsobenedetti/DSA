package bt

func walkIn(curr *Node, path *[]int) []int {
	if curr == nil {
		return *path
	}

	//pre
	walkIn(curr.Left, path)

	//visit node
	*path = append(*path, curr.Value)

	//recurse
	walkIn(curr.Right, path)

	//post
	return *path
}

func InOrderSearch(head *Node) (path []int) {
	return walkIn(head, &path)
}
