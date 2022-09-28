package main

import "fmt"

type MaxHeap struct {
	array []int
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func (h *MaxHeap) swap(parent, child int) {
	h.array[parent], h.array[child] = h.array[child], h.array[parent]
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	left, right := leftChild(index), rightChild(index)

	if left >= len(h.array)-1 {
		return
	}

	if h.array[left] >= h.array[index] && h.array[left] > h.array[right] {
		h.swap(index, left)
		h.heapifyDown(left)
	}

	if h.array[right] >= h.array[index] && h.array[right] >= h.array[left] {
		h.swap(index, right)
		h.heapifyDown(right)
	}
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("MaxHeap Extract err: Heap is empty")
		return -1
	}

	first := h.array[0]

	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.heapifyDown(0)

	return first
}
