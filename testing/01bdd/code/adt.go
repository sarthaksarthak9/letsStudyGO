package adt

type Stack struct {
	isEmpty bool
	size    int
	value   int
}

func NewStack() *Stack {
	return &Stack{true, 0, -1}
}

func (s *Stack) Add(value int) {
	s.isEmpty = false
	s.size++
	s.value = value
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Pop() int {
	s.size--
	return s.value
}
