package HeapSort

type Heap struct {
}

func (heap *Heap) HeapSort(array []int) {
	heap.BuildHeap(array)

	for length:= len(array); length > 1; length-- {
		heap.RemoveTop(array, length)
	}
}

func (heap *Heap) BuildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap.Heapify(array, i, len(array))
	}
}

func (heap *Heap) RemoveTop(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

func (heap *Heap) Heapify(array []int, root, length int) {
	var max = root
	var l, r = heap.Left(array, root), heap.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		heap.Heapify(array, max, length)
	}
}

func (*Heap) Left(array []int, root int) int {
	return (root * 2) + 1
}

func (*Heap) Right(array []int, root int) int {
	return (root * 2) + 2
}
