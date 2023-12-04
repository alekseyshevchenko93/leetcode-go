package helpers

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalHelper(root *TreeNode, results []string) []string {
	if root == nil {
		return results
	}

	if root.Left != nil {
		results = inorderTraversalHelper(root.Left, results)
	} else {
		results = append(results, "null")
	}

	results = append(results, fmt.Sprintf("\"%s\"", fmt.Sprint(root.Val)))

	if root.Right != nil {
		results = inorderTraversalHelper(root.Right, results)
	} else {
		results = append(results, "null")
	}

	return results
}

func InorderTraversal(root *TreeNode) []string {
	results := make([]string, 0)
	return inorderTraversalHelper(root, results)
}

func PreorderTraversal(root *TreeNode) []string {
	results := make([]string, 0)
	return preorderTraversalHelper(root, results)
}

func preorderTraversalHelper(root *TreeNode, results []string) []string {
	if root == nil {
		return results
	}

	results = append(results, fmt.Sprintf("\"%s\"", fmt.Sprint(root.Val)))

	if root.Left != nil {
		results = preorderTraversalHelper(root.Left, results)
	} else {
		results = append(results, "null")
	}

	if root.Right != nil {
		results = preorderTraversalHelper(root.Right, results)
	} else {
		results = append(results, "null")
	}

	return results
}
