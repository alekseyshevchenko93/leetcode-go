package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.False(t, containsDuplicate([]int{1, 2, 3, 4, 5}))
	require.True(t, containsDuplicate([]int{1, 2, 3, 3, 5}))
}
