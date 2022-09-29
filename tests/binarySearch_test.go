package algorithms_test

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"

	"github.com/celso-patiri/DSA/algorithms"
)

var maxSize = 100

func newRandomSortedSlice() []int {
	size := rand.Intn(maxSize) + 3
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func TestBinarySearch(t *testing.T) {
	for a := 0; a < 100; a++ {
		arr := newRandomSortedSlice()

		notFound := algorithms.BinarySearch(arr, -1)
		if !reflect.DeepEqual(notFound, -1) {
			t.Errorf("Binary search: expected %d got %d\n", -1, notFound)
		}

		for i, v := range arr {
			index := algorithms.BinarySearch(arr, v)
			if !reflect.DeepEqual(index, i) {
				t.Errorf("Binary search: expected %d got %d\n", i, index)
			}
		}
	}

}
