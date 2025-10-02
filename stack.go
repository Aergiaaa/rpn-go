package main

import "fmt"

type Stack struct {
	size  int
	top   int
	items []any
}

func StackInit(size int) Stack {
	return Stack{
		size:  size,
		top:   -1,
		items: make([]any, size),
	}
}

func (s Stack) IsFull() bool {
	return s.top == s.size
}

func (s Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(item any) (any, error) {
	if s.IsFull() {
		return nil, fmt.Errorf("error: stack already full")
	}

	s.items[s.top+1] = item
	s.top++

	fmt.Println("Pushed item:", item) // Debugging line
	fmt.Println("Top: ", s.top)

	return item, nil
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("error: stack is empty")
	}

	item := s.items[s.top]
	s.items[s.top] = nil // Clear the popped item
	s.top--

	fmt.Println("Popped item:", item) // Debugging line
	fmt.Println("Top: ", s.top)

	return item, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("error: stack is empty")
	}

	return s.items[s.top], nil
}

func (s *Stack) PeekAt(index int) (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("error: stack is empty")
	}

	if index >= s.top {
		return nil, fmt.Errorf("error: stack item are not that many")
	}

	return s.items[index], nil
}

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
