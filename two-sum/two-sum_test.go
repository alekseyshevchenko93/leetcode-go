package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.Equal(t, twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1})
	require.Equal(t, twoSum([]int{3, 2, 4}, 6), []int{1, 2})
	require.Equal(t, twoSum([]int{3, 3}, 6), []int{0, 1})
}
