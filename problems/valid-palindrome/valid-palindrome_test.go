package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	tt := []struct {
		input  string
		output bool
	}{
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"race a car",
			false,
		},
		{
			" ",
			true,
		},
	}

	for _, testCase := range tt {
		require.Equal(t, testCase.output, isPalindrome(testCase.input))
	}
}
