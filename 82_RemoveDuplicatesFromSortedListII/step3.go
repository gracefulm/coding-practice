package main

// deleteDuplicates_2_withSideEffectは、引数のリンクトリストの重複を排除して返します。
// ただし、引数のリンクトリストを変更します。
// 空間計算量: O(1)
// 時間計算量: O(n)
func deleteDuplicates_3_withSideEffect(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
    tail := dummy
    for current := head; current != nil; {
        next := current.Next
        for next != nil && current.Val == next.Val {
            next = next.Next
        }
        if current.Next == next {
            tail.Next = current
            tail = current
        }
        current = next
    }
	tail.Next = nil
    return dummy.Next
}

// deleteDuplicates_2_withoutSideEffectは、引数のリンクトリストの重複を排除して返します。
// ただし、引数のリンクトリストを変更しません。
// 空間計算量: O(n)
// 時間計算量: O(n)
func deleteDuplicates_3_withoutSideEffect(head *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for current := head; current != nil; {
        next := current.Next
        for next != nil && current.Val == next.Val {
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
