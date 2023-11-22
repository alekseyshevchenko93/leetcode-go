package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range s {
		_, found := m1[v]

		if found {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}

	for _, v := range t {
		_, found := m2[v]

		if found {
			m2[v]++
		} else {
			m2[v] = 1
		}
	}

	for k, v := range m1 {
		v2, found := m2[k]

		if !found || (found && v2 != v) {
			return false
		}
	}

	for k, v := range m2 {
		v1, found := m1[k]

		if !found || (found && v1 != v) {
			return false
		}
	}

	return true
}
