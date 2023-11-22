package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 2}
	node4 := ListNode{Val: 1}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	require.True(t, isPalindrome(&node1))
}

func TestProblem2(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}

	node1.Next = &node2

	require.False(t, isPalindrome(&node1))
}

func TestProblem3(t *testing.T) {
	node1 := ListNode{Val: 1}

	require.True(t, isPalindrome(&node1))
}

func TestProblem4(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 2}
	node5 := ListNode{Val: 1}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	require.True(t, isPalindrome(&node1))
}
