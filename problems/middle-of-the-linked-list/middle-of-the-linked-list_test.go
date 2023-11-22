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

	require.Equal(t, &node2, middleNode(&node1))
}

func TestProblem2(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	require.Equal(t, &node3, middleNode(&node1))
}

func TestProblem3(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	require.Equal(t, &node3, middleNode(&node1))
}
