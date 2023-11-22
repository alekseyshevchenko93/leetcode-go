package main

func productExceptSelf(nums []int) []int {
	arr1 := make([]int, len(nums))
	arr2 := make([]int, len(nums))

	for i, v := range nums {
		if i == 0 {
			arr1[i] = v
		} else {
			arr1[i] = v * arr1[i-1]
		}
	}

	l := len(nums) - 1

	for i := range nums {
		if i == 0 {
			arr2[i] = nums[l-i]
		} else {
			arr2[i] = nums[l-i] * arr2[i-1]
		}
	}

	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = arr2[l-i-1]
		} else if i == len(nums)-1 {
			result[i] = arr1[i-1]
		} else {
			result[i] = arr1[i-1] * arr2[l-i-1]
		}
	}

	return result
}
