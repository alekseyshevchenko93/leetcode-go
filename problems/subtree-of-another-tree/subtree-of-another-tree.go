package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, results []string) []string {
	if root == nil {
		return results
	}

	results = append(results, fmt.Sprintf("\"%s\"", fmt.Sprint(root.Val)))

	if root.Left != nil {
		results = dfs(root.Left, results)
	} else {
		results = append(results, "null")
	}

	if root.Right != nil {
		results = dfs(root.Right, results)
	} else {
		results = append(results, "null")
	}

	return results
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	arr1 := make([]string, 0)
	dfs1 := dfs(root, arr1)

	arr2 := make([]string, 0)
	dfs2 := dfs(subRoot, arr2)

	s1 := strings.Join(dfs1, ",")
	s2 := strings.Join(dfs2, ",")

	fmt.Println(s1, s2)

	return strings.Contains(s1, s2)
}
