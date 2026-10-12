package step1

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		prev *ListNode = &ListNode{Val: head.Val}
		next *ListNode = head.Next.Next
	)
	for node := head.Next; node != nil; {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}
