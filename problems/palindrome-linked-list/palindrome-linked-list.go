package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	p1 := head
	p2 := head

	for p1.Next != nil && p1.Next.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next
	}

	if p1.Next != nil {
		return p2.Next
	}

	return p2
}

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

func isPalindrome(head *ListNode) bool {
	count := 0
	node := head

	for node != nil {
		count++
		node = node.Next
	}

	middle := middleNode(head)

	left := head
	var right *ListNode

	if count%2 == 0 {
		right = reverseList(middle)
	} else {
		right = reverseList(middle.Next)
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}
