package algorithms

import (
	"math"
	"sort"
)

func BinarySearch(arr []int, value int) int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return binarySearchRec(arr, value, 0, len(arr))
}

func binarySearchRec(arr []int, value, low, high int) int {
	middle := int(math.Floor(float64(low + (high-low)/2)))

	if arr[middle] == value {
		return middle
	}

	if value > arr[middle] {
		return binarySearchRec(arr, value, middle+1, high)
	}
	if value < arr[middle] {
		return binarySearchRec(arr, value, low, middle)
	}

	return -1
}
