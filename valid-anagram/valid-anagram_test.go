package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	require.False(t, isAnagram("a", "aaa"))
	require.True(t, isAnagram("anagram", "nagaram"))
	require.False(t, isAnagram("rat", "car"))
}
