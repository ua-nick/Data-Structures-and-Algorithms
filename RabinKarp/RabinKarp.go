package RabinKarp

import "hash/adler32"

func calculateHash(text string) uint32 {
	return adler32.Checksum([]byte(text))
}

func RabinKarp(text, pattern string) int {
	textLength := len(text)
	patternLength := len(pattern)
	if patternLength > textLength {
		return -1
	}
	for i := 0; i < textLength-patternLength+1; i++ {
		if calculateHash(pattern) != calculateHash(text[i:i+patternLength]) {
			continue
		}
		matchesCount := 0
		for j := 0; j < patternLength; j++ {
			if text[i+j] != pattern[j] {
				break
			}
			matchesCount++
		}
		if matchesCount == patternLength {
			return i
		}
	}
	return -1
}
