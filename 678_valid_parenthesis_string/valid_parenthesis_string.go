package main

import "fmt"

func checkValidString(s string) bool {
	var stack []rune
	for _, t := range s {
		if t == '{' || t == '[' || t == '(' {
			stack = append(stack, t)
		} else  {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (t == '}' && top == '{') || (t == ']' && top == '[') || (t == ')' && top == '('){
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}
	return len(stack) == 0
}

func main() {
	fmt.Println(checkValidString("(*)"))

	fmt.Println(checkValidString("(*))"))
}
