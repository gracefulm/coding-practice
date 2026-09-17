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

func (s *stack) push(r rune) {
	s.closeBrackets = append(s.closeBrackets, r)
}

func (s *stack) pop() (rune, bool) {
	if s.size() == 0 {
		return -1, false
	}
	r := s.closeBrackets[s.size()-1]
	s.closeBrackets = s.closeBrackets[:s.size()-1]
	return r, true
}

var openToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	closeBrackets := newStack(len(s) / 2)
	for _, r := range s {
		if v, ok := openToClose[r]; ok {
			closeBrackets.push(v)
			continue
		}
		if expected, ok := closeBrackets.pop(); !ok || expected != r {
			return false
		}
	}
	return closeBrackets.size() == 0
}
