package algorithms

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) (pivot int) {
	pivot = arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	//Fit pivot right after idx
	arr[hi], arr[idx] = arr[idx], pivot

	return idx
}

func QuickSort(arr []int) []int {
	qs(arr, 0, len(arr)-1)
	return arr
}
