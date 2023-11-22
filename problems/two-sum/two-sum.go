package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		idx, found := m[target-v]

		if found {
			return []int{idx, i}
		}

		m[v] = i
	}

	return nil
}
