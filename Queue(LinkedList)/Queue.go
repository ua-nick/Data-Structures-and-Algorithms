package main

type Node struct {
	data int
	next *Node
}

type Queue struct {
	rear *Node
}

func (list *Queue) enqueue(i int) {
	data := &Node{data: i}
	if list.rear != nil {
		data.next = list.rear
	}
	list.rear = data
}

func (list *Queue) dequeue() (int, bool) {
	if list.rear == nil {
		return 0, false
	}
	if list.rear.next == nil {
		i := list.rear.data
		list.rear = nil
		return i, true
	}
	current := list.rear
	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, true
		}
		current = current.next
	}
}

func (list *Queue) peek() (int, bool) {
	if list.rear == nil {
		return 0, false
	}
	return list.rear.data, true
}

func (list *Queue) get() []int {
	var items []int
	current := list.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *Queue) isEmpty() bool {
	return list.rear == nil
}

func (list *Queue) empty() {
	list.rear = nil
}
