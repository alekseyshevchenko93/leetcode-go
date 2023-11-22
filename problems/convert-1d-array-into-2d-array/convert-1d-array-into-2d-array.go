package main

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	matrix := make([][]int, m)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	c := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = original[c]
			c++
		}
	}

	return matrix
}
