package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.Equal(t, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}), [][]string{
		[]string{"bat"},
		[]string{"nat", "tan"},
		[]string{"ate", "eat", "tea"},
	})
}
