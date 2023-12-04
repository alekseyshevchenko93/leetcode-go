package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := TreeNode{Val: 3}
	node2 := TreeNode{Val: 4}
	node3 := TreeNode{Val: 5}
	node4 := TreeNode{Val: 1}
	node5 := TreeNode{Val: 2}

	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node2.Right = &node5

	node6 := TreeNode{Val: 4}
	node7 := TreeNode{Val: 1}
	node8 := TreeNode{Val: 2}
	node6.Left = &node7
	node6.Right = &node8

	require.True(t, isSubtree(&node1, &node6))
}

func TestProblem2(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}

	node1.Left = &node2
	node1.Right = &node3

	node6 := TreeNode{Val: 1}
	node7 := TreeNode{Val: 2}
	node6.Left = &node7

	require.True(t, isSubtree(&node1, &node6))
}

func TestProblem3(t *testing.T) {
	node1 := TreeNode{Val: 12}
	node2 := TreeNode{Val: 2}

	require.False(t, isSubtree(&node1, &node2))
}
