package NaiveStringSearch

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNaiveStringSearch(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	text := make([]rune, random.Intn(15-5)+5)
	for i := range text {
		text[i] = letters[rand.Intn(len(letters))]
	}
	end := random.Intn(len(text)-5) + 5
	start := random.Intn(end)
	fmt.Println(len(text), end, start)
	result := NaiveStringSearch(string(text), string(text[start:end]))
	if result == -1 {
		t.Fail()
	}
}
