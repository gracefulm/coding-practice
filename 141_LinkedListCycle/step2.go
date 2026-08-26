package main

// Step1を改善したもの
// 時間計算量: O(n^2)
// 空間計算量: O(1)
func hasCycle_V2(head *ListNode) bool {
	for node := head; node != nil; node = node.Next {
		// node.Next がこの区間を指し返していれば循環している。
		for visited := head; ; visited = visited.Next {
			if node.Next == visited {
				return true
			}
			if visited == node { // 自分自身まで来たら未訪問領域に入る
				break
			}
		}
	}
	return false
}

// ハッシュセットで既訪問を記録する方法
// 時間計算量: O(n)
// 空間計算量: O(n)
func hasCycle_Map(head *ListNode) bool {
	visited := make(map[*ListNode]struct{})
	for node := head; node != nil; node = node.Next {
		if _, seen := visited[node]; seen {
			return true
		}
		visited[node] = struct{}{}
	}
	return false
}

// Floyd の循環検出
// 時間計算量: O(n)
// 空間計算量: O(1)
func hasCycle_Floyd(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Brent の循環検出
// 時間計算量: O(n)
// 空間計算量: O(1)
func hasCycle_Brent(head *ListNode) bool {
	if head == nil {
		return false
	}
	power, lambda := 1, 1
	tortoise, hare := head, head.Next
	for hare != nil && tortoise != hare {
		if power == lambda { // テレポート: 探索範囲を倍にする
			tortoise = hare
			power *= 2
			lambda = 0
		}
		hare = hare.Next
		lambda++
	}
	return hare != nil
}
