package JumpSearch

import (
	"math/rand"
	"testing"
	"time"
)

func SortArray(array []int) {
	for itemIndex, itemValue := range array {
		for itemIndex != 0 && array[itemIndex-1] > itemValue {
			array[itemIndex] = array[itemIndex-1]
			itemIndex -= 1
		}
		array[itemIndex] = itemValue
	}
}

func TestJumpSearch(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, random.Intn(100-10)+10)
	for i := range array {
		array[i] = random.Intn(100)
	}
	SortArray(array)
	for _, value := range array {
		result := JumpSearch(array, value)
		if result == -1 {
			t.Fail()
		}
	}
}
