package sort

import "math"

func MergeSort(values []int) {
	mergesort(values, 0, len(values)-1)
}

/*
           [4, 3, 1, 6, 7], 0, 4
      [4, 3, 1], 0, 2
   [4, 3], 0, 1
   4, [0, 0]   3, [1, 1]

*/

func mergesort(values []int, start, end int) {
	if start == end {
		return
	}
	mid := int(math.Floor(float64((start + end) / 2)))
	mergesort(values, start, mid)
	mergesort(values, mid+1, end)
	merge(values, start, mid, end)
}

func merge(values []int, start, mid, end int) {
	i := start
	j := mid + 1
	for i <= mid && j <= end {
		if values[j] < values[i] {
			tmp := values[j]
			values[j] = values[i]
			values[i] = tmp
			j += 1
		} else {
			i += 1
		}
	}
}
