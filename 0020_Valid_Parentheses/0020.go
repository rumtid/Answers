package main

func isValid(s string) bool {
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, 1)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != 1 {
				return false
			}
			stack = stack[:len(stack)-1]
		case '{':
			stack = append(stack, 2)
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != 2 {
				return false
			}
			stack = stack[:len(stack)-1]
		case '[':
			stack = append(stack, 3)
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != 3 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
