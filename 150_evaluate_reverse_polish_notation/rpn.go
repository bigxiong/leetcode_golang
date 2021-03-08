package main

import (
	"strconv"
)

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

func main() {
	tokens := []string{"4","13","5","/","+"}
	evalRPN(tokens)
}