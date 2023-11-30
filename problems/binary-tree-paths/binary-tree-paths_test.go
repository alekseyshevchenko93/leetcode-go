package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 5}

	node1.Left = &node2
	node1.Right = &node3
	node2.Right = &node4

	require.Equal(t, []string{
		"1->2->5", "1->3",
	}, binaryTreePaths(&node1))
}

func TestProblem2(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}

	node1.Left = &node2
	node1.Right = &node3
	node3.Left = &node4

	require.Equal(t, []string{
		"1->2", "1->3->4",
	}, binaryTreePaths(&node1))
}
