package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.Equal(t, 2, missingNumber([]int{3, 0, 1}))
	require.Equal(t, 2, missingNumber([]int{0, 1}))
	require.Equal(t, 8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

}
