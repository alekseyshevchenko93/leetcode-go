package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, num := range nums {
		_, found := m[num]

		if found {
			return true
		}

		m[num] = struct{}{}
	}

	return false
}
