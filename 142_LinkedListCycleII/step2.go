package main

// detectCycleFloyd は、Floydのアルゴリズムを使用して、リストの循環を検出します。
// 循環が見つかった場合は、循環の開始ノードを返します。
// 循環が見つからなかった場合は、nilを返します。
func detectCycleFloyd(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}

// detectCycleFloydV2 は、detectCycleFloyd の改良版です。
func detectCycleFloydV2(head *ListNode) *ListNode {
	collisionNode := findCollisionNode(head)
	if collisionNode == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next {
		if node == collisionNode {
			return node
		}
		collisionNode = collisionNode.Next
	}
	return nil
}

// findCollisionNode は、リストの循環を検出します。
// 循環が見つかった場合は、循環の衝突ノードを返します。
// 循環が見つからなかった場合は、nilを返します。
func findCollisionNode(head *ListNode) *ListNode {
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
