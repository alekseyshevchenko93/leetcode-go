package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-z0-9]`)
	rs := re.ReplaceAll([]byte(strings.ToLower(s)), []byte(""))

	l := 0
	r := len(rs) - 1

	for l <= r {
		if rs[l] != rs[r] {
			return false
		}

		l++
		r--
	}

	return true
}
