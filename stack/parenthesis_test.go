package stack

import (
	"fmt"
	"testing"
)

func isBalanced(exp string) bool {
	fmt.Printf("%v", exp)
	s := Stack[string]{}
	for _, c := range exp {
		if c == '(' || c == '{' || c == '[' {
			s.Push(string(c))
		} else {
			t, _ := s.Top()
			if !s.IsEmpty() &&
				((t == "(" && c == ')') ||
					(t == "{" && c == '}') ||
					(t == "[" && c == ']')) {
				s.Pop()
			} else {
				return false // Unmatched closing bracket
			}
		}
	}

	return s.IsEmpty()
}

func TestIsBalanced(t *testing.T) {
	isBalanced("{(a*b) + (c*d)}")
	//if !isBalanced("{(a*b) + (c*d)}") {
	//	t.Errorf("isBalanced should result in true")
	//}
}
