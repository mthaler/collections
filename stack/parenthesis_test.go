package stack

import (
	"testing"
)

func isBalanced(exp string) bool {
	s := Stack[string]{}

	for _, char := range exp {
		switch char {
		case '(':
			s.Push(string(char))
		case '{':
			s.Push(string(char))
		case '[':
			s.Push(string(char))
		case ')':
			top, _ := s.Top()
			if top == "(" {
				s.Pop()
			} else {
				return false
			}
		case '}':
			top, _ := s.Top()
			if top == "{" {
				s.Pop()
			} else {
				return false
			}
		case ']':
			top, _ := s.Top()
			if top == "[" {
				s.Pop()
			} else {
				return false
			}
		}
	}

	return s.IsEmpty()
}

func TestIsBalanced(t *testing.T) {
	if !isBalanced("{(a*b) + (c*d)}") {
		t.Errorf("isBalanced should result in true")
	}
	if isBalanced("(a+b") {
		t.Errorf("isBalanced should result in false")
	}
}
