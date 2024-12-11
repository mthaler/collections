package stack

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(dara T) {
	s.items = append(s.items, dara)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpy() {
		var result T
		return result
	}
	result := s.items[len(s.items)]
	s.items = s.items[:len(s.items)-1]
	return result
}

func (s *Stack[T]) IsEmpy() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) ToSliice() []T {
	return s.items
}
