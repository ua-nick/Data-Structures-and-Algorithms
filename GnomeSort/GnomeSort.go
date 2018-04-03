package GnomeSort

func GnomeSort(array []int) {
	itemIndex := 0
	for itemIndex < len(array)-1 {
		if array[itemIndex] > array[itemIndex+1] {
			array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
			if itemIndex != 0 {
				itemIndex -= 1
			}
		} else {
			itemIndex += 1
		}
	}
}
