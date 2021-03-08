package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func calculate(s string) int {
	rpn := infixToPostfix(s)
	fmt.Println(rpn)

	return evalRPN(rpn)
}

func evalRPN(tokens []string) int {
	var stack []int
	for _, val := range tokens{
		var op1 int
		var op2 int
		switch val {
		case "+":
			op1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			op2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, op1 + op2)
		case "-":
			op1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			op2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, op2 - op1)
		case "*":
			op1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			op2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, op2 * op1)
		case "/":
			op1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			op2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, op2 / op1)
		default:
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func infixToPostfix(infix string) []string {
	var postfix []string
	var stack []byte
	for i := 0; i < len(infix); i++ {
		c := infix[i]
		if unicode.IsSpace(rune(c)) {
			continue
		} else if unicode.IsDigit(rune(c)) {
			tmp:=string(c)
			for i < len(infix)-1 && infix[i+1]>='0' && infix[i+1]<='9'{
				tmp+=string(infix[i+1])
				i++
			}
			postfix = append(postfix, tmp)
		} else if c == '(' {
			stack = append(stack, c)
		} else if c ==')' {
			var top byte
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				stack, top = stack[:len(stack)-1], stack[len(stack)-1]
				postfix = append(postfix, string(top))
			}
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			}
		} else {
			if len(stack) == 0 || precedence(c) > precedence(stack[len(stack)-1]) {
				stack = append(stack, c)
			} else {
				var top byte
				for len(stack) > 0 && precedence(c) <= precedence(stack[len(stack)-1]) {
					stack, top = stack[:len(stack)-1], stack[len(stack)-1]
					postfix = append(postfix, string(top))
				}
				stack = append(stack, c)
			}
		}
	}


	for len(stack) > 0 {
		var top byte
		stack, top = stack[:len(stack)-1], stack[len(stack)-1]
		postfix = append(postfix, string(top))
	}
	return postfix
}

func precedence(c byte) int {
	if c =='+' || c == '-' {
		return 1
	}
	if c == '*' || c == '/' {
		return 2
	}
	return -1
}

func main() {
	fmt.Println(calculate("2147483647"))
}