package main

func BubbleSort(array []int) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		firstItem := 0
		secondItem := 1
		for i := 0; i < len(array); i++ {
			if array[firstItem] > array[secondItem] {
				array[firstItem], array[secondItem] = array[secondItem], array[firstItem]
				swapCount += 1
			}
			firstItem += 1
			secondItem += 1
			if secondItem == len(array) {
				break
			}
		}
	}
}
