package RabinKarp

import (
	"math/rand"
	"testing"
	"time"
)

func TestRabinKarp(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	text := make([]rune, random.Intn(15-10)+10)
	for i := range text {
		text[i] = letters[rand.Intn(len(letters))]
	}
	end := random.Intn(len(text)-5) + 5
	start := random.Intn(end)
	result := RabinKarp(string(text), string(text[start:end]))
	if result == -1 {
		t.Fail()
	}
}
