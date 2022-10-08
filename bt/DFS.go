package bt

func DFS(node *Node, target int) bool {
	if node == nil {
		return false
	}

	if node.Value == target {
		return true
	}

	if target < node.Value {
		return DFS(node.Left, target)
	} else {
		return DFS(node.Right, target)
	}
}
