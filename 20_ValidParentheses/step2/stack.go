package main

type stack struct {
	closeBrackets []rune
}

func newStack(size int) *stack {
	return &stack{
		closeBrackets: make([]rune, 0, size),
	}
}

func (s *stack) size() int {
	return len(s.closeBrackets)
}

func (s *stack) push(closeBracket rune) {
	s.closeBrackets = append(s.closeBrackets, closeBracket)
}

func (s *stack) pop() (rune, bool) {
	if s.size() == 0 {
		return -1, false
	}

	r := s.closeBrackets[s.size()-1]
	s.closeBrackets = s.closeBrackets[:s.size()-1]
	return r, true
}
