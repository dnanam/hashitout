package stack

import "fmt"

type Stack struct {
	stack []int
}

// NewStack returns a Stack object
func NewStack() *Stack {
	return &Stack{
		stack: []int{},
	}
}

// Push will add a value to the end
func (s *Stack) Push(x int) {
	s.stack = append(s.stack, x)
}

// Pop removes the last element that was added to the stack
func (s *Stack) Pop() {
	// wrap in condition to check if stack is not empty
	if len(s.stack) != 0 {
		s.stack = s.stack[:len(s.stack)-1]
	} else {
		fmt.Println("the stack seems to be empty")
	}
}
