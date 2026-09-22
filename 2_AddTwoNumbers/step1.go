package main

func addTwoNumbers_1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	n1 := l1
	n2 := l2
	for n1 != nil || n2 != nil {
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
	if carry != 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
