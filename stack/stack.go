package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(dara T) {
	s.items = append(s.items, dara)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var result T
		return result
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		var result T
		return result, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) ToSlice() []T {
	return s.items
}
