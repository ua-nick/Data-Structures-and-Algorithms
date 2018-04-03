package InterpolationSearch

func InterpolationSearch(array []int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex < maxIndex {
		midIndex := minIndex + ((maxIndex - minIndex) / (array[maxIndex] - array[minIndex]) * (number - array[minIndex]))
		midItem := array[midIndex]
		if number == midItem {
			return midIndex
		}
		if midItem < number {
			minIndex = midIndex + 1
		} else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	if array[minIndex] == number {
		return minIndex
	} else if array[maxIndex] == number {
		return maxIndex
	} else {
		return -1
	}
}
