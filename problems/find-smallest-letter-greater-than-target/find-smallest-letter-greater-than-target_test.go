package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem(t *testing.T) {
	tt := []struct {
		input  []byte
		target byte
		output byte
	}{
		{
			[]byte("cfj"),
			[]byte("a")[0],
			[]byte("c")[0],
		},
		{
			[]byte("cfj"),
			[]byte("c")[0],
			[]byte("f")[0],
		},
		{
			[]byte("xxyy"),
			[]byte("z")[0],
			[]byte("x")[0],
		},
		{
			[]byte("abcdefg"),
			[]byte("c")[0],
			[]byte("d")[0],
		},
		{
			[]byte("cfj"),
			[]byte("d")[0],
			[]byte("f")[0],
		},
	}

	for _, v := range tt {
		require.Equal(t, v.output, nextGreatestLetter(v.input, v.target))
	}
}
