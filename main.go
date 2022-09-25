package main

import (
	"celso/data-structures/heap"
	"fmt"
)

func main() {
	maxHeap := heap.MaxHeap{}

	buildHeap := []int{1, 2, 136, 3, 147, 158, 169, 4, 6, 8, 10, 12, 14, 70, 81, 92, 103, 114, 125, 180, 191}

	for _, v := range buildHeap {
		maxHeap.Insert(v)
		fmt.Println(maxHeap)
	}

	for i, j := 0, len(buildHeap)-1; i < j; i, j = i+1, j-1 {
		buildHeap[i], buildHeap[j] = buildHeap[j], buildHeap[i]
	}

	for range buildHeap {
		maxHeap.Extract()
		fmt.Println(maxHeap)
	}

}
