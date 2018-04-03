package BinarySearch

import (
	"sort"
	"testing"
)

func SortArray(array []int) {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
}

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		slice    []int
		number   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 1},
		{[]int{5, 3, 7, 8, 2}, 2, 0},
		{[]int{5, -1, 7, 8, 2}, 2, 1},
		{[]int{5, 3, 7, 8, 2}, 9, -1},
	}

	for i, tc := range testCases {
		SortArray(tc.slice)
		index := BinarySearch(tc.slice, tc.number)
		if index != tc.expected {
			t.Errorf("BinarySearch(%v, %d) test case #%d: expected %d, got %d instead", tc.slice, tc.number, i+1, tc.expected, index)
		}
	}
}
