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
	node3.Next = &node1

	require.True(t, hasCycle(&node1))
}

func TestProblem2(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node1

	require.True(t, hasCycle2(&node1))
}
