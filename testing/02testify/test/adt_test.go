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
