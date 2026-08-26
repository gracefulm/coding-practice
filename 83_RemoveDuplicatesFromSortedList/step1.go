package main

// deleteDuplicates_1 は、引数として受け取ったリストの重複する要素を削除する関数です。
// 引数の値を直接更新します。
func deleteDuplicates_1(head *ListNode) *ListNode {
	for node := head; node != nil; {
		if node.Next == nil {
			break
		}
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
	return head
}
