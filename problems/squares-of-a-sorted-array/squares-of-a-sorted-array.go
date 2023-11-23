package main

import (
	"math"
)

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1
	lenNums := r

	result := make([]int, len(nums))
	i := 0

	for l <= r {
		left := math.Pow(float64(nums[l]), 2)
		right := math.Pow(float64(nums[r]), 2)

		if left > right {
			result[lenNums-i] = int(left)
			l++
		} else {
			result[lenNums-i] = int(right)
			r--
		}

		i++
	}

	return result
}
