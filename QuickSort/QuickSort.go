package main

func QuickSort(array []int, p, r int) {
	if p < r {
		q := partition(array, p, r)
		QuickSort(array, p, q-1)
		QuickSort(array, q+1, r)
	}
}

func partition(array []int, p, r int) int {
	x := array[r]
	i := p - 1
	for j := p; j < r; j++ {
		if array[j] <= x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[r], array[i+1] = array[i+1], array[r]
	return i + 1
}
