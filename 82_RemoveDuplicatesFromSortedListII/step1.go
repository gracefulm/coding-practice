package main

func deleteDuplicates_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		newHead    *ListNode
		newTail    *ListNode
		headExists bool
		shouldSkip bool
	)
	for current := head; current.Next != nil; current = current.Next {
		if current.Val == current.Next.Val {
			shouldSkip = true
			continue
		}
		if shouldSkip {
			shouldSkip = false
		} else {
			if !headExists {
				newTail = &ListNode{Val: current.Val, Next: nil}
				newHead = newTail
				headExists = true
			} else {
				newTail.Next = &ListNode{Val: current.Val, Next: nil}
				newTail = newTail.Next
			}
		}

		// ループの最後で、重複が発生していない場合は最後の要素を追加する
		if current.Next.Next == nil {
			if !headExists {
				return &ListNode{Val: current.Next.Val, Next: nil}
			}
			newTail.Next = &ListNode{Val: current.Next.Val, Next: nil}
		}
	}
	return newHead
}
