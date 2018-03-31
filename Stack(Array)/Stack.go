package main

const arraySize = 10

type Stack struct {
	top  int
	data [arraySize]int
}

func (s *Stack) push(i int) bool {
	if s.top == len(s.data) {
		return false
	}
	s.data[s.top] = i
	s.top += 1
	return true
}

func (s *Stack) pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	}
	i := s.data[s.top-1]
	s.top -= 1
	return i, true
}

func (s *Stack) peek() int {
	return s.data[s.top-1]
}

func (s *Stack) get() []int {
	return s.data[:s.top]
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) empty() {
	s.top = 0
}
