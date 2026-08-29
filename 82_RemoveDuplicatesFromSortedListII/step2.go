package main

// deleteDuplicates_2_withSideEffectは、引数のリンクトリストの重複を排除して返します。
// ただし、引数のリンクトリストを変更します。
// 空間計算量: O(1)
// 時間計算量: O(n)
func deleteDuplicates_2_withSideEffect(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	lastKept := dummy
	for current := head; current != nil; {
		// current と同じ値が続く区間を読み飛ばす
		next := current.Next
		for next != nil && next.Val == current.Val {
			next = next.Next
		}
		if current.Next == next { // 区間の長さが1、つまり重複していない
			lastKept.Next = current
			lastKept = current
		}
		current = next
	}
	lastKept.Next = nil
	return dummy.Next
}

// deleteDuplicates_2_withoutSideEffectは、引数のリンクトリストの重複を排除して返します。
// ただし、引数のリンクトリストを変更しません。
// 空間計算量: O(n)
// 時間計算量: O(n)
func deleteDuplicates_2_withoutSideEffect(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for current := head; current != nil; {
		next := current.Next
		for next != nil && next.Val == current.Val {
			next = next.Next
		}
		if current.Next == next {
			tail.Next = &ListNode{Val: current.Val}
			tail = tail.Next
		}
		current = next
	}
	return dummy.Next
}
