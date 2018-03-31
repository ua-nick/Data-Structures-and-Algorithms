package main

const arraySize = 10

func ExponentialSearch(array [arraySize]int, number int) int {
	if array[0] == number {
		return 0
	}
	boundValue := 1
	for boundValue < len(array) && array[boundValue] < number {
		boundValue *= 2
	}
	if boundValue > len(array) {
		boundValue = len(array)
	}
	return BinarySearch(array, boundValue, number)
}

func BinarySearch(array [arraySize]int, bound, number int) int {
	minIndex := 0
	maxIndex := bound
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
