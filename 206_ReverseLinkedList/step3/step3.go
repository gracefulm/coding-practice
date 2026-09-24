package step3

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode
	for node := head; node != nil; node = node.Next {
		reversed = &ListNode{Val: node.Val, Next: reversed}
	}
	return reversed
}
