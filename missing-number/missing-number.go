package main

func missingNumber(nums []int) int {
	arr := make([]int, len(nums)+1)

	for _, v := range nums {
		arr[v] = 1
	}

	for k, v := range arr {
		if v == 0 {
			return k
		}
	}

	return 0
}
