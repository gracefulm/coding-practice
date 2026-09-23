package main

var openToClose = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// isValid_step2_map は、mapを使って括弧のペアを管理する実装。
func isValid_step2_map(s string) bool {
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

// isValid_step2_switch は、switchを使って括弧のペアを管理する実装。
// "[(hoge)]"のような括弧以外の文字列を含むケースにも対応でき、実行時のパフォーマンスは、isValid_step2_mapよりも若干良い。
// ただし、コード量が多くなり、可読性が低下する可能性がある。
func isValid_step2_switch(s string) bool {
	st := newStack(len(s) / 2)
	for _, r := range s {
		switch r {
		case '(':
			st.push(')')
		case '[':
			st.push(']')
		case '{':
			st.push('}')
		case ')', ']', '}':
			if v, ok := st.pop(); !ok || v != r {
				return false
			}
		default:
			// skip other than parentheses
		}
	}
	return st.size() == 0
}
