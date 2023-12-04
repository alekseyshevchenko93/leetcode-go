package main

import "fmt"

func isSubsequence(s string, t string) bool {
	var (
		i    = 0
		j    = 0
		len1 = len(s)
		len2 = len(t)
	)

	for i < len1 && j < len2 {
		if s[i] == t[j] {
			i++
			j++
			continue
		}

		j++
	}

	fmt.Println("i,j", i, j)

	return i == len1
}
