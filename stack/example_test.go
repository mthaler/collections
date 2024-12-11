package stack

func checkBalance(exp string) bool {
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

	return false
}
