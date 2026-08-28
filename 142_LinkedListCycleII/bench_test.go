package main

import (
	"fmt"
	"testing"
)

// buildList は n 個のノードを繋ぎ、pos >= 0 なら末尾を nodes[pos] に繋いで循環を作る。
// 戻り値は (先頭, 期待される循環の入口)。
func buildList(n, pos int) (*ListNode, *ListNode) {
	nodes := make([]*ListNode, n)
	for i := range nodes {
		nodes[i] = &ListNode{Val: i}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if pos >= 0 {
		nodes[n-1].Next = nodes[pos]
		return nodes[0], nodes[pos]
	}
	return nodes[0], nil
}

var sizes = []int{100, 10_000, 1_000_000}

var cases = []struct {
	name string
	pos  func(n int) int
}{
	{"cycle_head", func(n int) int { return 0 }},     // リスト全体が循環
	{"cycle_mid", func(n int) int { return n / 2 }},  // 前半は直線、後半が循環
	{"cycle_tail", func(n int) int { return n - 1 }}, // 末尾の自己ループ（Floyd の最悪ケース）
	{"no_cycle", func(n int) int { return -1 }},      // 循環なし（HashMap の最悪ケース）
}

// ベンチの前に両実装が同じ答えを返すことを確認する。
func TestSameResult(t *testing.T) {
	for _, n := range sizes {
		for _, c := range cases {
			head, want := buildList(n, c.pos(n))
			if got := detectCycle(head); got != want {
				t.Errorf("detectCycle   n=%d %s: got %p, want %p", n, c.name, got, want)
			}
			if got := detectCycleFloyd(head); got != want {
				t.Errorf("detectCycleV2 n=%d %s: got %p, want %p", n, c.name, got, want)
			}
		}
	}
}

var sink *ListNode

// go test ./142_LinkedListCycleII/ -bench . -benchmem
func BenchmarkDetectCycle(b *testing.B) {
	for _, n := range sizes {
		for _, c := range cases {
			head, _ := buildList(n, c.pos(n))

			b.Run(fmt.Sprintf("n=%d/%s/HashMap", n, c.name), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					sink = detectCycle(head)
				}
			})
			b.Run(fmt.Sprintf("n=%d/%s/Floyd", n, c.name), func(b *testing.B) {
				b.ReportAllocs()
				for b.Loop() {
					sink = detectCycleFloyd(head)
				}
			})
		}
	}
}
