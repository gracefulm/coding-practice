package main

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	current := head
	for {
		// Nextがnilの場合は末尾 && 循環しないため false
		if current.Next == nil {
			return false
		}

		current2 := head
		for {
			// Nextが他のリストの要素をさしている場合は循環しているとみなす
			if current.Next == current2 {
				return true
			}

			// 自分自身までたどり着いたら親のループに戻る
			if current == current2 {
				break
			}

			// current2を1つ進める
			current2 = current2.Next
		}

		// currentを1つ進める
		current = current.Next
	}
}
