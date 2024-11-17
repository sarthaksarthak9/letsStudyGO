package test

import (
	adt "testify/letsStudyGO/testing/02testify/code"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
	stack *adt.Stack // see SetupTest
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

func (s *StackSuite) SetupTest() {
	s.stack = adt.NewStack()
}

func (s *StackSuite) TestIsEmpty() {
	s.stack.Bury("red")
	s.False(s.stack.IsEmpty())
}

func (s *StackSuite) TestUnburySingleTest(){
	s.stack.Bury("blue")

	s.stack.Unbury()
	s.Zero(0, s.stack.Size())
	s.True(s.stack.IsEmpty())
}
func (s *StackSuite) TestBury(){
	s.stack.Bury("red")
	s.stack.Bury("yellow")

	s.Equal(2, s.stack.Size()) // Equal could compare any alphanumeric char
}

func (s *StackSuite) TestUnBury(){
	s.stack.Bury("red")
	s.stack.Bury("blue")

	s.Equal("red", s.stack.Unbury())
	s.Equal(1, s.stack.Size())
}


