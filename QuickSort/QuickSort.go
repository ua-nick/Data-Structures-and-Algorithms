package main

func QuickSort(array []int, left, right int) {
	if left < right {
		pivot := partition(array, left, right)
		QuickSort(array, left, pivot - 1)
		QuickSort(array, pivot + 1, right)
	}
}

func partition(array []int, left, right int) int {
	pivot := array[left]
	i, j := left, right
	for {
		for array[j] > pivot {
			j--
		}
		for array[i] <= pivot && i < j {
			i++
		}
		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}
	array[i], array[left] = array[left], array[i]
	return i
}

