package DoublyLinkedList

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (list *LinkedList) InsertFirst(i int) {
	data := &Node{data: i}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList) InsertLast(i int) {
	data := &Node{data: i}
	if list.head == nil {
		list.head = data
		list.tail = data
		return
	}
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.tail
	}
	list.tail = data
}

func (list *LinkedList) RemoveByValue(i int) bool {
	if list.head == nil {
		return false
	}
	if list.head.data == i {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if list.tail.data == i {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if current.next.data == i {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList) SearchValue(i int) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if current.data == i {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList) GetFirst() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	return list.head.data, true
}

func (list *LinkedList) GetLast() (int, bool) {
	if list.head == nil {
		return 0, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

func (list *LinkedList) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}

	return count
}

func (list *LinkedList) GetItemsFromStart() []int {
	var items []int
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *LinkedList) GetItemsFromEnd() []int {
	var items []int
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}
