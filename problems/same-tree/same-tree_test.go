package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node1.Left = &node2
	node1.Right = &node3

	node4 := TreeNode{Val: 1}
	node5 := TreeNode{Val: 2}
	node6 := TreeNode{Val: 3}
	node4.Left = &node5
	node4.Right = &node6

	require.True(t, isSameTree(&node1, &node4))
}

func TestProblem2(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 1}
	node1.Left = &node2
	node1.Right = &node3

	node4 := TreeNode{Val: 1}
	node5 := TreeNode{Val: 2}
	node6 := TreeNode{Val: 1}
	node4.Left = &node6
	node4.Right = &node5

	require.False(t, isSameTree(&node1, &node2))
}

func TestProblem3(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node1.Left = &node2

	node4 := TreeNode{Val: 1}
	node5 := TreeNode{Val: 2}
	node4.Right = &node5

	require.False(t, isSameTree(&node1, &node2))
}
