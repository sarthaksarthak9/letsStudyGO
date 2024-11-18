package adt

type Set struct {
	empty bool
	size  int
}

func NewSet() *Set {
	return &Set{true, 0}
}

func (s *Set) IsEmpty() bool {
	return s.size == 0
}

func (s *Set) ADD(v string) {

	s.size++
}

func (s *Set) Size() int {
	return s.size
}

func (s *Set) Contains(v string)bool{
	
	return s.IsEmpty()
}