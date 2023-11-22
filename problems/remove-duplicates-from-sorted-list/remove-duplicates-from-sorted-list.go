package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func allElementAreUnique(head *ListNode) bool {
	if head == nil {
		return true
	}

	m := make(map[int]struct{})

	for head != nil {
		_, found := m[head.Val]

		if found {
			return false
		}

		m[head.Val] = struct{}{}

		head = head.Next
	}

	return true
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Val: -101}
	dummyHead := dummy

	for head != nil {
		if dummy.Val != head.Val {
			dummy.Next = head
			dummy = head
		}

		head = head.Next
	}

	dummy.Next = nil

	return dummyHead.Next
}
