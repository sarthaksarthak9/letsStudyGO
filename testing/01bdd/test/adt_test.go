package adt_test

import (
	"testing"

	adt "adt.go/code"
)

func TestSize(t *testing.T) {
	s := adt.NewStack()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Size() != 3 {
		t.Errorf("Expected: 3, found: %d", s.Size())
	}
}

func TestPop(t *testing.T) {
	s := adt.NewStack()
	s.Add(1)
	s.Add(2)

	v1 := s.Pop()
	if v1 != 2 {
		t.Errorf("Expexted: 2, found: %d", s.Pop())
	}

	v2 := s.Pop()
	if v2 != 1 {
		t.Errorf("Expected: 1, found: %d", s.Pop())
	}

}
