package algorithms

func BFS(head *Node, target int) bool {
	q := NodeQueue{}
	q.Push(head)

	for q.Size() > 0 {
		curr, _ := q.Dequeue()

		if curr.Value == target {
			return true
		}

		if curr.Left != nil {
			q.Push(curr.Left)
		}

		if curr.Right != nil {
			q.Push(curr.Right)
		}

	}

	return false
}
