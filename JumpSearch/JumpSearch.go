package JumpSearch

import "math"

func JumpSearch(array []int, number int) int {
	jumpValue := int(math.Floor(math.Sqrt(float64(len(array)))))
	minIndex := 0
	maxIndex := jumpValue
	for array[maxIndex] <= number {
		minIndex += jumpValue
		maxIndex = minIndex + jumpValue
		if maxIndex >= len(array) {
			maxIndex = len(array)
			minIndex = maxIndex - jumpValue
			break
		}
	}
	for i := minIndex; i < maxIndex; i++ {
		if array[i] == number {
			return i
		}
	}
	return -1
}
