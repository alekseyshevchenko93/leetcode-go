package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	tt := []struct {
		input  []int
		output []int
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0, 0, 1, 1},
			[]int{1, 1, 0, 0},
		},
		{
			[]int{0, 1, 0, 1},
			[]int{1, 1, 0, 0},
		},
		{
			[]int{0, 1, 0, 0, 0, 0, 1},
			[]int{1, 1, 0, 0, 0, 0, 0},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
	}

	for _, v := range tt {
		moveZeroes(v.input)
		require.Equal(t, v.output, v.input)
	}
}
