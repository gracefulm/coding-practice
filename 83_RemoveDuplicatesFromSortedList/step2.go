package main

// deleteDuplicates_2 は、引数として受け取ったリストの重複する要素を削除する関数です。
// 元のリストを変更しません
func deleteDuplicates_2(head *ListNode) *ListNode {
	copied := deepCopy(head)
	for node := copied; node != nil; {
		if node.Next == nil {
			break
		}
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
			continue
		}
		node = node.Next
	}
	return copied
}

func deepCopy(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	newHead := &ListNode{Val: head.Val, Next: nil}
	prev := newHead
	for node := head.Next; node != nil; node = node.Next {
		n := &ListNode{Val: node.Val, Next: nil}
		prev.Next = n
		prev = n
	}
	return newHead
}

// deleteDuplicates_2V2 は、deleteDuplicates_2の改良版
// 元のリストを変更しません
func deleteDuplicates_2V2(head *ListNode) *ListNode {
	ln := &ListNode{}
	cursor := ln
	for node := head; node != nil; node = node.Next {
		if node.Next != nil && node.Val == node.Next.Val {
			continue // 同じ値が続くので、この要素は捨てて最後の1つだけ残す
		}
		cursor.Next = &ListNode{Val: node.Val}
		cursor = cursor.Next
	}
	return ln.Next
}

// deleteDuplicates_2_recursive は、再帰的にリストの重複する要素を削除する関数です。
func deleteDuplicates_2_recursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		head = deleteDuplicates_2_recursive(head.Next)
	} else {
		head.Next = deleteDuplicates_2_recursive(head.Next)
	}
	return head
}
