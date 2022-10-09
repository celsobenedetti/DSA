package bt

func BFS(head *Node, target int) bool {
	q := NodeQueue{}
	q.Enqueue(head)

	for q.Size() > 0 {
		curr, _ := q.Dequeue()

		if curr.Value == target {
			return true
		}

		if curr.Left != nil {
			q.Enqueue(curr.Left)
		}

		if curr.Right != nil {
			q.Enqueue(curr.Right)
		}

	}

	return false
}
