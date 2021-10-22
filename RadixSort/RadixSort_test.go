package RadixSort

import (
	"sort"
	"testing"
)

func TestRadixSort(t *testing.T) {
	array1 := []int{421, 15, -175, 90, -2, 214, -52, -166}
	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)
	RadixSort(array1)
	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}
}
