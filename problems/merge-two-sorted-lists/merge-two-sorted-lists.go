package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummyHead := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			dummy.Next = list1
			dummy = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			dummy = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		dummy.Next = list1
	}

	if list2 != nil {
		dummy.Next = list2
	}

	return dummyHead.Next
}
