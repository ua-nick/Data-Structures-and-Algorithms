package RadixSort

import (
	"bytes"
	"encoding/binary"
)

const digit = 4
const maxbit = -1 << 31

func RadixSort(data []int) {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(data))
	for i, e := range data {
		binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		data[i] = int(w^maxbit)
	}
}
