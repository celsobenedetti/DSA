package structures

type MinHeap struct {
	items []int
}

func (h MinHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (h MinHeap) leftChild(idx int) int {
	return (2 * idx) + 1
}

func (h MinHeap) rightChild(idx int) int {
	return (2 * idx) + 2
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	currentValue := h.items[idx]

	parentIdx := h.parent(idx)
	parentValue := h.items[parentIdx]

	if parentValue > currentValue {
		h.items[idx], h.items[parentIdx] = parentValue, currentValue
		h.heapifyUp(parentIdx)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	if idx >= len(h.items) {
		return
	}

	leftIdx, rightIdx := h.leftChild(idx), h.leftChild(idx)

	if leftIdx >= len(h.items) {
		return
	}

	currentValue, leftValue, rightValue := h.items[idx], h.items[leftIdx], h.items[rightIdx]

	if leftValue >= rightValue && currentValue >= rightValue {
		h.items[idx], h.items[rightIdx] = rightValue, currentValue
		h.heapifyDown(rightIdx)
	}
	if rightValue >= leftValue && currentValue >= leftValue {
		h.items[idx], h.items[leftIdx] = leftValue, currentValue
		h.heapifyDown(leftIdx)
	}
}

func (h *MinHeap) Insert(data int) {
	h.items = append(h.items, data)
	h.heapifyUp(len(h.items))
}

func (h *MinHeap) Pop() int {
	if len(h.items) == 0 {
		return -1
	}
	out := h.items[0]

	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]

	h.heapifyDown(0)
	return out
}
