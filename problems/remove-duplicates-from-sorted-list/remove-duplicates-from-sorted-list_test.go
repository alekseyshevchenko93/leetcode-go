package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 2}
	node4 := ListNode{Val: 3}
	node5 := ListNode{Val: 3}
	node6 := ListNode{Val: 4}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	head := deleteDuplicates(&node1)

	require.True(t, allElementAreUnique(head))
}

func TestProblem2(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 1}
	node3 := ListNode{Val: 2}
	node4 := ListNode{Val: 3}
	node5 := ListNode{Val: 3}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	head := deleteDuplicates(&node1)

	require.True(t, allElementAreUnique(head))
}

func TestProblem3(t *testing.T) {
	node1 := ListNode{Val: 0}
	node2 := ListNode{Val: 0}
	node3 := ListNode{Val: 0}
	node4 := ListNode{Val: 0}
	node5 := ListNode{Val: 0}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	head := deleteDuplicates(&node1)

	require.True(t, allElementAreUnique(head))
}
