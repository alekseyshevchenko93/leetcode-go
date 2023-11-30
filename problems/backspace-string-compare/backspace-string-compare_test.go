package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	tt := []struct {
		s1, s2 string
		output bool
	}{
		{
			"ab#c", "ad#c",
			true,
		},
		{
			"ab##", "c#d#",
			true,
		},
		{
			"a#c", "b",
			false,
		},
		{
			"y#fo##f", "y#f#o##f",
			true,
		},
	}

	for _, v := range tt {
		require.Equal(t, v.output, backspaceCompare(v.s1, v.s2))
	}
}
