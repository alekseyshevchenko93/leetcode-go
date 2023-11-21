package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	m := make(map[*ListNode]struct{})

	for head != nil {
		_, found := m[head]

		if found {
			return true
		}

		m[head] = struct{}{}
		head = head.Next
	}

	return false
}
