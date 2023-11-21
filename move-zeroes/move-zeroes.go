package main

func moveZeroes(nums []int) {
	s := 0

	for i, v := range nums {
		if v != 0 {
			if i != s {
				temp := nums[i]
				nums[i] = nums[s]
				nums[s] = temp
			}

			s++
		}
	}
}
