package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInorderTraversal(t *testing.T) {
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

	require.Equal(t, []int{1, 2, 3, 4, 6, 7, 9}, InorderTraversal(&node1))
}
