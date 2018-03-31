package main

const arraySize = 10

func LinearSearch(array [arraySize]int, number int) int {
	for index, value := range array {
		if value == number {
			return index
		}
	}
	return -1
}
