package CircularBuffer

const arraySize = 10

type CircularBuffer struct {
	data    [arraySize]int
	pointer int
}

func (b *CircularBuffer) InsertValue(i int) {
	if b.pointer == len(b.data) {
		b.pointer = 0
	}
	b.data[b.pointer] = i
	b.pointer += 1
}

func (b *CircularBuffer) GetValues() [arraySize]int {
	return b.data
}

func (b *CircularBuffer) GetValuesFromPosition(i int) ([arraySize]int, bool) {
	var out [arraySize]int

	if i >= len(out) {
		return out, false
	}

	for u := 0; u < len(out); u++ {
		if i >= len(b.data) {
			i = 0
		}
		out[u] = b.data[i]
		i += 1
	}
	return out, true
}
