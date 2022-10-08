package algorithms

func BST_DFSFind(node *Node, target int) bool {
	if node == nil {
		return false
	}

	if node.Value == target {
		return true
	}

	if target < node.Value {
		return BST_DFSFind(node.Left, target)
	} else {
		return BST_DFSFind(node.Right, target)
	}
}
