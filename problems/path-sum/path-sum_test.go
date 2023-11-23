package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := TreeNode{Val: 5}
	node2 := TreeNode{Val: 4}
	node3 := TreeNode{Val: 8}
	node4 := TreeNode{Val: 11}
	node5 := TreeNode{Val: 7}
	node6 := TreeNode{Val: 2}
	node7 := TreeNode{Val: 13}
	node8 := TreeNode{Val: 4}
	node9 := TreeNode{Val: 1}

	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node4.Right = &node5
	node4.Right = &node6
	node3.Left = &node7
	node3.Right = &node8
	node8.Right = &node9

	require.Equal(t, true, hasPathSum(&node1, 22))
}

func TestProblem2(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}

	node1.Left = &node2
	node1.Right = &node3

	require.Equal(t, false, hasPathSum(&node1, 5))
}

func TestProblem3(t *testing.T) {
	require.Equal(t, false, hasPathSum(nil, 0))
}

func TestProblem4(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}

	node1.Left = &node2

	require.Equal(t, false, hasPathSum(&node1, 1))
}
