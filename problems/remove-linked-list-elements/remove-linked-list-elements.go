package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func checkIfExists(head *ListNode, val int) bool {
	if head == nil {
		return false
	}

	for head != nil {
		if head.Val == val {
			return true
		}

		head = head.Next
	}

	return false
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummyHead := dummy

	for head != nil {
		if head.Val != val {
			dummy.Next = head
			dummy = head
		}

		head = head.Next
	}

	dummy.Next = nil

	return dummyHead.Next
}
