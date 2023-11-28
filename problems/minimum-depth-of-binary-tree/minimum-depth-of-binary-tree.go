package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	q := make([]*TreeNode, 1)
	q[0] = root
	lvl := 0

	for len(q) != 0 {
		items := q[:]

		for _, item := range items {
			if item.Left == nil && item.Right == nil {
				return lvl + 1
			}

			if item.Left != nil {
				q = append(q, item.Left)
			}

			if item.Right != nil {
				q = append(q, item.Right)
			}
		}

		q = q[len(items):]
		lvl++
	}

	return lvl
}
