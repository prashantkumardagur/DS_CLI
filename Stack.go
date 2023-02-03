package main

import (
	"fmt"
)


// struct for stack
type Stack struct {
	array []item
	size int
	kind string
}
func NewStack(varkind string) *Stack {
	return &Stack{kind: varkind}
}
func (s *Stack) Kind() string { return s.kind }


// function to push a value into the stack
func (s *Stack) Push(value item) {
	s.array = append(s.array, value)
	s.size++
}


// function to pop a value from the stack
func (s *Stack) Pop() item {
	if s.size == 0 { return -1 }
	value := s.array[s.size-1]
	s.array = s.array[:s.size-1]
	s.size--
	return value
}


// function to get the top value of the stack
func (s *Stack) Top() item {
	if s.size == 0 { return -1 }
	return s.array[s.size-1]
}


// function to get the size of the stack
func (s *Stack) Size() int {
	return s.size
}


// function to check if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}


// function to print the stack
func (s *Stack) Print() {
	fmt.Printf("cli_out> Stack: Top ")
	for i := s.size-1; i >= 0; i-- {
		fmt.Printf("<- %v ", s.array[i])
	}
	fmt.Printf("<- Bottom\n")
}