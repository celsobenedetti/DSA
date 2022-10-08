package bt

func walkPost(curr *Node, path *[]int) []int {
	if curr == nil {
		return *path
	}

	//pre
	//recurse
	walkPost(curr.Left, path)
	walkPost(curr.Right, path)

	//post
	//visit node
	*path = append(*path, curr.Value)

	return *path
}

func PostOrderSearch(head *Node) (path []int) {
	return walkPost(head, &path)
}
