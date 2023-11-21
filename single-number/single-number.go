package main

func singleNumber(nums []int) int {
	mask := 0

	for _, v := range nums {
		mask ^= v
	}

	return mask
}
