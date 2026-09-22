package main

import "iter"

// addTwoNumbers_2 は2つのリストを加算し、結果をリストで返します。
// step1.goを改善したナイーブな実装です。
func addTwoNumbers_2(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	n1, n2 := l1, l2
	for n1 != nil || n2 != nil || carry > 0 {
		sum := carry
		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	return dummy.Next
}

// addTwoNumbers_2_recursive は2つのリストを加算し、結果をリストで返します。
// 再帰的な実装です。
func addTwoNumbers_2_recursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers_recursive_inner(l1, l2, 0)
}

func addTwoNumbers_recursive_inner(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	return &ListNode{Val: sum % 10, Next: addTwoNumbers_recursive_inner(l1, l2, sum/10)}
}

// addTwoNumbers_2_iter は2つのリストを加算し、結果をリストで返します。
// iter.Seq2を使用する実装です。
func addTwoNumbers_2_iter(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	for d1, d2 := range digitPairs(l1, l2) {
		sum := d1 + d2 + carry
		carry = sum / 10
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
	}
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// digitPairs は2つのリストの桁を下位から順に組で返します。短い方は 0 で埋めます。
func digitPairs(l1, l2 *ListNode) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for n1, n2 := l1, l2; n1 != nil || n2 != nil; {
			d1, d2 := 0, 0
			if n1 != nil {
				d1 = n1.Val
				n1 = n1.Next
			}
			if n2 != nil {
				d2 = n2.Val
				n2 = n2.Next
			}
			if !yield(d1, d2) {
				return
			}
		}
	}
}
