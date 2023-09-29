package main

import (
	"github.com/dnanam/hashitout/ds/stack"
)

// main creates a stack object and calls Push and Pop on the returned Stack object.
func main() {
	s := stack.NewStack()
	s.Push(1)
	s.Push(2)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
}
