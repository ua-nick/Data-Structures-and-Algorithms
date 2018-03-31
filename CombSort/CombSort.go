package main

const arraySize = 10

func CombSort(array *[arraySize]int) {
	gapValue := len(array)
	swapCount := 1
	for gapValue >= 1 && swapCount != 0 {
		if gapValue != 1 {
			gapValue = int(float64(gapValue) / float64(1.3))
		}
		swapCount = 0
		firstItem := 0
		secondItem := gapValue
		for secondItem != len(array) {
			if array[firstItem] > array[secondItem] {
				array[firstItem], array[secondItem] = array[secondItem], array[firstItem]
				swapCount += 1
			}
			firstItem += 1
			secondItem += 1
		}
	}
}
