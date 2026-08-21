package main

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		if _, seen := visited[node]; seen {
			return node
		}
		visited[node] = struct{}{}
	}
	return nil
}
