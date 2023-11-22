package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		runeToCountMap := make(map[rune]int)

		for _, v := range str {
			_, found := runeToCountMap[v]

			if found {
				runeToCountMap[v]++
			} else {
				runeToCountMap[v] = 1
			}
		}

		resultKey := make([]string, 0)

		for k, v := range runeToCountMap {
			resultKey = append(resultKey, string(k)+fmt.Sprint(v))
		}

		sort.Strings(resultKey)

		resultKeyStr := strings.Join(resultKey, ":")

		_, found := m[resultKeyStr]

		if found {
			m[resultKeyStr] = append(m[resultKeyStr], str)
		} else {
			m[resultKeyStr] = []string{str}
		}
	}

	var result [][]string

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
