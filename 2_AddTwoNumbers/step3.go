package main

// addTwoNumbers_3 は2つのリストを加算し、結果をリストで返します。
func addTwoNumbers_3(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	for n1,n2 := l1,l2; n1 != nil || n2 != nil || carry != 0; {
		sum := carry
		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}
		carry = sum/10
		tail.Next = &ListNode{Val: sum%10}
		tail = tail.Next
	}
	return dummy.Next
}
