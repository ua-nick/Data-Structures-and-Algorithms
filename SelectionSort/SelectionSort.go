package SelectionSort

func SelectionSort(array []int) {
	for arrayIndex := range array {
		minValue := array[arrayIndex]
		minIndex := arrayIndex
		for subArrayIndex := arrayIndex + 1; subArrayIndex < len(array); subArrayIndex++ {
			if array[subArrayIndex] < minValue {
				minValue = array[subArrayIndex]
				minIndex = subArrayIndex
			}
		}
		array[minIndex], array[arrayIndex] = array[arrayIndex], array[minIndex]
	}
}
