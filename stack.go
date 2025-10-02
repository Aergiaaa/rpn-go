package main

import "fmt"

// Stack is a simple stack implementation in Go.
// It supports basic operations like Push, Pop, Peek, and Search.
// The stack is implemented using a slice to hold the items.
// It also includes methods to check if the stack is full or empty.
type Stack struct {
	size  int
	top   int
	items []any
}

// StackInit initializes a new stack with a given size.
// It sets the top index to -1, indicating that the stack is empty.
// The items slice is allocated with the specified size.
func StackInit(size int) Stack {
	return Stack{
		size:  size,
		top:   -1,
		items: make([]any, size),
	}
}

// IsFull checks if the stack is full.
func (s Stack) IsFull() bool {
	return s.top == s.size
}

// IsEmpty checks if the stack is empty.
func (s Stack) IsEmpty() bool {
	return s.top == -1
}

// Push adds an item to the top of the stack.
// If the stack is full, it returns an error.
// Otherwise, it increments the top index and adds the item to the stack.
func (s *Stack) Push(item any) (any, error) {
	if s.IsFull() {
		return nil, fmt.Errorf("error: stack already full")
	}

	s.items[s.top+1] = item
	s.top++

	return item, nil
}

// Pop removes and returns the item at the top of the stack.
// If the stack is empty, it returns an error.
// Otherwise, it retrieves the item at the top index, clears that position,
// and decrements the top index.
func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("error: stack is empty")
	}

	item := s.items[s.top]
	s.items[s.top] = nil // Clear the popped item
	s.top--

	return item, nil
}

// Peek returns the item at the top of the stack without removing it.
// If the stack is empty, it returns an error.
// Otherwise, it retrieves the item at the top index.
func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("error: stack is empty")
	}

	return s.items[s.top], nil
}

// PeekAt returns the item at a specific index in the stack without removing it.
// If the stack is empty or the index is out of bounds, it returns an error.
// Otherwise, it retrieves the item at the specified index.
// Note: The index is zero-based, so the top item is at index 0.
func (s *Stack) PeekAt(index int) (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("error: stack is empty")
	}

	if index >= s.top {
		return nil, fmt.Errorf("error: stack item are not that many")
	}

	return s.items[index], nil
}

// Search looks for an item in the stack and returns its index.
// If the stack is empty, it returns an error.
// If the item is found, it returns the index; otherwise, it returns an error indicating that the item is not found.
func (s *Stack) Search(item any) (i int, err error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("error: stack is empty")
	}

	for i = s.top - 1; i >= 0; i-- {
		if item == s.items[i] {
			return i, nil
		}
	}

	return -1, fmt.Errorf("error: there is no such item in this stack")
}
