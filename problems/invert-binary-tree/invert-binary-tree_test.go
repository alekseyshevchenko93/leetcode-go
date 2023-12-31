package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := TreeNode{Val: 4}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 7}
	node4 := TreeNode{Val: 1}
	node5 := TreeNode{Val: 3}
	node6 := TreeNode{Val: 6}
	node7 := TreeNode{Val: 9}

	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	node3.Left = &node6
	node3.Right = &node7

	root := invertTree(&node1)

	require.Equal(t, 7, root.Left.Val)
	require.Equal(t, 2, root.Right.Val)
	require.Equal(t, 9, root.Left.Left.Val)
	require.Equal(t, 3, root.Right.Left.Val)
	require.Equal(t, 1, root.Right.Right.Val)
}
