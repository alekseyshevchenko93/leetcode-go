package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := TreeNode{Val: 3}
	node2 := TreeNode{Val: 9}
	node3 := TreeNode{Val: 20}
	node4 := TreeNode{Val: 15}
	node5 := TreeNode{Val: 7}

	node1.Left = &node2
	node1.Right = &node3
	node3.Left = &node4
	node3.Right = &node5

	require.Equal(t, 2, minDepth(&node1))
}

func TestProblem2(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5}

	node1.Right = &node2
	node2.Right = &node3
	node3.Right = &node4
	node4.Right = &node5

	require.Equal(t, 5, minDepth(&node1))
}
