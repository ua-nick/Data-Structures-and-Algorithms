package ExponentialSearch

func ExponentialSearch(array []int, number int) int {
	boundValue := 1
	for boundValue < len(array) && array[boundValue] < number {
		boundValue *= 2
	}
	if boundValue > len(array) {
		boundValue = len(array) - 1
	}
	return BinarySearch(array, boundValue+1, number)
}

func BinarySearch(array []int, bound, number int) int {
	minIndex := 0
	maxIndex := bound - 1
	for minIndex <= maxIndex {
		midIndex := int((maxIndex + minIndex) / 2)
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
	return -1
}
