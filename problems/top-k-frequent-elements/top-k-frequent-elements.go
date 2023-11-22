package main

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		_, found := m[v]

		if found {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	pairs := make([][]int, 0, len(m))

	for k, v := range m {
		pairs = append(pairs, []int{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		pair1 := pairs[i]
		pair2 := pairs[j]

		return pair1[1] > pair2[1]
	})

	topK := pairs[:k]
	result := make([]int, 0, k)

	for _, v := range topK {
		result = append(result, v[0])
	}

	return result
}
