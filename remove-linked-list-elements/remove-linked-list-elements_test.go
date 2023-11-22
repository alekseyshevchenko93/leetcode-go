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

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	head := removeElements(&node1, 2)

	require.False(t, checkIfExists(head, 2))
}

func TestProblem2(t *testing.T) {
	head := removeElements(nil, 2)

	require.False(t, checkIfExists(head, 2))
}

func TestProblem3(t *testing.T) {
	node1 := ListNode{Val: 7}
	node2 := ListNode{Val: 7}
	node3 := ListNode{Val: 7}
	node4 := ListNode{Val: 7}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	head := removeElements(&node1, 7)

	require.False(t, checkIfExists(head, 7))
}

func TestProblem4(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node6 := ListNode{Val: 6}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	head := removeElements(&node1, 6)

	require.False(t, checkIfExists(head, 6))
}
