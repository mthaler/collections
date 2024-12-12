package stack

import (
	"fmt"
	"testing"
)

func isBalanced(exp string) bool {
	fmt.Printf("%v", exp)
	s := Stack[rune]{}
	for _, c := range exp {
		if c == '(' || c == '{' || c == '[' {
			fmt.Printf("Pushed %s", string(c))
			s.Push(c)
		} else {
			t, _ := s.Top()
			if !s.IsEmpty() &&
				((t == '(' && c == ')') ||
					(t == '{' && c == '}') ||
					(t == '[' && c == ']')) {
				r := s.Pop()
				fmt.Printf("popped: %s", string(r))
			} else {
				t, _ := s.Top()
				fmt.Printf("popped: %s", string(t))
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
