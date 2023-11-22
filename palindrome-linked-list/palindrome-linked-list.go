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

func isPalindrome(head *ListNode) bool {
	middle := middleNode(head)

}
