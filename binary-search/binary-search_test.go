package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	tt := []struct {
		input  []int
		target int
		output int
	}{
		{
			[]int{-1, 0, 3, 5, 9, 12},
			9,
			4,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			2,
			-1,
		},
	}

	for _, v := range tt {
		require.Equal(t, v.output, search(v.input, v.target))
	}
}
