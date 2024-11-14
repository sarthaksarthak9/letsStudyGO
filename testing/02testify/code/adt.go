package adt

type Stack struct {
	size int
	item [6]string
}

func NewStack() *Stack {
	return &Stack{0, [6]string{}}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Bury(item string) {
	s.item[s.size] = item
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Unbury() string {
	s.size--
	return s.item[s.size]
}
