package TernarySearch

func TernarySearch(array []int, number int) int {
	minIndex := 0
	maxIndex := len(array) - 1
	for minIndex <= maxIndex {
		midIndex1 := minIndex + int((maxIndex-minIndex)/3)
		midIndex2 := maxIndex - int((maxIndex-minIndex)/3)
		midItem1 := array[midIndex1]
		midItem2 := array[midIndex2]
		if midItem1 == number {
			return midIndex1
		} else if midItem2 == number {
			return midIndex2
		}
		if midItem1 < number {
			minIndex = midIndex1 + 1
		} else if midItem2 > number {
			maxIndex = midIndex2 - 1
		} else {
			minIndex = midIndex1 + 1
			maxIndex = midIndex2 - 1
		}
	}
	return -1
}
