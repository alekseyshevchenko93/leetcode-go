package main

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	result := make([]int, 0)

	for i := 1; i <= len(nums); i++ {
		_, found := m[i]

		if !found {
			result = append(result, i)
		}
	}

	return result
}
