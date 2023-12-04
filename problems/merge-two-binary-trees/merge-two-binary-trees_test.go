package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	var (
		root1 *TreeNode
		root2 *TreeNode
		root3 *TreeNode
	)

	{
		node1 := TreeNode{Val: 1}
		node2 := TreeNode{Val: 3}
		node3 := TreeNode{Val: 2}
		node4 := TreeNode{Val: 5}
		node1.Left = &node2
		node1.Right = &node3
		node2.Left = &node4

		root1 = &node1
	}

	{
		node1 := TreeNode{Val: 2}
		node2 := TreeNode{Val: 1}
		node3 := TreeNode{Val: 3}
		node4 := TreeNode{Val: 4}
		node5 := TreeNode{Val: 7}

		node1.Left = &node2
		node1.Right = &node3
		node2.Right = &node4
		node3.Right = &node5

		root2 = &node1
	}

	{
		node1 := TreeNode{Val: 3}
		node2 := TreeNode{Val: 4}
		node3 := TreeNode{Val: 5}
		node4 := TreeNode{Val: 5}
		node5 := TreeNode{Val: 4}
		node6 := TreeNode{Val: 7}

		node1.Left = &node2
		node1.Right = &node3
		node2.Left = &node4
		node2.Right = &node5
		node3.Right = &node6

		root3 = &node1
	}

	result := mergeTrees(root1, root2)

	require.Equal(t, root3.Val, result.Val)
	require.Equal(t, root3.Left.Val, result.Left.Val)
	require.Equal(t, root3.Right.Val, result.Right.Val)
	require.Equal(t, root3.Left.Left.Val, result.Left.Left.Val)
	require.Equal(t, root3.Left.Right.Val, result.Left.Right.Val)
	require.Equal(t, root3.Right.Right.Val, result.Right.Right.Val)
}
