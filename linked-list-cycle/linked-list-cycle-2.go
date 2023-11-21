package main

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	p1 := head
	p2 := head

	for p1 != nil || p1.Next != nil {
		p1 = p1.Next.Next
		p2 = p2.Next

		if p1 == p2 {
			return true
		}
	}

	return false
}
