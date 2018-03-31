package main

type Node struct {
	data int
	next *Node
}

type Stack struct {
	top *Node
}

func (list *Stack) push(i int) {
	data := &Node{data: i}
	if list.top != nil {
		data.next = list.top
	}
	list.top = data
}

func (list *Stack) pop() (int, bool) {
	if list.top == nil {
		return 0, false
	}
	i := list.top.data
	list.top = list.top.next
	return i, true
}

func (list *Stack) peek() (int, bool) {
	if list.top == nil {
		return 0, false
	}
	return list.top.data, true
}

func (list *Stack) get() []int {

	var items []int

	current := list.top
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *Stack) isEmpty() bool {
	return list.top == nil
}

func (list *Stack) empty() {
	list.top = nil
}
