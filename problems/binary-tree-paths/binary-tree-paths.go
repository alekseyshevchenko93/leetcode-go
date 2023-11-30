package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(root *TreeNode, current string, results []string) []string {
	if root == nil {
		return results
	}

	if root.Left == nil && root.Right == nil {
		return append(results, current)
	}

	if root.Left != nil {
		results = helper(root.Left, fmt.Sprintf("%s->%s", current, fmt.Sprint(root.Left.Val)), results)
	}

	if root.Right != nil {
		results = helper(root.Right, fmt.Sprintf("%s->%s", current, fmt.Sprint(root.Right.Val)), results)
	}

	return results
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	results := make([]string, 0)
	return helper(root, fmt.Sprint(root.Val), results)
}
