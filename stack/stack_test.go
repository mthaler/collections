package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	var s Stack[int]
	if !s.IsEmpy() {
		t.Errorf("Stack should be empty")
	}
	s.Push(1)
	if s.IsEmpy() {
		t.Errorf("Stack should be empty")
	}
}
