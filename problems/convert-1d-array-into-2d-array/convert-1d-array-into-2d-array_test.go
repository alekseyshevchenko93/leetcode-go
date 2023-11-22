package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.Equal(t, [][]int{
		{1, 2},
		{3, 4},
	}, construct2DArray([]int{1, 2, 3, 4}, 2, 2))

	require.Equal(t, [][]int{
		{1, 2, 3},
	}, construct2DArray([]int{1, 2, 3}, 1, 3))

	require.Equal(t, [][]int{}, construct2DArray([]int{1, 2}, 1, 1))

	require.Equal(t, [][]int{}, construct2DArray([]int{3}, 1, 2))
}
