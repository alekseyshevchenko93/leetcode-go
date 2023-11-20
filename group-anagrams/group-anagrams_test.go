package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.Equal(t, [][]string{
		[]string{"eat", "tea", "ate"},
		[]string{"tan", "nat"},
		[]string{"bat"},
	}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
