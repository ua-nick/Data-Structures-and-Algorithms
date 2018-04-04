package InterpolationSearch

func InterpolationSearch(array []int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex && number >= array[minIndex] && number <= array[maxIndex] {
		midIndex := minIndex + (number-array[minIndex])*(maxIndex-minIndex)/(array[maxIndex]-array[minIndex])
		midItem := array[midIndex]
		if midItem == number {
			return midIndex
		} else if midItem < number {
			minIndex = midIndex + 1
		} else if midItem > number {
			maxIndex = midIndex - 1
		}
	}
	return -1
}
