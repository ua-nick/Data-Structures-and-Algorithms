package MergeSort

func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}

func MergeSort(array []int) {
	if len(array) <= 1 {
		return
	}
	left, right := array[:len(array)/2], array[len(array)/2:]
	MergeSort(left)
	MergeSort(right)
	copy(array, merge(left, right))
}
