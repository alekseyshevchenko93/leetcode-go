package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPreorderTraversal(t *testing.T) {
	node1 := TreeNode{Val: 1}
	node2 := TreeNode{Val: 2}
	node3 := TreeNode{Val: 3}

	node1.Left = &node2
	node1.Right = &node3

	require.Equal(t, []string{"\"1\"", "\"2\"", "null", "null", "\"3\"", "null", "null"}, PreorderTraversal(&node1))
}
