package main

func backspaceCompare(s string, t string) bool {
	stack1 := make([]rune, 0, len(s))
	stack2 := make([]rune, 0, len(t))

	for _, c := range s {
		if c == '#' {
			if len(stack1) > 0 {
				stack1 = stack1[:len(stack1)-1]
			}
		} else {
			stack1 = append(stack1, c)
		}
	}

	for _, c := range t {
		if c == '#' {
			if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			}
		} else {
			stack2 = append(stack2, c)
		}
	}

	if len(stack1) != len(stack2) {
		return false
	}

	i := 0

	for i < len(stack1) {
		if stack1[i] != stack2[i] {
			return false
		}
		i++
	}

	return true
}
