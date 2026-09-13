package main

type stack struct {
	inner []string
}

func (s *stack) size() int {
	return len(s.inner)
}

func (s *stack) push(char string) {
	s.inner = append(s.inner, char)
}

func (s *stack) pop() *string {
	if s.size() == 0 {
		return nil
	}
	char := s.inner[s.size()-1]
	s.inner = s.inner[:s.size()-1]
	return &char
}

func isValid(s string) bool {
	bracketPairs := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	st := new(stack)
	for _, r := range s {
		char := string(r)
		val, ok := bracketPairs[char]
		if !ok {
			if expected := st.pop(); expected == nil || *expected != char {
				return false
			}
			continue
		}
		st.push(val)
	}
	return st.size() == 0
}
