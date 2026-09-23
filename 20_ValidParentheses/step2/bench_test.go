package main

import (
	"fmt"
	"strings"
	"testing"
)

var sizes = []int{100, 10_000, 1_000_000}

// 入力の形によってスタックの伸び方と打ち切り位置が変わるため、代表的なパターンを用意する。
var cases = []struct {
	name string
	want bool
	gen  func(n int) string // 長さおよそ n の入力を生成する
}{
	// スタックが深さ 1 までしか伸びない。最も素直なケース。
	{"valid_flat", true, func(n int) string {
		return strings.Repeat("()", n/2)
	}},
	// スタックが n/2 まで伸びる。append の再確保が効いてくる。
	{"valid_nested", true, func(n int) string {
		return strings.Repeat("(", n/2) + strings.Repeat(")", n/2)
	}},
	// 3 種類の括弧が混ざる。map のキー分散・switch の case 分岐がばらける。
	{"valid_mixed", true, func(n int) string {
		return strings.Repeat("({[]})", n/6)
	}},
	// 先頭で即 false。早期リターンの速さを測る。
	{"invalid_head", false, func(n int) string {
		return ")" + strings.Repeat("()", n/2)
	}},
	// 最後の 1 文字で false。全体を走査しきる最悪ケース。
	{"invalid_tail", false, func(n int) string {
		return strings.Repeat("(", n/2) + strings.Repeat(")", n/2-1) + "]"
	}},
}

// ベンチの前に両実装が同じ答えを返すことを確認する。
func TestSameResult(t *testing.T) {
	for _, n := range sizes {
		for _, c := range cases {
			in := c.gen(n)
			if got := isValid_step2_map(in); got != c.want {
				t.Errorf("isValid_step2_map    n=%d %s: got %v, want %v", n, c.name, got, c.want)
			}
			if got := isValid_step2_switch(in); got != c.want {
				t.Errorf("isValid_step2_switch n=%d %s: got %v, want %v", n, c.name, got, c.want)
			}
		}
	}
}

// TestNonBracketInput は括弧以外の文字を含む入力に対する両実装の差を固定する。
// map 版は「開き括弧でない = 閉じ括弧」と決め打ちしているため対応できない。
func TestNonBracketInput(t *testing.T) {
	for _, in := range []string{"[(hoge)]", "a(b)c", "[(ほげ)]"} {
		if got := isValid_step2_switch(in); !got {
			t.Errorf("isValid_step2_switch(%q) = false, want true", in)
		}
		if got := isValid_step2_map(in); got {
			t.Errorf("isValid_step2_map(%q) = true; 括弧以外を扱えない既知の制限に変化あり", in)
		}
	}
}

var sink bool

// go test ./20_ValidParentheses/step2/ -bench . -benchmem
func BenchmarkIsValid(b *testing.B) {
	for _, n := range sizes {
		for _, c := range cases {
			in := c.gen(n)

			b.Run(fmt.Sprintf("n=%d/%s/map", n, c.name), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					sink = isValid_step2_map(in)
				}
			})
			b.Run(fmt.Sprintf("n=%d/%s/switch", n, c.name), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					sink = isValid_step2_switch(in)
				}
			})
		}
	}
}

// BenchmarkNonBracket は括弧以外の文字を含む入力を switch 版のみで測る。
// map 版は同じ入力を正しく扱えないため比較対象にならない。
func BenchmarkNonBracket(b *testing.B) {
	for _, n := range sizes {
		in := strings.Repeat("[(hoge)]", n/8)
		b.Run(fmt.Sprintf("n=%d/switch", n), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				sink = isValid_step2_switch(in)
			}
		})
	}
}
