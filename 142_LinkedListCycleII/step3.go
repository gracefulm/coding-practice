package main

func detectCycle_3(head *ListNode) *ListNode {
	collisonNode := findCollision(head)
	if collisonNode == nil {
		return nil
	}

	for node := head; ; node = node.Next {
		if collisonNode == node {
			return collisonNode
		}
		collisonNode = collisonNode.Next
	}
	return nil
}

func findCollision(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return fast
		}
	}
	return nil
}
