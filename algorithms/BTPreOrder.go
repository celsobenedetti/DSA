package algorithms

func walkPre(curr *Node, path *[]int) []int {
	if curr == nil {
		return *path
	}

	//pre
	*path = append(*path, curr.Value)

	//recurse
	walkPre(curr.Left, path)
	walkPre(curr.Right, path)

	//post
	return *path
}

func PreOrderSearch(head *Node) (path []int) {
	return walkPre(head, &path)
}
