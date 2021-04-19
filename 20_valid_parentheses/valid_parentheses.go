package main

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if pairs[c] > 0 && pairs[c] != top {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	isValid("()")
}
