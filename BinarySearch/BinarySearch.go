package main

const arraySize = 10

func BinarySearch(array [arraySize]int, number int) int {
	minIndex := 0
	maxIndex := len(array)
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
