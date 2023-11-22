package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 -> 2 -> 3
func reverseList(head *ListNode) *ListNode {
	var last *ListNode = nil
	node := head

	for node != nil {
		next := node.Next
		node.Next = last
		last = node
		node = next
	}

	return last
}
