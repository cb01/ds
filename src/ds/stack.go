package main 

import (
	"fmt"
)

// Stack is an object implementing a classic LIFO stack data structure.
type Stack struct {
	head *stackNode
}

// Stack
type stackNode struct {
	next *stackNode
	data int
}

// NewStack creates and returns an empty stack data structure.
func NewStack() Stack {
	return Stack{}
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(val int) error {
	
	// Create a new node containing the pushed value and a pointer
	// to the current head node. It is important to point to the
	// current head before setting s.head because keeping the current
	// stack nodes in memory depends on there not being a gap in them
	// being pointed to by something.
	n := stackNode{s.head, val}

	// Set the current head of the stack to point to the new node.
	s.head = &n

	return nil
}

// Pop removes and returns the value from the top of the stack.
//
// A non-nil error will be returned if Pop is called on an empty stack.
func (s *Stack) Pop() (int, error) {
	
	// If the stack is empty, return an error
	if s.head == nil {
		return -1, fmt.Errorf("can't pop a value from an empty stack")
	}

	// Get the value that will be returned
	val := (*s.head).data

	// If the head node of the stack is the last
	if (*s.head).next == nil {
		s.head = (*s.head).next
		return val, nil
	}

	// Since the head node is not the last node in the stack, the head
	// node is set to simply the next node in the stack beyond the 
	// head node.
	s.head = (*s.head).next

	return val, nil
}

