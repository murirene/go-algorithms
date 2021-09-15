package sort

import "math"

func MergeSort(values []int) {
	buffer := make([]int, len(values))
	mergesort(values, 0, len(values)-1, buffer)
}

/*
           [4, 3, 1, 6, 7], 0, 4
      [4, 3, 1], 0, 2
   [4, 3], 0, 1
   4, [0, 0]   3, [1, 1]

*/

func mergesort(values []int, start, end int, buffer []int) {
	if start == end {
		return
	}
	mid := int(math.Floor(float64((start + end) / 2)))
	mergesort(values, start, mid, buffer)
	mergesort(values, mid+1, end, buffer)
	merge(values, start, mid, end, buffer)
}

func merge(values []int, start, mid, end int, buffer []int) {
	leftIdx := start
	sortedIdx := start
	rightIdx := mid + 1
	for leftIdx <= mid || rightIdx <= end {
		if rightIdx > end || (leftIdx <= mid && values[leftIdx] < values[rightIdx]) {
			buffer[sortedIdx] = values[leftIdx]
			leftIdx += 1
		} else {
			buffer[sortedIdx] = values[rightIdx]
			rightIdx += 1
		}
		sortedIdx += 1
	}

	for i := start; i <= end; i++ {
		values[i] = buffer[i]
	}
}
