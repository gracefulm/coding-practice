package main

func hasCycle_V3(head *ListNode) bool {
	visited := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		if _, seen := visited[node]; seen {
			return true
		}
		visited[node] = struct{}{}
	}
	return false
}
