package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	results := make([]string, 0)
	var helper func(*TreeNode, string)

	helper = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			results = append(results, path)
			return
		}

		if node.Left != nil {
			helper(node.Left, fmt.Sprintf("%s->%s", path, fmt.Sprint(node.Left.Val)))
		}

		if node.Right != nil {
			helper(node.Right, fmt.Sprintf("%s->%s", path, fmt.Sprint(node.Right.Val)))
		}
	}

	helper(root, fmt.Sprint(root.Val))

	fmt.Println(results)

	return results
}
