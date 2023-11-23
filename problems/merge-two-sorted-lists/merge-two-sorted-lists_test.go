package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}

	node1.Next = &node2
	node2.Next = &node3

	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}

	node4.Next = &node5

	root := mergeTwoLists(&node1, &node4)

	for i := 1; i <= 5; i++ {
		require.NotNil(t, root)
		require.Equal(t, i, root.Val)
		root = root.Next
	}
}
