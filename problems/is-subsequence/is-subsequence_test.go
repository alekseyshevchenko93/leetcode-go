package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	tt := []struct {
		s      string
		t      string
		output bool
	}{
		{
			"abc",
			"ahbgdc",
			true,
		},
		{
			"axc",
			"ahbgdc",
			false,
		},
	}

	for _, v := range tt {
		require.Equal(t, v.output, isSubsequence(v.s, v.t))
	}
}
