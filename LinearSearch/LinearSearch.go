package LinearSearch

func LinearSearch(array []int, number int) int {
	for index, value := range array {
		if value == number {
			return index
		}
	}
	return -1
}
