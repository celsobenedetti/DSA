package algorithms_test

import (
	"math/rand"
	"testing"

	"github.com/celso-patiri/DSA/algorithms"
)

func NewRandomSlice() []int {
	size := rand.Intn(maxSize) + 3
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

func TestQuickSort(t *testing.T) {
	for i := 0; i < maxSize; i++ {
		arr := NewRandomSlice()
		arr = algorithms.QuickSort(arr)

		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				t.Errorf("QuickSort: unsorted array")
			}
		}
	}
}
