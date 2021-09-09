package sort_test

import (
	"go-algorithms/sort"
	"testing"
)

func MergeSortTest(t *testing.T) {
	values := []int{4, 3, 6, 7, 1}
	results := []int{1, 3, 4, 6, 7}
	sort.MergeSort(values)
	for v, i := range values {
		if v != results[i] {
			t.Fatalf("Failed to merge sort %v", values)
		}
	}
}

func MergeSortTest2(t *testing.T) {
    values := []int{3}
    results := []int{3}
	sort.MergeSort(values)
	for v, i := range values {
		if v != results[i] {
			t.Fatalf("Failed to merge sort %v", values)
		}
	}

}

func MergeSortTest3(t *testing.T) {
    values := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
    results := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.MergeSort(values)
	for v, i := range values {
		if v != results[i] {
			t.Fatalf("Failed to merge sort %v", values)
		}
	}

}

func MergeSortTest4(t *testing.T) {
    values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    results := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sort.MergeSort(values)
	for v, i := range values {
		if v != results[i] {
			t.Fatalf("Failed to merge sort %v", values)
		}
	}

}
