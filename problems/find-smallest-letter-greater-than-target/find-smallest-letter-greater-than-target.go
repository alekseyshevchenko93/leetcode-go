package main

func search(nums []byte, target byte) int {
	l := 0
	r := len(nums) - 1
	idx := -1

	for l <= r {
		mid := int((l + r) / 2)

		if nums[mid] == target {
			l = mid + 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
			idx = mid
		}
	}

	return idx
}

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := search(letters, target)

	if idx == -1 {
		return letters[0]
	}

	return letters[idx]
}
