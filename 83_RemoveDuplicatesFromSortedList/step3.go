package main

func deleteDuplicates_3_normal(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	for current := head.Next; current != nil; current = current.Next {
		if prev.Val == current.Val {
			prev.Next = current.Next
			continue
		}
		prev = current
	}
	return head
}

func deleteDuplicates_3_withoutSideEffect(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val, Next: nil}
	prev := newHead
	for current := head.Next; current != nil; current = current.Next {
		if prev.Val != current.Val {
			prev.Next = &ListNode{Val: current.Val, Next: nil}
			prev = prev.Next
		}
	}
	return newHead
}

func deleteDuplicates_3_recursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		head = deleteDuplicates_3_recursive(head.Next)
	} else {
		head.Next = deleteDuplicates_3_recursive(head.Next)
	}
	return head
}
