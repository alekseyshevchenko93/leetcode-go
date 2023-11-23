package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}
	results := []float64{}

	for len(q) != 0 {
		items := q[:]
		var avg float64 = 0

		for _, item := range items {
			avg += float64(item.Val)

			if item.Left != nil {
				q = append(q, item.Left)
			}

			if item.Right != nil {
				q = append(q, item.Right)
			}
		}

		results = append(results, avg/float64(len(items)))
		q = q[len(items):]
	}

	return results
}
