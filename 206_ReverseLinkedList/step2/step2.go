package step2

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList は、head を破壊的に変更して逆順にする。
// Step1の改良版。
func reverseList(head *ListNode) *ListNode {
	var reversed *ListNode
	for head != nil {
		next := head.Next
		head.Next = reversed
		reversed = head
		head = next
	}
	return reversed
}

// reverseListWithoutSideEffect は、head を破壊的に変更せずに逆順にする。
func reverseListWithoutSideEffect(head *ListNode) *ListNode {
	var reversed *ListNode
	for node := head; node != nil; node = node.Next {
		reversed = &ListNode{Val: node.Val, Next: reversed}
	}
	return reversed
}

// reverseListRecursive は、再帰で逆順にする。
// この実装は、head を破壊的に変更します。
func reverseListRecursive(head *ListNode) *ListNode {
	return reverse(head, nil)
}

func reverse(head *ListNode, reversed *ListNode) *ListNode {
	if head == nil {
		return reversed
	}
	next := head.Next
	head.Next = reversed
	reversed = head
	return reverse(next, head)
}
