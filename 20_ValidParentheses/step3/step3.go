package main

type stack struct {
	closeParentheses []rune
}

func newStack(size int) *stack {
	return &stack{
		closeParentheses: make([]rune, 0, size),
	}
}

func (s *stack) size() int {
	return len(s.closeParentheses)
}

func (s *stack) push(r rune) {
	s.closeParentheses = append(s.closeParentheses, r)
}

func (s *stack) pop() (rune, bool) {
	if s.size() == 0 {
		return -1, false
	}
	r := s.closeParentheses[s.size()-1]
	s.closeParentheses = s.closeParentheses[:s.size()-1]
	return r, true
}

var openToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

func isValid(s string) bool {
	st := newStack(len(s) / 2)
	for _, r := range s {
		if v, ok := openToClose[r]; ok {
			st.push(v)
			continue
		}
		if expected, ok := st.pop(); !ok || expected != r {
			return false
		}
	}
	return st.size() == 0
}
